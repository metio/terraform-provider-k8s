/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package volsync_backube_v1alpha1

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
	_ datasource.DataSource = &VolsyncBackubeReplicationSourceV1Alpha1Manifest{}
)

func NewVolsyncBackubeReplicationSourceV1Alpha1Manifest() datasource.DataSource {
	return &VolsyncBackubeReplicationSourceV1Alpha1Manifest{}
}

type VolsyncBackubeReplicationSourceV1Alpha1Manifest struct{}

type VolsyncBackubeReplicationSourceV1Alpha1ManifestData struct {
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
		External *struct {
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			Provider   *string            `tfsdk:"provider" json:"provider,omitempty"`
		} `tfsdk:"external" json:"external,omitempty"`
		Paused *bool `tfsdk:"paused" json:"paused,omitempty"`
		Rclone *struct {
			AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
			Capacity    *string   `tfsdk:"capacity" json:"capacity,omitempty"`
			CopyMethod  *string   `tfsdk:"copy_method" json:"copyMethod,omitempty"`
			CustomCA    *struct {
				ConfigMapName *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				Key           *string `tfsdk:"key" json:"key,omitempty"`
				SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"custom_ca" json:"customCA,omitempty"`
			MoverPodLabels *map[string]string `tfsdk:"mover_pod_labels" json:"moverPodLabels,omitempty"`
			MoverResources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"mover_resources" json:"moverResources,omitempty"`
			MoverSecurityContext *struct {
				FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
				RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
				RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
				RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				SeLinuxOptions      *struct {
					Level *string `tfsdk:"level" json:"level,omitempty"`
					Role  *string `tfsdk:"role" json:"role,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					User  *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
				SeccompProfile *struct {
					LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
					Type             *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
				SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
				Sysctls            *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"sysctls" json:"sysctls,omitempty"`
				WindowsOptions *struct {
					GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
					HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
					RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
			} `tfsdk:"mover_security_context" json:"moverSecurityContext,omitempty"`
			MoverServiceAccount     *string `tfsdk:"mover_service_account" json:"moverServiceAccount,omitempty"`
			RcloneConfig            *string `tfsdk:"rclone_config" json:"rcloneConfig,omitempty"`
			RcloneConfigSection     *string `tfsdk:"rclone_config_section" json:"rcloneConfigSection,omitempty"`
			RcloneDestPath          *string `tfsdk:"rclone_dest_path" json:"rcloneDestPath,omitempty"`
			StorageClassName        *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeSnapshotClassName *string `tfsdk:"volume_snapshot_class_name" json:"volumeSnapshotClassName,omitempty"`
		} `tfsdk:"rclone" json:"rclone,omitempty"`
		Restic *struct {
			AccessModes           *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
			CacheAccessModes      *[]string `tfsdk:"cache_access_modes" json:"cacheAccessModes,omitempty"`
			CacheCapacity         *string   `tfsdk:"cache_capacity" json:"cacheCapacity,omitempty"`
			CacheStorageClassName *string   `tfsdk:"cache_storage_class_name" json:"cacheStorageClassName,omitempty"`
			Capacity              *string   `tfsdk:"capacity" json:"capacity,omitempty"`
			CopyMethod            *string   `tfsdk:"copy_method" json:"copyMethod,omitempty"`
			CustomCA              *struct {
				ConfigMapName *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				Key           *string `tfsdk:"key" json:"key,omitempty"`
				SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"custom_ca" json:"customCA,omitempty"`
			MoverPodLabels *map[string]string `tfsdk:"mover_pod_labels" json:"moverPodLabels,omitempty"`
			MoverResources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"mover_resources" json:"moverResources,omitempty"`
			MoverSecurityContext *struct {
				FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
				RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
				RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
				RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				SeLinuxOptions      *struct {
					Level *string `tfsdk:"level" json:"level,omitempty"`
					Role  *string `tfsdk:"role" json:"role,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					User  *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
				SeccompProfile *struct {
					LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
					Type             *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
				SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
				Sysctls            *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"sysctls" json:"sysctls,omitempty"`
				WindowsOptions *struct {
					GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
					HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
					RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
			} `tfsdk:"mover_security_context" json:"moverSecurityContext,omitempty"`
			MoverServiceAccount *string `tfsdk:"mover_service_account" json:"moverServiceAccount,omitempty"`
			PruneIntervalDays   *int64  `tfsdk:"prune_interval_days" json:"pruneIntervalDays,omitempty"`
			Repository          *string `tfsdk:"repository" json:"repository,omitempty"`
			Retain              *struct {
				Daily   *int64  `tfsdk:"daily" json:"daily,omitempty"`
				Hourly  *int64  `tfsdk:"hourly" json:"hourly,omitempty"`
				Last    *string `tfsdk:"last" json:"last,omitempty"`
				Monthly *int64  `tfsdk:"monthly" json:"monthly,omitempty"`
				Weekly  *int64  `tfsdk:"weekly" json:"weekly,omitempty"`
				Within  *string `tfsdk:"within" json:"within,omitempty"`
				Yearly  *int64  `tfsdk:"yearly" json:"yearly,omitempty"`
			} `tfsdk:"retain" json:"retain,omitempty"`
			StorageClassName        *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			Unlock                  *string `tfsdk:"unlock" json:"unlock,omitempty"`
			VolumeSnapshotClassName *string `tfsdk:"volume_snapshot_class_name" json:"volumeSnapshotClassName,omitempty"`
		} `tfsdk:"restic" json:"restic,omitempty"`
		Rsync *struct {
			AccessModes    *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
			Address        *string            `tfsdk:"address" json:"address,omitempty"`
			Capacity       *string            `tfsdk:"capacity" json:"capacity,omitempty"`
			CopyMethod     *string            `tfsdk:"copy_method" json:"copyMethod,omitempty"`
			MoverPodLabels *map[string]string `tfsdk:"mover_pod_labels" json:"moverPodLabels,omitempty"`
			MoverResources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"mover_resources" json:"moverResources,omitempty"`
			MoverServiceAccount     *string `tfsdk:"mover_service_account" json:"moverServiceAccount,omitempty"`
			Path                    *string `tfsdk:"path" json:"path,omitempty"`
			Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
			ServiceType             *string `tfsdk:"service_type" json:"serviceType,omitempty"`
			SshKeys                 *string `tfsdk:"ssh_keys" json:"sshKeys,omitempty"`
			SshUser                 *string `tfsdk:"ssh_user" json:"sshUser,omitempty"`
			StorageClassName        *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeSnapshotClassName *string `tfsdk:"volume_snapshot_class_name" json:"volumeSnapshotClassName,omitempty"`
		} `tfsdk:"rsync" json:"rsync,omitempty"`
		RsyncTLS *struct {
			AccessModes    *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
			Address        *string            `tfsdk:"address" json:"address,omitempty"`
			Capacity       *string            `tfsdk:"capacity" json:"capacity,omitempty"`
			CopyMethod     *string            `tfsdk:"copy_method" json:"copyMethod,omitempty"`
			KeySecret      *string            `tfsdk:"key_secret" json:"keySecret,omitempty"`
			MoverPodLabels *map[string]string `tfsdk:"mover_pod_labels" json:"moverPodLabels,omitempty"`
			MoverResources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"mover_resources" json:"moverResources,omitempty"`
			MoverSecurityContext *struct {
				FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
				RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
				RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
				RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				SeLinuxOptions      *struct {
					Level *string `tfsdk:"level" json:"level,omitempty"`
					Role  *string `tfsdk:"role" json:"role,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					User  *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
				SeccompProfile *struct {
					LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
					Type             *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
				SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
				Sysctls            *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"sysctls" json:"sysctls,omitempty"`
				WindowsOptions *struct {
					GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
					HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
					RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
			} `tfsdk:"mover_security_context" json:"moverSecurityContext,omitempty"`
			MoverServiceAccount     *string `tfsdk:"mover_service_account" json:"moverServiceAccount,omitempty"`
			Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
			StorageClassName        *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeSnapshotClassName *string `tfsdk:"volume_snapshot_class_name" json:"volumeSnapshotClassName,omitempty"`
		} `tfsdk:"rsync_tls" json:"rsyncTLS,omitempty"`
		SourcePVC *string `tfsdk:"source_pvc" json:"sourcePVC,omitempty"`
		Syncthing *struct {
			ConfigAccessModes      *[]string          `tfsdk:"config_access_modes" json:"configAccessModes,omitempty"`
			ConfigCapacity         *string            `tfsdk:"config_capacity" json:"configCapacity,omitempty"`
			ConfigStorageClassName *string            `tfsdk:"config_storage_class_name" json:"configStorageClassName,omitempty"`
			MoverPodLabels         *map[string]string `tfsdk:"mover_pod_labels" json:"moverPodLabels,omitempty"`
			MoverResources         *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"mover_resources" json:"moverResources,omitempty"`
			MoverSecurityContext *struct {
				FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
				RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
				RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
				RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				SeLinuxOptions      *struct {
					Level *string `tfsdk:"level" json:"level,omitempty"`
					Role  *string `tfsdk:"role" json:"role,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					User  *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
				SeccompProfile *struct {
					LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
					Type             *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
				SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
				Sysctls            *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"sysctls" json:"sysctls,omitempty"`
				WindowsOptions *struct {
					GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
					HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
					RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
			} `tfsdk:"mover_security_context" json:"moverSecurityContext,omitempty"`
			MoverServiceAccount *string `tfsdk:"mover_service_account" json:"moverServiceAccount,omitempty"`
			Peers               *[]struct {
				ID         *string `tfsdk:"id" json:"ID,omitempty"`
				Address    *string `tfsdk:"address" json:"address,omitempty"`
				Introducer *bool   `tfsdk:"introducer" json:"introducer,omitempty"`
			} `tfsdk:"peers" json:"peers,omitempty"`
			ServiceType *string `tfsdk:"service_type" json:"serviceType,omitempty"`
		} `tfsdk:"syncthing" json:"syncthing,omitempty"`
		Trigger *struct {
			Manual   *string `tfsdk:"manual" json:"manual,omitempty"`
			Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		} `tfsdk:"trigger" json:"trigger,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VolsyncBackubeReplicationSourceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_volsync_backube_replication_source_v1alpha1_manifest"
}

func (r *VolsyncBackubeReplicationSourceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ReplicationSource defines the source for a replicated volume",
		MarkdownDescription: "ReplicationSource defines the source for a replicated volume",
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
				Description:         "spec is the desired state of the ReplicationSource, including the replication method to use and its configuration.",
				MarkdownDescription: "spec is the desired state of the ReplicationSource, including the replication method to use and its configuration.",
				Attributes: map[string]schema.Attribute{
					"external": schema.SingleNestedAttribute{
						Description:         "external defines the configuration when using an external replication provider.",
						MarkdownDescription: "external defines the configuration when using an external replication provider.",
						Attributes: map[string]schema.Attribute{
							"parameters": schema.MapAttribute{
								Description:         "parameters are provider-specific key/value configuration parameters. For more information, please see the documentation of the specific replication provider being used.",
								MarkdownDescription: "parameters are provider-specific key/value configuration parameters. For more information, please see the documentation of the specific replication provider being used.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider": schema.StringAttribute{
								Description:         "provider is the name of the external replication provider. The name should be of the form: domain.com/provider.",
								MarkdownDescription: "provider is the name of the external replication provider. The name should be of the form: domain.com/provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": schema.BoolAttribute{
						Description:         "paused can be used to temporarily stop replication. Defaults to 'false'.",
						MarkdownDescription: "paused can be used to temporarily stop replication. Defaults to 'false'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rclone": schema.SingleNestedAttribute{
						Description:         "rclone defines the configuration when using Rclone-based replication.",
						MarkdownDescription: "rclone defines the configuration when using Rclone-based replication.",
						Attributes: map[string]schema.Attribute{
							"access_modes": schema.ListAttribute{
								Description:         "accessModes can be used to override the accessModes of the PiT image.",
								MarkdownDescription: "accessModes can be used to override the accessModes of the PiT image.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capacity": schema.StringAttribute{
								Description:         "capacity can be used to override the capacity of the PiT image.",
								MarkdownDescription: "capacity can be used to override the capacity of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"copy_method": schema.StringAttribute{
								Description:         "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								MarkdownDescription: "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Direct", "None", "Clone", "Snapshot"),
								},
							},

							"custom_ca": schema.SingleNestedAttribute{
								Description:         "customCA is a custom CA that will be used to verify the remote",
								MarkdownDescription: "customCA is a custom CA that will be used to verify the remote",
								Attributes: map[string]schema.Attribute{
									"config_map_name": schema.StringAttribute{
										Description:         "The name of a ConfigMap that contains the custom CA certificate If ConfigMapName is used then SecretName should not be set",
										MarkdownDescription: "The name of a ConfigMap that contains the custom CA certificate If ConfigMapName is used then SecretName should not be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "The key within the Secret or ConfigMap containing the CA certificate",
										MarkdownDescription: "The key within the Secret or ConfigMap containing the CA certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of a Secret that contains the custom CA certificate If SecretName is used then ConfigMapName should not be set",
										MarkdownDescription: "The name of a Secret that contains the custom CA certificate If SecretName is used then ConfigMapName should not be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mover_pod_labels": schema.MapAttribute{
								Description:         "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								MarkdownDescription: "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_resources": schema.SingleNestedAttribute{
								Description:         "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
								MarkdownDescription: "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
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

							"mover_security_context": schema.SingleNestedAttribute{
								Description:         "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								MarkdownDescription: "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								Attributes: map[string]schema.Attribute{
									"fs_group": schema.Int64Attribute{
										Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fs_group_change_policy": schema.StringAttribute{
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_group": schema.Int64Attribute{
										Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_non_root": schema.BoolAttribute{
										Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user": schema.Int64Attribute{
										Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"se_linux_options": schema.SingleNestedAttribute{
										Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"level": schema.StringAttribute{
												Description:         "Level is SELinux level label that applies to the container.",
												MarkdownDescription: "Level is SELinux level label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a SELinux role label that applies to the container.",
												MarkdownDescription: "Role is a SELinux role label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type is a SELinux type label that applies to the container.",
												MarkdownDescription: "Type is a SELinux type label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is a SELinux user label that applies to the container.",
												MarkdownDescription: "User is a SELinux user label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"seccomp_profile": schema.SingleNestedAttribute{
										Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"localhost_profile": schema.StringAttribute{
												Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"supplemental_groups": schema.ListAttribute{
										Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysctls": schema.ListNestedAttribute{
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of a property to set",
													MarkdownDescription: "Name of a property to set",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value of a property to set",
													MarkdownDescription: "Value of a property to set",
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

									"windows_options": schema.SingleNestedAttribute{
										Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										Attributes: map[string]schema.Attribute{
											"gmsa_credential_spec": schema.StringAttribute{
												Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gmsa_credential_spec_name": schema.StringAttribute{
												Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_process": schema.BoolAttribute{
												Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user_name": schema.StringAttribute{
												Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"mover_service_account": schema.StringAttribute{
								Description:         "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								MarkdownDescription: "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rclone_config": schema.StringAttribute{
								Description:         "RcloneConfig is the rclone secret name",
								MarkdownDescription: "RcloneConfig is the rclone secret name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rclone_config_section": schema.StringAttribute{
								Description:         "RcloneConfigSection is the section in rclone_config file to use for the current job.",
								MarkdownDescription: "RcloneConfigSection is the section in rclone_config file to use for the current job.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rclone_dest_path": schema.StringAttribute{
								Description:         "RcloneDestPath is the remote path to sync to.",
								MarkdownDescription: "RcloneDestPath is the remote path to sync to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class_name": schema.StringAttribute{
								Description:         "storageClassName can be used to override the StorageClass of the PiT image.",
								MarkdownDescription: "storageClassName can be used to override the StorageClass of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_snapshot_class_name": schema.StringAttribute{
								Description:         "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								MarkdownDescription: "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"restic": schema.SingleNestedAttribute{
						Description:         "restic defines the configuration when using Restic-based replication.",
						MarkdownDescription: "restic defines the configuration when using Restic-based replication.",
						Attributes: map[string]schema.Attribute{
							"access_modes": schema.ListAttribute{
								Description:         "accessModes can be used to override the accessModes of the PiT image.",
								MarkdownDescription: "accessModes can be used to override the accessModes of the PiT image.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_access_modes": schema.ListAttribute{
								Description:         "CacheAccessModes can be used to set the accessModes of restic metadata cache volume",
								MarkdownDescription: "CacheAccessModes can be used to set the accessModes of restic metadata cache volume",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_capacity": schema.StringAttribute{
								Description:         "cacheCapacity can be used to set the size of the restic metadata cache volume",
								MarkdownDescription: "cacheCapacity can be used to set the size of the restic metadata cache volume",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_storage_class_name": schema.StringAttribute{
								Description:         "cacheStorageClassName can be used to set the StorageClass of the restic metadata cache volume",
								MarkdownDescription: "cacheStorageClassName can be used to set the StorageClass of the restic metadata cache volume",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capacity": schema.StringAttribute{
								Description:         "capacity can be used to override the capacity of the PiT image.",
								MarkdownDescription: "capacity can be used to override the capacity of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"copy_method": schema.StringAttribute{
								Description:         "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								MarkdownDescription: "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Direct", "None", "Clone", "Snapshot"),
								},
							},

							"custom_ca": schema.SingleNestedAttribute{
								Description:         "customCA is a custom CA that will be used to verify the remote",
								MarkdownDescription: "customCA is a custom CA that will be used to verify the remote",
								Attributes: map[string]schema.Attribute{
									"config_map_name": schema.StringAttribute{
										Description:         "The name of a ConfigMap that contains the custom CA certificate If ConfigMapName is used then SecretName should not be set",
										MarkdownDescription: "The name of a ConfigMap that contains the custom CA certificate If ConfigMapName is used then SecretName should not be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "The key within the Secret or ConfigMap containing the CA certificate",
										MarkdownDescription: "The key within the Secret or ConfigMap containing the CA certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of a Secret that contains the custom CA certificate If SecretName is used then ConfigMapName should not be set",
										MarkdownDescription: "The name of a Secret that contains the custom CA certificate If SecretName is used then ConfigMapName should not be set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mover_pod_labels": schema.MapAttribute{
								Description:         "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								MarkdownDescription: "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_resources": schema.SingleNestedAttribute{
								Description:         "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
								MarkdownDescription: "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
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

							"mover_security_context": schema.SingleNestedAttribute{
								Description:         "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								MarkdownDescription: "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								Attributes: map[string]schema.Attribute{
									"fs_group": schema.Int64Attribute{
										Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fs_group_change_policy": schema.StringAttribute{
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_group": schema.Int64Attribute{
										Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_non_root": schema.BoolAttribute{
										Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user": schema.Int64Attribute{
										Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"se_linux_options": schema.SingleNestedAttribute{
										Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"level": schema.StringAttribute{
												Description:         "Level is SELinux level label that applies to the container.",
												MarkdownDescription: "Level is SELinux level label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a SELinux role label that applies to the container.",
												MarkdownDescription: "Role is a SELinux role label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type is a SELinux type label that applies to the container.",
												MarkdownDescription: "Type is a SELinux type label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is a SELinux user label that applies to the container.",
												MarkdownDescription: "User is a SELinux user label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"seccomp_profile": schema.SingleNestedAttribute{
										Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"localhost_profile": schema.StringAttribute{
												Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"supplemental_groups": schema.ListAttribute{
										Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysctls": schema.ListNestedAttribute{
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of a property to set",
													MarkdownDescription: "Name of a property to set",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value of a property to set",
													MarkdownDescription: "Value of a property to set",
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

									"windows_options": schema.SingleNestedAttribute{
										Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										Attributes: map[string]schema.Attribute{
											"gmsa_credential_spec": schema.StringAttribute{
												Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gmsa_credential_spec_name": schema.StringAttribute{
												Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_process": schema.BoolAttribute{
												Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user_name": schema.StringAttribute{
												Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"mover_service_account": schema.StringAttribute{
								Description:         "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								MarkdownDescription: "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prune_interval_days": schema.Int64Attribute{
								Description:         "PruneIntervalDays define how often to prune the repository",
								MarkdownDescription: "PruneIntervalDays define how often to prune the repository",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "Repository is the secret name containing repository info",
								MarkdownDescription: "Repository is the secret name containing repository info",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retain": schema.SingleNestedAttribute{
								Description:         "ResticRetainPolicy define the retain policy",
								MarkdownDescription: "ResticRetainPolicy define the retain policy",
								Attributes: map[string]schema.Attribute{
									"daily": schema.Int64Attribute{
										Description:         "Daily defines the number of snapshots to be kept daily",
										MarkdownDescription: "Daily defines the number of snapshots to be kept daily",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hourly": schema.Int64Attribute{
										Description:         "Hourly defines the number of snapshots to be kept hourly",
										MarkdownDescription: "Hourly defines the number of snapshots to be kept hourly",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"last": schema.StringAttribute{
										Description:         "Last defines the number of snapshots to be kept",
										MarkdownDescription: "Last defines the number of snapshots to be kept",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"monthly": schema.Int64Attribute{
										Description:         "Monthly defines the number of snapshots to be kept monthly",
										MarkdownDescription: "Monthly defines the number of snapshots to be kept monthly",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"weekly": schema.Int64Attribute{
										Description:         "Weekly defines the number of snapshots to be kept weekly",
										MarkdownDescription: "Weekly defines the number of snapshots to be kept weekly",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"within": schema.StringAttribute{
										Description:         "Within defines the number of snapshots to be kept Within the given time period",
										MarkdownDescription: "Within defines the number of snapshots to be kept Within the given time period",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"yearly": schema.Int64Attribute{
										Description:         "Yearly defines the number of snapshots to be kept yearly",
										MarkdownDescription: "Yearly defines the number of snapshots to be kept yearly",
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
								Description:         "storageClassName can be used to override the StorageClass of the PiT image.",
								MarkdownDescription: "storageClassName can be used to override the StorageClass of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unlock": schema.StringAttribute{
								Description:         "unlock is a string value that schedules an unlock on the restic repository during the next sync operation. Once a sync completes then status.restic.lastUnlocked is set to the same string value. To unlock a repository, set spec.restic.unlock to a known value and then wait for lastUnlocked to be updated by the operator to the same value, which means that the sync unlocked the repository by running a restic unlock command and then ran a backup. Unlock will not be run again unless spec.restic.unlock is set to a different value.",
								MarkdownDescription: "unlock is a string value that schedules an unlock on the restic repository during the next sync operation. Once a sync completes then status.restic.lastUnlocked is set to the same string value. To unlock a repository, set spec.restic.unlock to a known value and then wait for lastUnlocked to be updated by the operator to the same value, which means that the sync unlocked the repository by running a restic unlock command and then ran a backup. Unlock will not be run again unless spec.restic.unlock is set to a different value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_snapshot_class_name": schema.StringAttribute{
								Description:         "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								MarkdownDescription: "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rsync": schema.SingleNestedAttribute{
						Description:         "rsync defines the configuration when using Rsync-based replication.",
						MarkdownDescription: "rsync defines the configuration when using Rsync-based replication.",
						Attributes: map[string]schema.Attribute{
							"access_modes": schema.ListAttribute{
								Description:         "accessModes can be used to override the accessModes of the PiT image.",
								MarkdownDescription: "accessModes can be used to override the accessModes of the PiT image.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"address": schema.StringAttribute{
								Description:         "address is the remote address to connect to for replication.",
								MarkdownDescription: "address is the remote address to connect to for replication.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capacity": schema.StringAttribute{
								Description:         "capacity can be used to override the capacity of the PiT image.",
								MarkdownDescription: "capacity can be used to override the capacity of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"copy_method": schema.StringAttribute{
								Description:         "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								MarkdownDescription: "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Direct", "None", "Clone", "Snapshot"),
								},
							},

							"mover_pod_labels": schema.MapAttribute{
								Description:         "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								MarkdownDescription: "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_resources": schema.SingleNestedAttribute{
								Description:         "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
								MarkdownDescription: "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
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

							"mover_service_account": schema.StringAttribute{
								Description:         "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as the ReplicationSource.",
								MarkdownDescription: "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as the ReplicationSource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "path is the remote path to rsync to. Defaults to '/'",
								MarkdownDescription: "path is the remote path to rsync to. Defaults to '/'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "port is the SSH port to connect to for replication. Defaults to 22.",
								MarkdownDescription: "port is the SSH port to connect to for replication. Defaults to 22.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(65535),
								},
							},

							"service_type": schema.StringAttribute{
								Description:         "serviceType determines the Service type that will be created for incoming SSH connections.",
								MarkdownDescription: "serviceType determines the Service type that will be created for incoming SSH connections.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_keys": schema.StringAttribute{
								Description:         "sshKeys is the name of a Secret that contains the SSH keys to be used for authentication. If not provided, the keys will be generated.",
								MarkdownDescription: "sshKeys is the name of a Secret that contains the SSH keys to be used for authentication. If not provided, the keys will be generated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssh_user": schema.StringAttribute{
								Description:         "sshUser is the username for outgoing SSH connections. Defaults to 'root'.",
								MarkdownDescription: "sshUser is the username for outgoing SSH connections. Defaults to 'root'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class_name": schema.StringAttribute{
								Description:         "storageClassName can be used to override the StorageClass of the PiT image.",
								MarkdownDescription: "storageClassName can be used to override the StorageClass of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_snapshot_class_name": schema.StringAttribute{
								Description:         "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								MarkdownDescription: "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rsync_tls": schema.SingleNestedAttribute{
						Description:         "rsyncTLS defines the configuration when using Rsync-based replication over TLS.",
						MarkdownDescription: "rsyncTLS defines the configuration when using Rsync-based replication over TLS.",
						Attributes: map[string]schema.Attribute{
							"access_modes": schema.ListAttribute{
								Description:         "accessModes can be used to override the accessModes of the PiT image.",
								MarkdownDescription: "accessModes can be used to override the accessModes of the PiT image.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"address": schema.StringAttribute{
								Description:         "address is the remote address to connect to for replication.",
								MarkdownDescription: "address is the remote address to connect to for replication.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capacity": schema.StringAttribute{
								Description:         "capacity can be used to override the capacity of the PiT image.",
								MarkdownDescription: "capacity can be used to override the capacity of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"copy_method": schema.StringAttribute{
								Description:         "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								MarkdownDescription: "copyMethod describes how a point-in-time (PiT) image of the source volume should be created.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Direct", "None", "Clone", "Snapshot"),
								},
							},

							"key_secret": schema.StringAttribute{
								Description:         "keySecret is the name of a Secret that contains the TLS pre-shared key to be used for authentication. If not provided, the key will be generated.",
								MarkdownDescription: "keySecret is the name of a Secret that contains the TLS pre-shared key to be used for authentication. If not provided, the key will be generated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_pod_labels": schema.MapAttribute{
								Description:         "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								MarkdownDescription: "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_resources": schema.SingleNestedAttribute{
								Description:         "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
								MarkdownDescription: "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
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

							"mover_security_context": schema.SingleNestedAttribute{
								Description:         "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								MarkdownDescription: "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								Attributes: map[string]schema.Attribute{
									"fs_group": schema.Int64Attribute{
										Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fs_group_change_policy": schema.StringAttribute{
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_group": schema.Int64Attribute{
										Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_non_root": schema.BoolAttribute{
										Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user": schema.Int64Attribute{
										Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"se_linux_options": schema.SingleNestedAttribute{
										Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"level": schema.StringAttribute{
												Description:         "Level is SELinux level label that applies to the container.",
												MarkdownDescription: "Level is SELinux level label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a SELinux role label that applies to the container.",
												MarkdownDescription: "Role is a SELinux role label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type is a SELinux type label that applies to the container.",
												MarkdownDescription: "Type is a SELinux type label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is a SELinux user label that applies to the container.",
												MarkdownDescription: "User is a SELinux user label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"seccomp_profile": schema.SingleNestedAttribute{
										Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"localhost_profile": schema.StringAttribute{
												Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"supplemental_groups": schema.ListAttribute{
										Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysctls": schema.ListNestedAttribute{
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of a property to set",
													MarkdownDescription: "Name of a property to set",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value of a property to set",
													MarkdownDescription: "Value of a property to set",
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

									"windows_options": schema.SingleNestedAttribute{
										Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										Attributes: map[string]schema.Attribute{
											"gmsa_credential_spec": schema.StringAttribute{
												Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gmsa_credential_spec_name": schema.StringAttribute{
												Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_process": schema.BoolAttribute{
												Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user_name": schema.StringAttribute{
												Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"mover_service_account": schema.StringAttribute{
								Description:         "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								MarkdownDescription: "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "port is the port to connect to for replication. Defaults to 8000.",
								MarkdownDescription: "port is the port to connect to for replication. Defaults to 8000.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(65535),
								},
							},

							"storage_class_name": schema.StringAttribute{
								Description:         "storageClassName can be used to override the StorageClass of the PiT image.",
								MarkdownDescription: "storageClassName can be used to override the StorageClass of the PiT image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_snapshot_class_name": schema.StringAttribute{
								Description:         "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								MarkdownDescription: "volumeSnapshotClassName can be used to specify the VSC to be used if copyMethod is Snapshot. If not set, the default VSC is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_pvc": schema.StringAttribute{
						Description:         "sourcePVC is the name of the PersistentVolumeClaim (PVC) to replicate.",
						MarkdownDescription: "sourcePVC is the name of the PersistentVolumeClaim (PVC) to replicate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"syncthing": schema.SingleNestedAttribute{
						Description:         "syncthing defines the configuration when using Syncthing-based replication.",
						MarkdownDescription: "syncthing defines the configuration when using Syncthing-based replication.",
						Attributes: map[string]schema.Attribute{
							"config_access_modes": schema.ListAttribute{
								Description:         "Used to set the accessModes of Syncthing config volume.",
								MarkdownDescription: "Used to set the accessModes of Syncthing config volume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_capacity": schema.StringAttribute{
								Description:         "Used to set the size of the Syncthing config volume.",
								MarkdownDescription: "Used to set the size of the Syncthing config volume.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_storage_class_name": schema.StringAttribute{
								Description:         "Used to set the StorageClass of the Syncthing config volume.",
								MarkdownDescription: "Used to set the StorageClass of the Syncthing config volume.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_pod_labels": schema.MapAttribute{
								Description:         "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								MarkdownDescription: "Labels that should be added to data mover pods These will be in addition to any labels that VolSync may add",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mover_resources": schema.SingleNestedAttribute{
								Description:         "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
								MarkdownDescription: "Resources represents compute resources required by the data mover container. Immutable. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ This should only be used by advanced users as this can result in a mover pod being unschedulable or crashing due to limited resources.",
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

							"mover_security_context": schema.SingleNestedAttribute{
								Description:         "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								MarkdownDescription: "MoverSecurityContext allows specifying the PodSecurityContext that will be used by the data mover",
								Attributes: map[string]schema.Attribute{
									"fs_group": schema.Int64Attribute{
										Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fs_group_change_policy": schema.StringAttribute{
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_group": schema.Int64Attribute{
										Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_non_root": schema.BoolAttribute{
										Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user": schema.Int64Attribute{
										Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"se_linux_options": schema.SingleNestedAttribute{
										Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"level": schema.StringAttribute{
												Description:         "Level is SELinux level label that applies to the container.",
												MarkdownDescription: "Level is SELinux level label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a SELinux role label that applies to the container.",
												MarkdownDescription: "Role is a SELinux role label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type is a SELinux type label that applies to the container.",
												MarkdownDescription: "Type is a SELinux type label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "User is a SELinux user label that applies to the container.",
												MarkdownDescription: "User is a SELinux user label that applies to the container.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"seccomp_profile": schema.SingleNestedAttribute{
										Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"localhost_profile": schema.StringAttribute{
												Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"supplemental_groups": schema.ListAttribute{
										Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysctls": schema.ListNestedAttribute{
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of a property to set",
													MarkdownDescription: "Name of a property to set",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value of a property to set",
													MarkdownDescription: "Value of a property to set",
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

									"windows_options": schema.SingleNestedAttribute{
										Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
										Attributes: map[string]schema.Attribute{
											"gmsa_credential_spec": schema.StringAttribute{
												Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gmsa_credential_spec_name": schema.StringAttribute{
												Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_process": schema.BoolAttribute{
												Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user_name": schema.StringAttribute{
												Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"mover_service_account": schema.StringAttribute{
								Description:         "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								MarkdownDescription: "MoverServiceAccount allows specifying the name of the service account that will be used by the data mover. This should only be used by advanced users who want to override the service account normally used by the mover. The service account needs to exist in the same namespace as this CR.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"peers": schema.ListNestedAttribute{
								Description:         "List of Syncthing peers to be connected for syncing",
								MarkdownDescription: "List of Syncthing peers to be connected for syncing",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "The peer's Syncthing ID.",
											MarkdownDescription: "The peer's Syncthing ID.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"address": schema.StringAttribute{
											Description:         "The peer's address that our Syncthing node will connect to.",
											MarkdownDescription: "The peer's address that our Syncthing node will connect to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"introducer": schema.BoolAttribute{
											Description:         "A flag that determines whether this peer should introduce us to other peers sharing this volume. It is HIGHLY recommended that two Syncthing peers do NOT set each other as introducers as you will have a difficult time disconnecting the two.",
											MarkdownDescription: "A flag that determines whether this peer should introduce us to other peers sharing this volume. It is HIGHLY recommended that two Syncthing peers do NOT set each other as introducers as you will have a difficult time disconnecting the two.",
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

							"service_type": schema.StringAttribute{
								Description:         "Type of service to be used when exposing the Syncthing peer",
								MarkdownDescription: "Type of service to be used when exposing the Syncthing peer",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"trigger": schema.SingleNestedAttribute{
						Description:         "trigger determines when the latest state of the volume will be captured (and potentially replicated to the destination).",
						MarkdownDescription: "trigger determines when the latest state of the volume will be captured (and potentially replicated to the destination).",
						Attributes: map[string]schema.Attribute{
							"manual": schema.StringAttribute{
								Description:         "manual is a string value that schedules a manual trigger. Once a sync completes then status.lastManualSync is set to the same string value. A consumer of a manual trigger should set spec.trigger.manual to a known value and then wait for lastManualSync to be updated by the operator to the same value, which means that the manual trigger will then pause and wait for further updates to the trigger.",
								MarkdownDescription: "manual is a string value that schedules a manual trigger. Once a sync completes then status.lastManualSync is set to the same string value. A consumer of a manual trigger should set spec.trigger.manual to a known value and then wait for lastManualSync to be updated by the operator to the same value, which means that the manual trigger will then pause and wait for further updates to the trigger.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schedule": schema.StringAttribute{
								Description:         "schedule is a cronspec (https://en.wikipedia.org/wiki/Cron#Overview) that can be used to schedule replication to occur at regular, time-based intervals. nolint:lll",
								MarkdownDescription: "schedule is a cronspec (https://en.wikipedia.org/wiki/Cron#Overview) that can be used to schedule replication to occur at regular, time-based intervals. nolint:lll",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(@(annually|yearly|monthly|weekly|daily|hourly))|((((\d+,)*\d+|(\d+(\/|-)\d+)|\*(\/\d+)?)\s?){5})$`), ""),
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

func (r *VolsyncBackubeReplicationSourceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_volsync_backube_replication_source_v1alpha1_manifest")

	var model VolsyncBackubeReplicationSourceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("volsync.backube/v1alpha1")
	model.Kind = pointer.String("ReplicationSource")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
