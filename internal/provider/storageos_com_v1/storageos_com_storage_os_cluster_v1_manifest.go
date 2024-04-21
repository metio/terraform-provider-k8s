/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storageos_com_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &StorageosComStorageOsclusterV1Manifest{}
)

func NewStorageosComStorageOsclusterV1Manifest() datasource.DataSource {
	return &StorageosComStorageOsclusterV1Manifest{}
}

type StorageosComStorageOsclusterV1Manifest struct{}

type StorageosComStorageOsclusterV1ManifestData struct {
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
		ContainerResources *struct {
			ApiManagerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"api_manager_container" json:"apiManagerContainer,omitempty"`
			CliContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"cli_container" json:"cliContainer,omitempty"`
			CsiExternalAttacherContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_external_attacher_container" json:"csiExternalAttacherContainer,omitempty"`
			CsiExternalProvisionerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_external_provisioner_container" json:"csiExternalProvisionerContainer,omitempty"`
			CsiExternalResizerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_external_resizer_container" json:"csiExternalResizerContainer,omitempty"`
			CsiExternalSnapshotterContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_external_snapshotter_container" json:"csiExternalSnapshotterContainer,omitempty"`
			CsiLivenessProbeContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_liveness_probe_container" json:"csiLivenessProbeContainer,omitempty"`
			CsiNodeDriverRegistrarContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"csi_node_driver_registrar_container" json:"csiNodeDriverRegistrarContainer,omitempty"`
			InitContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"init_container" json:"initContainer,omitempty"`
			KubeSchedulerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"kube_scheduler_container" json:"kubeSchedulerContainer,omitempty"`
			MetricsExporterContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"metrics_exporter_container" json:"metricsExporterContainer,omitempty"`
			NodeContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"node_container" json:"nodeContainer,omitempty"`
			NodeManagerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"node_manager_container" json:"nodeManagerContainer,omitempty"`
			PortalManagerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"portal_manager_container" json:"portalManagerContainer,omitempty"`
			SnapshotControllerContainer *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"snapshot_controller_container" json:"snapshotControllerContainer,omitempty"`
		} `tfsdk:"container_resources" json:"containerResources,omitempty"`
		Csi *struct {
			AttacherTimeout              *string `tfsdk:"attacher_timeout" json:"attacherTimeout,omitempty"`
			DeploymentStrategy           *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
			DeviceDir                    *string `tfsdk:"device_dir" json:"deviceDir,omitempty"`
			DriverRegisterationMode      *string `tfsdk:"driver_registeration_mode" json:"driverRegisterationMode,omitempty"`
			DriverRequiresAttachment     *string `tfsdk:"driver_requires_attachment" json:"driverRequiresAttachment,omitempty"`
			Enable                       *bool   `tfsdk:"enable" json:"enable,omitempty"`
			EnableControllerExpandCreds  *bool   `tfsdk:"enable_controller_expand_creds" json:"enableControllerExpandCreds,omitempty"`
			EnableControllerPublishCreds *bool   `tfsdk:"enable_controller_publish_creds" json:"enableControllerPublishCreds,omitempty"`
			EnableNodePublishCreds       *bool   `tfsdk:"enable_node_publish_creds" json:"enableNodePublishCreds,omitempty"`
			EnableProvisionCreds         *bool   `tfsdk:"enable_provision_creds" json:"enableProvisionCreds,omitempty"`
			Endpoint                     *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			KubeletDir                   *string `tfsdk:"kubelet_dir" json:"kubeletDir,omitempty"`
			KubeletRegistrationPath      *string `tfsdk:"kubelet_registration_path" json:"kubeletRegistrationPath,omitempty"`
			PluginDir                    *string `tfsdk:"plugin_dir" json:"pluginDir,omitempty"`
			ProvisionerTimeout           *string `tfsdk:"provisioner_timeout" json:"provisionerTimeout,omitempty"`
			ProvisionerWorkerCount       *int64  `tfsdk:"provisioner_worker_count" json:"provisionerWorkerCount,omitempty"`
			RegistrarSocketDir           *string `tfsdk:"registrar_socket_dir" json:"registrarSocketDir,omitempty"`
			RegistrationDir              *string `tfsdk:"registration_dir" json:"registrationDir,omitempty"`
			ResizerTimeout               *string `tfsdk:"resizer_timeout" json:"resizerTimeout,omitempty"`
			SnapshotterTimeout           *string `tfsdk:"snapshotter_timeout" json:"snapshotterTimeout,omitempty"`
			Version                      *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"csi" json:"csi,omitempty"`
		Debug               *bool              `tfsdk:"debug" json:"debug,omitempty"`
		DisableCLI          *bool              `tfsdk:"disable_cli" json:"disableCLI,omitempty"`
		DisableFencing      *bool              `tfsdk:"disable_fencing" json:"disableFencing,omitempty"`
		DisableScheduler    *bool              `tfsdk:"disable_scheduler" json:"disableScheduler,omitempty"`
		DisableTCMU         *bool              `tfsdk:"disable_tcmu" json:"disableTCMU,omitempty"`
		DisableTelemetry    *bool              `tfsdk:"disable_telemetry" json:"disableTelemetry,omitempty"`
		EnablePortalManager *bool              `tfsdk:"enable_portal_manager" json:"enablePortalManager,omitempty"`
		Environment         *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
		ForceTCMU           *bool              `tfsdk:"force_tcmu" json:"forceTCMU,omitempty"`
		Images              *struct {
			ApiManagerContainer                *string `tfsdk:"api_manager_container" json:"apiManagerContainer,omitempty"`
			CliContainer                       *string `tfsdk:"cli_container" json:"cliContainer,omitempty"`
			CsiClusterDriverRegistrarContainer *string `tfsdk:"csi_cluster_driver_registrar_container" json:"csiClusterDriverRegistrarContainer,omitempty"`
			CsiExternalAttacherContainer       *string `tfsdk:"csi_external_attacher_container" json:"csiExternalAttacherContainer,omitempty"`
			CsiExternalProvisionerContainer    *string `tfsdk:"csi_external_provisioner_container" json:"csiExternalProvisionerContainer,omitempty"`
			CsiExternalResizerContainer        *string `tfsdk:"csi_external_resizer_container" json:"csiExternalResizerContainer,omitempty"`
			CsiExternalSnapshotterContainer    *string `tfsdk:"csi_external_snapshotter_container" json:"csiExternalSnapshotterContainer,omitempty"`
			CsiLivenessProbeContainer          *string `tfsdk:"csi_liveness_probe_container" json:"csiLivenessProbeContainer,omitempty"`
			CsiNodeDriverRegistrarContainer    *string `tfsdk:"csi_node_driver_registrar_container" json:"csiNodeDriverRegistrarContainer,omitempty"`
			HyperkubeContainer                 *string `tfsdk:"hyperkube_container" json:"hyperkubeContainer,omitempty"`
			InitContainer                      *string `tfsdk:"init_container" json:"initContainer,omitempty"`
			KubeSchedulerContainer             *string `tfsdk:"kube_scheduler_container" json:"kubeSchedulerContainer,omitempty"`
			MetricsExporterContainer           *string `tfsdk:"metrics_exporter_container" json:"metricsExporterContainer,omitempty"`
			NfsContainer                       *string `tfsdk:"nfs_container" json:"nfsContainer,omitempty"`
			NodeContainer                      *string `tfsdk:"node_container" json:"nodeContainer,omitempty"`
			NodeGuardContainer                 *string `tfsdk:"node_guard_container" json:"nodeGuardContainer,omitempty"`
			NodeManagerContainer               *string `tfsdk:"node_manager_container" json:"nodeManagerContainer,omitempty"`
			PortalManagerContainer             *string `tfsdk:"portal_manager_container" json:"portalManagerContainer,omitempty"`
			SnapshotControllerContainer        *string `tfsdk:"snapshot_controller_container" json:"snapshotControllerContainer,omitempty"`
		} `tfsdk:"images" json:"images,omitempty"`
		Ingress *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Enable      *bool              `tfsdk:"enable" json:"enable,omitempty"`
			Hostname    *string            `tfsdk:"hostname" json:"hostname,omitempty"`
			Tls         *bool              `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Join      *string `tfsdk:"join" json:"join,omitempty"`
		K8sDistro *string `tfsdk:"k8s_distro" json:"k8sDistro,omitempty"`
		KvBackend *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Backend *string `tfsdk:"backend" json:"backend,omitempty"`
		} `tfsdk:"kv_backend" json:"kvBackend,omitempty"`
		Metrics *struct {
			DisabledCollectors *[]string `tfsdk:"disabled_collectors" json:"disabledCollectors,omitempty"`
			Enabled            *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			LogLevel           *string   `tfsdk:"log_level" json:"logLevel,omitempty"`
			Timeout            *int64    `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Namespace           *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		NodeManagerFeatures *map[string]string `tfsdk:"node_manager_features" json:"nodeManagerFeatures,omitempty"`
		NodeSelectorTerms   *[]struct {
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
		Pause     *bool `tfsdk:"pause" json:"pause,omitempty"`
		Resources *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		SecretRefName      *string `tfsdk:"secret_ref_name" json:"secretRefName,omitempty"`
		SecretRefNamespace *string `tfsdk:"secret_ref_namespace" json:"secretRefNamespace,omitempty"`
		Service            *struct {
			Annotations  *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			ExternalPort *int64             `tfsdk:"external_port" json:"externalPort,omitempty"`
			InternalPort *int64             `tfsdk:"internal_port" json:"internalPort,omitempty"`
			Name         *string            `tfsdk:"name" json:"name,omitempty"`
			Type         *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		SharedDir *string `tfsdk:"shared_dir" json:"sharedDir,omitempty"`
		Snapshots *struct {
			VolumeSnapshotClassName *string `tfsdk:"volume_snapshot_class_name" json:"volumeSnapshotClassName,omitempty"`
		} `tfsdk:"snapshots" json:"snapshots,omitempty"`
		StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		TlsEtcdSecretRefName      *string `tfsdk:"tls_etcd_secret_ref_name" json:"tlsEtcdSecretRefName,omitempty"`
		TlsEtcdSecretRefNamespace *string `tfsdk:"tls_etcd_secret_ref_namespace" json:"tlsEtcdSecretRefNamespace,omitempty"`
		Tolerations               *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StorageosComStorageOsclusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storageos_com_storage_os_cluster_v1_manifest"
}

func (r *StorageosComStorageOsclusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageOSCluster is the Schema for the storageosclusters API",
		MarkdownDescription: "StorageOSCluster is the Schema for the storageosclusters API",
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
				Description:         "StorageOSClusterSpec defines the desired state of StorageOSCluster",
				MarkdownDescription: "StorageOSClusterSpec defines the desired state of StorageOSCluster",
				Attributes: map[string]schema.Attribute{
					"container_resources": schema.SingleNestedAttribute{
						Description:         "ContainerResources is to set the resource requirements of each individual container managed by the operator.",
						MarkdownDescription: "ContainerResources is to set the resource requirements of each individual container managed by the operator.",
						Attributes: map[string]schema.Attribute{
							"api_manager_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"cli_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_external_attacher_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_external_provisioner_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_external_resizer_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_external_snapshotter_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_liveness_probe_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"csi_node_driver_registrar_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"init_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"kube_scheduler_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"metrics_exporter_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"node_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"node_manager_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"portal_manager_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"snapshot_controller_container": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"csi": schema.SingleNestedAttribute{
						Description:         "CSI defines the configurations for CSI.",
						MarkdownDescription: "CSI defines the configurations for CSI.",
						Attributes: map[string]schema.Attribute{
							"attacher_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deployment_strategy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device_dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"driver_registeration_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"driver_requires_attachment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_controller_expand_creds": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_controller_publish_creds": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_node_publish_creds": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_provision_creds": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubelet_dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubelet_registration_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provisioner_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provisioner_worker_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registrar_socket_dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_dir": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resizer_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshotter_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
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

					"debug": schema.BoolAttribute{
						Description:         "Debug is to set debug mode of the cluster.",
						MarkdownDescription: "Debug is to set debug mode of the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_cli": schema.BoolAttribute{
						Description:         "Disable StorageOS CLI deployment.",
						MarkdownDescription: "Disable StorageOS CLI deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_fencing": schema.BoolAttribute{
						Description:         "Disable Pod Fencing.  With StatefulSets, Pods are only re-scheduled if the Pod has been marked as killed.  In practice this means that failover of a StatefulSet pod is a manual operation.  By enabling Pod Fencing and setting the 'storageos.com/fenced=true' label on a Pod, StorageOS will enable automated Pod failover (by killing the application Pod on the failed node) if the following conditions exist:  - Pod fencing has not been explicitly disabled. - StorageOS has determined that the node the Pod is running on is offline.  StorageOS uses Gossip and TCP checks and will retry for 30 seconds.  At this point all volumes on the failed node are marked offline (irrespective of whether fencing is enabled) and volume failover starts. - The Pod has the label 'storageos.com/fenced=true' set. - The Pod has at least one StorageOS volume attached. - Each StorageOS volume has at least 1 healthy replica.  When Pod Fencing is disabled, StorageOS will not perform any interaction with Kubernetes when it detects that a node has gone offline. Additionally, the Kubernetes permissions required for Fencing will not be added to the StorageOS role. Deprecated: Not used any more, fencing is enabled/disabled by storageos.com/fenced label on pod.",
						MarkdownDescription: "Disable Pod Fencing.  With StatefulSets, Pods are only re-scheduled if the Pod has been marked as killed.  In practice this means that failover of a StatefulSet pod is a manual operation.  By enabling Pod Fencing and setting the 'storageos.com/fenced=true' label on a Pod, StorageOS will enable automated Pod failover (by killing the application Pod on the failed node) if the following conditions exist:  - Pod fencing has not been explicitly disabled. - StorageOS has determined that the node the Pod is running on is offline.  StorageOS uses Gossip and TCP checks and will retry for 30 seconds.  At this point all volumes on the failed node are marked offline (irrespective of whether fencing is enabled) and volume failover starts. - The Pod has the label 'storageos.com/fenced=true' set. - The Pod has at least one StorageOS volume attached. - Each StorageOS volume has at least 1 healthy replica.  When Pod Fencing is disabled, StorageOS will not perform any interaction with Kubernetes when it detects that a node has gone offline. Additionally, the Kubernetes permissions required for Fencing will not be added to the StorageOS role. Deprecated: Not used any more, fencing is enabled/disabled by storageos.com/fenced label on pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_scheduler": schema.BoolAttribute{
						Description:         "Disable StorageOS scheduler extender.",
						MarkdownDescription: "Disable StorageOS scheduler extender.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_tcmu": schema.BoolAttribute{
						Description:         "Disable TCMU can be set to true to disable the TCMU storage driver.  This is required when there are multiple storage systems running on the same node and you wish to avoid conflicts.  Only one TCMU-based storage system can run on a node at a time.  Disabling TCMU will degrade performance. Deprecated: Not used any more.",
						MarkdownDescription: "Disable TCMU can be set to true to disable the TCMU storage driver.  This is required when there are multiple storage systems running on the same node and you wish to avoid conflicts.  Only one TCMU-based storage system can run on a node at a time.  Disabling TCMU will degrade performance. Deprecated: Not used any more.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_telemetry": schema.BoolAttribute{
						Description:         "Disable Telemetry.",
						MarkdownDescription: "Disable Telemetry.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_portal_manager": schema.BoolAttribute{
						Description:         "EnablePortalManager enables Portal Manager.",
						MarkdownDescription: "EnablePortalManager enables Portal Manager.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"environment": schema.MapAttribute{
						Description:         "Environment contains environment variables that are passed to StorageOS.",
						MarkdownDescription: "Environment contains environment variables that are passed to StorageOS.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"force_tcmu": schema.BoolAttribute{
						Description:         "Force TCMU can be set to true to ensure that TCMU is enabled or cause StorageOS to abort startup.  At startup, StorageOS will automatically fallback to non-TCMU mode if another TCMU-based storage system is running on the node.  Since non-TCMU will degrade performance, this may not always be desired. Deprecated: Not used any more.",
						MarkdownDescription: "Force TCMU can be set to true to ensure that TCMU is enabled or cause StorageOS to abort startup.  At startup, StorageOS will automatically fallback to non-TCMU mode if another TCMU-based storage system is running on the node.  Since non-TCMU will degrade performance, this may not always be desired. Deprecated: Not used any more.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"images": schema.SingleNestedAttribute{
						Description:         "Images defines the various container images used in the cluster.",
						MarkdownDescription: "Images defines the various container images used in the cluster.",
						Attributes: map[string]schema.Attribute{
							"api_manager_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cli_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_cluster_driver_registrar_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_external_attacher_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_external_provisioner_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_external_resizer_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_external_snapshotter_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_liveness_probe_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"csi_node_driver_registrar_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hyperkube_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"init_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_scheduler_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_exporter_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nfs_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_guard_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_manager_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"portal_manager_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshot_controller_container": schema.StringAttribute{
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

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress defines the ingress configurations used in the cluster. Deprecated: Not used any more, please create your ingress for dashboard on your own.",
						MarkdownDescription: "Ingress defines the ingress configurations used in the cluster. Deprecated: Not used any more, please create your ingress for dashboard on your own.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.BoolAttribute{
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

					"join": schema.StringAttribute{
						Description:         "Join is the join token used for service discovery. Deprecated: Not used any more.",
						MarkdownDescription: "Join is the join token used for service discovery. Deprecated: Not used any more.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"k8s_distro": schema.StringAttribute{
						Description:         "K8sDistro is the name of the Kubernetes distribution where the operator is being deployed.  It should be in the format: 'name[-1.0]', where the version is optional and should only be appended if known.  Suitable names include: 'openshift', 'rancher', 'aks', 'gke', 'eks', or the deployment method if using upstream directly, e.g 'minishift' or 'kubeadm'.  Setting k8sDistro is optional, and will be used to simplify cluster configuration by setting appropriate defaults for the distribution.  The distribution information will also be included in the product telemetry (if enabled), to help focus development efforts.",
						MarkdownDescription: "K8sDistro is the name of the Kubernetes distribution where the operator is being deployed.  It should be in the format: 'name[-1.0]', where the version is optional and should only be appended if known.  Suitable names include: 'openshift', 'rancher', 'aks', 'gke', 'eks', or the deployment method if using upstream directly, e.g 'minishift' or 'kubeadm'.  Setting k8sDistro is optional, and will be used to simplify cluster configuration by setting appropriate defaults for the distribution.  The distribution information will also be included in the product telemetry (if enabled), to help focus development efforts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kv_backend": schema.SingleNestedAttribute{
						Description:         "KVBackend defines the key-value store backend used in the cluster.",
						MarkdownDescription: "KVBackend defines the key-value store backend used in the cluster.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"backend": schema.StringAttribute{
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

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics feature configuration.",
						MarkdownDescription: "Metrics feature configuration.",
						Attributes: map[string]schema.Attribute{
							"disabled_collectors": schema.ListAttribute{
								Description:         "DisabledCollectors is a list of collectors that shall be disabled. By default, all are enabled.",
								MarkdownDescription: "DisabledCollectors is a list of collectors that shall be disabled. By default, all are enabled.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Verbosity of log messages. Accepts go.uber.org/zap log levels.",
								MarkdownDescription: "Verbosity of log messages. Accepts go.uber.org/zap log levels.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("debug", "info", "warn", "error", "dpanic", "panic", "fatal"),
								},
							},

							"timeout": schema.Int64Attribute{
								Description:         "Timeout in seconds to serve metrics.",
								MarkdownDescription: "Timeout in seconds to serve metrics.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace is the kubernetes Namespace where storageos resources are provisioned. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						MarkdownDescription: "Namespace is the kubernetes Namespace where storageos resources are provisioned. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_manager_features": schema.MapAttribute{
						Description:         "Node manager feature list with optional configurations.",
						MarkdownDescription: "Node manager feature list with optional configurations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector_terms": schema.ListNestedAttribute{
						Description:         "NodeSelectorTerms is to set the placement of storageos pods using node affinity requiredDuringSchedulingIgnoredDuringExecution.",
						MarkdownDescription: "NodeSelectorTerms is to set the placement of storageos pods using node affinity requiredDuringSchedulingIgnoredDuringExecution.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pause": schema.BoolAttribute{
						Description:         "Pause is to pause the operator for the cluster. Deprecated: Not used any more, operator is always running.",
						MarkdownDescription: "Pause is to pause the operator for the cluster. Deprecated: Not used any more, operator is always running.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources is to set the resource requirements of the storageos containers. Deprecated: Set resource requests for individual containers via ContainerResources field in spec.",
						MarkdownDescription: "Resources is to set the resource requirements of the storageos containers. Deprecated: Set resource requests for individual containers via ContainerResources field in spec.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"secret_ref_name": schema.StringAttribute{
						Description:         "SecretRefName is the name of the secret object that contains all the sensitive cluster configurations.",
						MarkdownDescription: "SecretRefName is the name of the secret object that contains all the sensitive cluster configurations.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"secret_ref_namespace": schema.StringAttribute{
						Description:         "SecretRefNamespace is the namespace of the secret reference. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						MarkdownDescription: "SecretRefNamespace is the namespace of the secret reference. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "Service is the Service configuration for the cluster nodes.",
						MarkdownDescription: "Service is the Service configuration for the cluster nodes.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"internal_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"shared_dir": schema.StringAttribute{
						Description:         "SharedDir is the shared directory to be used when the kubelet is running in a container. Typically: '/var/lib/kubelet/plugins/kubernetes.io~storageos'. If not set, defaults will be used.",
						MarkdownDescription: "SharedDir is the shared directory to be used when the kubelet is running in a container. Typically: '/var/lib/kubelet/plugins/kubernetes.io~storageos'. If not set, defaults will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshots": schema.SingleNestedAttribute{
						Description:         "Snapshots feature configuration.",
						MarkdownDescription: "Snapshots feature configuration.",
						Attributes: map[string]schema.Attribute{
							"volume_snapshot_class_name": schema.StringAttribute{
								Description:         "VolumeSnapshotClassName is the name of default VolumeSnapshotClass created for StorageOS volumes.",
								MarkdownDescription: "VolumeSnapshotClassName is the name of default VolumeSnapshotClass created for StorageOS volumes.",
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
						Description:         "StorageClassName is the name of default StorageClass created for StorageOS volumes.",
						MarkdownDescription: "StorageClassName is the name of default StorageClass created for StorageOS volumes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_etcd_secret_ref_name": schema.StringAttribute{
						Description:         "TLSEtcdSecretRefName is the name of the secret object that contains the etcd TLS certs. This secret is shared with etcd, therefore it's not part of the main storageos secret.",
						MarkdownDescription: "TLSEtcdSecretRefName is the name of the secret object that contains the etcd TLS certs. This secret is shared with etcd, therefore it's not part of the main storageos secret.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_etcd_secret_ref_namespace": schema.StringAttribute{
						Description:         "TLSEtcdSecretRefNamespace is the namespace of the etcd TLS secret object. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						MarkdownDescription: "TLSEtcdSecretRefNamespace is the namespace of the etcd TLS secret object. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations is to set the placement of storageos pods using pod toleration.",
						MarkdownDescription: "Tolerations is to set the placement of storageos pods using pod toleration.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *StorageosComStorageOsclusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storageos_com_storage_os_cluster_v1_manifest")

	var model StorageosComStorageOsclusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("storageos.com/v1")
	model.Kind = pointer.String("StorageOSCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
