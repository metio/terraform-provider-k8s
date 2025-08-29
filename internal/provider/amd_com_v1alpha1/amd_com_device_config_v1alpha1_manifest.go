/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package amd_com_v1alpha1

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
	_ datasource.DataSource = &AmdComDeviceConfigV1Alpha1Manifest{}
)

func NewAmdComDeviceConfigV1Alpha1Manifest() datasource.DataSource {
	return &AmdComDeviceConfigV1Alpha1Manifest{}
}

type AmdComDeviceConfigV1Alpha1Manifest struct{}

type AmdComDeviceConfigV1Alpha1ManifestData struct {
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
		CommonConfig *struct {
			InitContainerImage *string `tfsdk:"init_container_image" json:"initContainerImage,omitempty"`
			UtilsContainer     *struct {
				Image               *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy     *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				ImageRegistrySecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			} `tfsdk:"utils_container" json:"utilsContainer,omitempty"`
		} `tfsdk:"common_config" json:"commonConfig,omitempty"`
		ConfigManager *struct {
			Config *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigManagerTolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"config_manager_tolerations" json:"configManagerTolerations,omitempty"`
			Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
			Image               *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy     *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageRegistrySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			Selector      *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			UpgradePolicy *struct {
				MaxUnavailable  *int64  `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				UpgradeStrategy *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
			} `tfsdk:"upgrade_policy" json:"upgradePolicy,omitempty"`
		} `tfsdk:"config_manager" json:"configManager,omitempty"`
		DevicePlugin *struct {
			DevicePluginArguments       *map[string]string `tfsdk:"device_plugin_arguments" json:"devicePluginArguments,omitempty"`
			DevicePluginImage           *string            `tfsdk:"device_plugin_image" json:"devicePluginImage,omitempty"`
			DevicePluginImagePullPolicy *string            `tfsdk:"device_plugin_image_pull_policy" json:"devicePluginImagePullPolicy,omitempty"`
			DevicePluginTolerations     *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"device_plugin_tolerations" json:"devicePluginTolerations,omitempty"`
			EnableNodeLabeller  *bool `tfsdk:"enable_node_labeller" json:"enableNodeLabeller,omitempty"`
			ImageRegistrySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			NodeLabellerArguments       *[]string `tfsdk:"node_labeller_arguments" json:"nodeLabellerArguments,omitempty"`
			NodeLabellerImage           *string   `tfsdk:"node_labeller_image" json:"nodeLabellerImage,omitempty"`
			NodeLabellerImagePullPolicy *string   `tfsdk:"node_labeller_image_pull_policy" json:"nodeLabellerImagePullPolicy,omitempty"`
			NodeLabellerTolerations     *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"node_labeller_tolerations" json:"nodeLabellerTolerations,omitempty"`
			UpgradePolicy *struct {
				MaxUnavailable  *int64  `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				UpgradeStrategy *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
			} `tfsdk:"upgrade_policy" json:"upgradePolicy,omitempty"`
		} `tfsdk:"device_plugin" json:"devicePlugin,omitempty"`
		Driver *struct {
			AmdgpuInstallerRepoURL *string `tfsdk:"amdgpu_installer_repo_url" json:"amdgpuInstallerRepoURL,omitempty"`
			Blacklist              *bool   `tfsdk:"blacklist" json:"blacklist,omitempty"`
			DriverType             *string `tfsdk:"driver_type" json:"driverType,omitempty"`
			Enable                 *bool   `tfsdk:"enable" json:"enable,omitempty"`
			Image                  *string `tfsdk:"image" json:"image,omitempty"`
			ImageBuild             *struct {
				BaseImageRegistry    *string `tfsdk:"base_image_registry" json:"baseImageRegistry,omitempty"`
				BaseImageRegistryTLS *struct {
					Insecure              *bool `tfsdk:"insecure" json:"insecure,omitempty"`
					InsecureSkipTLSVerify *bool `tfsdk:"insecure_skip_tls_verify" json:"insecureSkipTLSVerify,omitempty"`
				} `tfsdk:"base_image_registry_tls" json:"baseImageRegistryTLS,omitempty"`
			} `tfsdk:"image_build" json:"imageBuild,omitempty"`
			ImageRegistrySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			ImageRegistryTLS *struct {
				Insecure              *bool `tfsdk:"insecure" json:"insecure,omitempty"`
				InsecureSkipTLSVerify *bool `tfsdk:"insecure_skip_tls_verify" json:"insecureSkipTLSVerify,omitempty"`
			} `tfsdk:"image_registry_tls" json:"imageRegistryTLS,omitempty"`
			ImageSign *struct {
				CertSecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"cert_secret" json:"certSecret,omitempty"`
				KeySecret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"key_secret" json:"keySecret,omitempty"`
			} `tfsdk:"image_sign" json:"imageSign,omitempty"`
			KernelModuleConfig *struct {
				LoadArgs   *[]string `tfsdk:"load_args" json:"loadArgs,omitempty"`
				Parameters *[]string `tfsdk:"parameters" json:"parameters,omitempty"`
				UnloadArgs *[]string `tfsdk:"unload_args" json:"unloadArgs,omitempty"`
			} `tfsdk:"kernel_module_config" json:"kernelModuleConfig,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UpgradePolicy *struct {
				Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
				MaxParallelUpgrades *int64  `tfsdk:"max_parallel_upgrades" json:"maxParallelUpgrades,omitempty"`
				MaxUnavailableNodes *string `tfsdk:"max_unavailable_nodes" json:"maxUnavailableNodes,omitempty"`
				NodeDrainPolicy     *struct {
					Force              *bool  `tfsdk:"force" json:"force,omitempty"`
					GracePeriodSeconds *int64 `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
					TimeoutSeconds     *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"node_drain_policy" json:"nodeDrainPolicy,omitempty"`
				PodDeletionPolicy *struct {
					Force              *bool  `tfsdk:"force" json:"force,omitempty"`
					GracePeriodSeconds *int64 `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
					TimeoutSeconds     *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"pod_deletion_policy" json:"podDeletionPolicy,omitempty"`
				RebootRequired *bool `tfsdk:"reboot_required" json:"rebootRequired,omitempty"`
			} `tfsdk:"upgrade_policy" json:"upgradePolicy,omitempty"`
			Version    *string `tfsdk:"version" json:"version,omitempty"`
			VfioConfig *struct {
				DeviceIDs *[]string `tfsdk:"device_i_ds" json:"deviceIDs,omitempty"`
			} `tfsdk:"vfio_config" json:"vfioConfig,omitempty"`
		} `tfsdk:"driver" json:"driver,omitempty"`
		MetricsExporter *struct {
			Config *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
			Image               *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy     *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageRegistrySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			NodePort                 *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
			PodResourceAPISocketPath *string `tfsdk:"pod_resource_api_socket_path" json:"podResourceAPISocketPath,omitempty"`
			Port                     *int64  `tfsdk:"port" json:"port,omitempty"`
			Prometheus               *struct {
				ServiceMonitor *struct {
					AttachMetadata *struct {
						Node *bool `tfsdk:"node" json:"node,omitempty"`
					} `tfsdk:"attach_metadata" json:"attachMetadata,omitempty"`
					Authorization *struct {
						Credentials *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"credentials" json:"credentials,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"authorization" json:"authorization,omitempty"`
					BearerTokenFile   *string            `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
					Enable            *bool              `tfsdk:"enable" json:"enable,omitempty"`
					HonorLabels       *bool              `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
					HonorTimestamps   *bool              `tfsdk:"honor_timestamps" json:"honorTimestamps,omitempty"`
					Interval          *string            `tfsdk:"interval" json:"interval,omitempty"`
					Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					MetricRelabelings *[]struct {
						Action       *string   `tfsdk:"action" json:"action,omitempty"`
						Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
					Relabelings *[]struct {
						Action       *string   `tfsdk:"action" json:"action,omitempty"`
						Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"relabelings" json:"relabelings,omitempty"`
					TlsConfig *struct {
						Ca *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"ca" json:"ca,omitempty"`
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"cert" json:"cert,omitempty"`
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
			RbacConfig *struct {
				ClientCAConfigMap *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_ca_config_map" json:"clientCAConfigMap,omitempty"`
				DisableHttps *bool   `tfsdk:"disable_https" json:"disableHttps,omitempty"`
				Enable       *bool   `tfsdk:"enable" json:"enable,omitempty"`
				Image        *string `tfsdk:"image" json:"image,omitempty"`
				Secret       *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				StaticAuthorization *struct {
					ClientName *string `tfsdk:"client_name" json:"clientName,omitempty"`
					Enable     *bool   `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"static_authorization" json:"staticAuthorization,omitempty"`
			} `tfsdk:"rbac_config" json:"rbacConfig,omitempty"`
			Selector    *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UpgradePolicy *struct {
				MaxUnavailable  *int64  `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				UpgradeStrategy *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
			} `tfsdk:"upgrade_policy" json:"upgradePolicy,omitempty"`
		} `tfsdk:"metrics_exporter" json:"metricsExporter,omitempty"`
		RemediationWorkflow *struct {
			ConditionalWorkflows *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"conditional_workflows" json:"conditionalWorkflows,omitempty"`
			Enable                *bool  `tfsdk:"enable" json:"enable,omitempty"`
			TtlForFailedWorkflows *int64 `tfsdk:"ttl_for_failed_workflows" json:"ttlForFailedWorkflows,omitempty"`
		} `tfsdk:"remediation_workflow" json:"remediationWorkflow,omitempty"`
		Selector   *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
		TestRunner *struct {
			Config *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			Enable              *bool   `tfsdk:"enable" json:"enable,omitempty"`
			Image               *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy     *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageRegistrySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_registry_secret" json:"imageRegistrySecret,omitempty"`
			LogsLocation *struct {
				HostPath          *string `tfsdk:"host_path" json:"hostPath,omitempty"`
				LogsExportSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"logs_export_secrets" json:"logsExportSecrets,omitempty"`
				MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			} `tfsdk:"logs_location" json:"logsLocation,omitempty"`
			Selector    *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UpgradePolicy *struct {
				MaxUnavailable  *int64  `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				UpgradeStrategy *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
			} `tfsdk:"upgrade_policy" json:"upgradePolicy,omitempty"`
		} `tfsdk:"test_runner" json:"testRunner,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AmdComDeviceConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_amd_com_device_config_v1alpha1_manifest"
}

func (r *AmdComDeviceConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DeviceConfig describes how to enable AMD GPU device",
		MarkdownDescription: "DeviceConfig describes how to enable AMD GPU device",
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
				Description:         "DeviceConfigSpec describes how the AMD GPU operator should enable AMD GPU device for customer's use.",
				MarkdownDescription: "DeviceConfigSpec describes how the AMD GPU operator should enable AMD GPU device for customer's use.",
				Attributes: map[string]schema.Attribute{
					"common_config": schema.SingleNestedAttribute{
						Description:         "common config",
						MarkdownDescription: "common config",
						Attributes: map[string]schema.Attribute{
							"init_container_image": schema.StringAttribute{
								Description:         "InitContainerImage is being used for the operands pods, i.e. metrics exporter, test runner, device plugin, device config manager and node labeller",
								MarkdownDescription: "InitContainerImage is being used for the operands pods, i.e. metrics exporter, test runner, device plugin, device config manager and node labeller",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"utils_container": schema.SingleNestedAttribute{
								Description:         "UtilsContainer contains parameters to configure operator's utils container",
								MarkdownDescription: "UtilsContainer contains parameters to configure operator's utils container",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "Image is the image of utils container",
										MarkdownDescription: "Image is the image of utils container",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
										},
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "image pull policy for utils container",
										MarkdownDescription: "image pull policy for utils container",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
										},
									},

									"image_registry_secret": schema.SingleNestedAttribute{
										Description:         "secret used for pull utils container image",
										MarkdownDescription: "secret used for pull utils container image",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

					"config_manager": schema.SingleNestedAttribute{
						Description:         "config manager",
						MarkdownDescription: "config manager",
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Description:         "config map to customize the config for config manager, if not specified default config will be applied",
								MarkdownDescription: "config map to customize the config for config manager, if not specified default config will be applied",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_manager_tolerations": schema.ListNestedAttribute{
								Description:         "tolerations for the device config manager DaemonSet",
								MarkdownDescription: "tolerations for the device config manager DaemonSet",
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

							"enable": schema.BoolAttribute{
								Description:         "enable config manager, disabled by default",
								MarkdownDescription: "enable config manager, disabled by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "config manager image",
								MarkdownDescription: "config manager image",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "image pull policy for config manager",
								MarkdownDescription: "image pull policy for config manager",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"image_registry_secret": schema.SingleNestedAttribute{
								Description:         "config manager image registry secret used to pull/push images",
								MarkdownDescription: "config manager image registry secret used to pull/push images",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": schema.MapAttribute{
								Description:         "Selector describes on which nodes to enable config manager",
								MarkdownDescription: "Selector describes on which nodes to enable config manager",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"upgrade_policy": schema.SingleNestedAttribute{
								Description:         "upgrade policy for config manager daemonset",
								MarkdownDescription: "upgrade policy for config manager daemonset",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.Int64Attribute{
										Description:         "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										MarkdownDescription: "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"upgrade_strategy": schema.StringAttribute{
										Description:         "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										MarkdownDescription: "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("RollingUpdate", "OnDelete"),
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

					"device_plugin": schema.SingleNestedAttribute{
						Description:         "device plugin",
						MarkdownDescription: "device plugin",
						Attributes: map[string]schema.Attribute{
							"device_plugin_arguments": schema.MapAttribute{
								Description:         "device plugin arguments is used to pass supported flags and their values while starting device plugin daemonset supported flag values: {'resource_naming_strategy': {'single', 'mixed'}}",
								MarkdownDescription: "device plugin arguments is used to pass supported flags and their values while starting device plugin daemonset supported flag values: {'resource_naming_strategy': {'single', 'mixed'}}",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device_plugin_image": schema.StringAttribute{
								Description:         "device plugin image",
								MarkdownDescription: "device plugin image",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"device_plugin_image_pull_policy": schema.StringAttribute{
								Description:         "image pull policy for device plugin",
								MarkdownDescription: "image pull policy for device plugin",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"device_plugin_tolerations": schema.ListNestedAttribute{
								Description:         "tolerations for the device plugin DaemonSet",
								MarkdownDescription: "tolerations for the device plugin DaemonSet",
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

							"enable_node_labeller": schema.BoolAttribute{
								Description:         "enable or disable the node labeller",
								MarkdownDescription: "enable or disable the node labeller",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_registry_secret": schema.SingleNestedAttribute{
								Description:         "node labeller image registry secret used to pull/push images",
								MarkdownDescription: "node labeller image registry secret used to pull/push images",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_labeller_arguments": schema.ListAttribute{
								Description:         "node labeller arguments is used to pass supported labels while starting node labeller daemonset some flags are enabled by default as they are applicable and bare minimum for all setups and are supported in all versions of node labeller default flags: {'vram', 'cu-count', 'simd-count', 'device-id', 'family', 'product-name', 'driver-version'} supported flags: {'compute-memory-partition', 'compute-partitioning-supported', 'memory-partitioning-supported'}",
								MarkdownDescription: "node labeller arguments is used to pass supported labels while starting node labeller daemonset some flags are enabled by default as they are applicable and bare minimum for all setups and are supported in all versions of node labeller default flags: {'vram', 'cu-count', 'simd-count', 'device-id', 'family', 'product-name', 'driver-version'} supported flags: {'compute-memory-partition', 'compute-partitioning-supported', 'memory-partitioning-supported'}",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_labeller_image": schema.StringAttribute{
								Description:         "node labeller image",
								MarkdownDescription: "node labeller image",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"node_labeller_image_pull_policy": schema.StringAttribute{
								Description:         "image pull policy for node labeller",
								MarkdownDescription: "image pull policy for node labeller",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"node_labeller_tolerations": schema.ListNestedAttribute{
								Description:         "tolerations for the node labeller DaemonSet",
								MarkdownDescription: "tolerations for the node labeller DaemonSet",
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

							"upgrade_policy": schema.SingleNestedAttribute{
								Description:         "upgrade policy for device plugin and node labeller daemons",
								MarkdownDescription: "upgrade policy for device plugin and node labeller daemons",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.Int64Attribute{
										Description:         "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										MarkdownDescription: "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"upgrade_strategy": schema.StringAttribute{
										Description:         "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										MarkdownDescription: "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("RollingUpdate", "OnDelete"),
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

					"driver": schema.SingleNestedAttribute{
						Description:         "driver",
						MarkdownDescription: "driver",
						Attributes: map[string]schema.Attribute{
							"amdgpu_installer_repo_url": schema.StringAttribute{
								Description:         "radeon repo URL for fetching amdgpu installer if building driver image on the fly installer URL is https://repo.radeon.com/amdgpu-install by default",
								MarkdownDescription: "radeon repo URL for fetching amdgpu installer if building driver image on the fly installer URL is https://repo.radeon.com/amdgpu-install by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"blacklist": schema.BoolAttribute{
								Description:         "blacklist amdgpu drivers on the host. Node reboot is required to apply the baclklist on the worker nodes. Not working for OpenShift cluster. OpenShift users please use the Machine Config Operator (MCO) resource to configure amdgpu blacklist. Example MCO resource is available at https://instinct.docs.amd.com/projects/gpu-operator/en/latest/installation/openshift-olm.html#create-blacklist-for-installing-out-of-tree-kernel-module",
								MarkdownDescription: "blacklist amdgpu drivers on the host. Node reboot is required to apply the baclklist on the worker nodes. Not working for OpenShift cluster. OpenShift users please use the Machine Config Operator (MCO) resource to configure amdgpu blacklist. Example MCO resource is available at https://instinct.docs.amd.com/projects/gpu-operator/en/latest/installation/openshift-olm.html#create-blacklist-for-installing-out-of-tree-kernel-module",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"driver_type": schema.StringAttribute{
								Description:         "specify the type of driver (container/vf-passthrough/pf-passthrough) to install on the worker node. default value is container. container: normal amdgpu-dkms driver for Bare Metal GPU nodes or guest VM. vf-passthrough: MxGPU GIM driver on the host machine to generate VF, then mount VF to vfio-pci pf-passthrough: directly mount PF device to vfio-pci",
								MarkdownDescription: "specify the type of driver (container/vf-passthrough/pf-passthrough) to install on the worker node. default value is container. container: normal amdgpu-dkms driver for Bare Metal GPU nodes or guest VM. vf-passthrough: MxGPU GIM driver on the host machine to generate VF, then mount VF to vfio-pci pf-passthrough: directly mount PF device to vfio-pci",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("container", "vf-passthrough", "pf-passthrough"),
								},
							},

							"enable": schema.BoolAttribute{
								Description:         "enable driver install. default value is true. disable is for skipping driver install/uninstall for dryrun or using in-tree amdgpu kernel module",
								MarkdownDescription: "enable driver install. default value is true. disable is for skipping driver install/uninstall for dryrun or using in-tree amdgpu kernel module",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "defines image that includes drivers and firmware blobs, don't include tag since it will be fully managed by operator for vanilla k8s the default value is image-registry:5000/$MOD_NAMESPACE/amdgpu_kmod for OpenShift the default value is image-registry.openshift-image-registry.svc:5000/$MOD_NAMESPACE/amdgpu_kmod image tag will be in the format of <linux distro>-<release version>-<kernel version>-<driver version> example tag is coreos-416.94-5.14.0-427.28.1.el9_4.x86_64-6.2.2 and ubuntu-22.04-5.15.0-94-generic-6.1.3 NOTE: Updating the driver image repository is not supported. Please delete the existing DeviceConfig and create a new one with the updated image repository",
								MarkdownDescription: "defines image that includes drivers and firmware blobs, don't include tag since it will be fully managed by operator for vanilla k8s the default value is image-registry:5000/$MOD_NAMESPACE/amdgpu_kmod for OpenShift the default value is image-registry.openshift-image-registry.svc:5000/$MOD_NAMESPACE/amdgpu_kmod image tag will be in the format of <linux distro>-<release version>-<kernel version>-<driver version> example tag is coreos-416.94-5.14.0-427.28.1.el9_4.x86_64-6.2.2 and ubuntu-22.04-5.15.0-94-generic-6.1.3 NOTE: Updating the driver image repository is not supported. Please delete the existing DeviceConfig and create a new one with the updated image repository",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[$a-zA-Z0-9_]+(?:[._-][$a-zA-Z0-9_]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"image_build": schema.SingleNestedAttribute{
								Description:         "image build configs",
								MarkdownDescription: "image build configs",
								Attributes: map[string]schema.Attribute{
									"base_image_registry": schema.StringAttribute{
										Description:         "image registry to fetch base image for building driver image, default value is docker.io, the builder will search for corresponding OS base image from given registry e.g. if your worker node is using Ubuntu 22.04, by default the base image would be docker.io/ubuntu:22.04 NOTE: this field won't apply for OpenShift since OpenShift is using its own DriverToolKit image to build driver image",
										MarkdownDescription: "image registry to fetch base image for building driver image, default value is docker.io, the builder will search for corresponding OS base image from given registry e.g. if your worker node is using Ubuntu 22.04, by default the base image would be docker.io/ubuntu:22.04 NOTE: this field won't apply for OpenShift since OpenShift is using its own DriverToolKit image to build driver image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"base_image_registry_tls": schema.SingleNestedAttribute{
										Description:         "TLS settings for fetching base image",
										MarkdownDescription: "TLS settings for fetching base image",
										Attributes: map[string]schema.Attribute{
											"insecure": schema.BoolAttribute{
												Description:         "If true, check if the container image already exists using plain HTTP.",
												MarkdownDescription: "If true, check if the container image already exists using plain HTTP.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_tls_verify": schema.BoolAttribute{
												Description:         "If true, skip any TLS server certificate validation",
												MarkdownDescription: "If true, skip any TLS server certificate validation",
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

							"image_registry_secret": schema.SingleNestedAttribute{
								Description:         "secrets used for pull/push images from/to private registry specified in driversImage",
								MarkdownDescription: "secrets used for pull/push images from/to private registry specified in driversImage",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_registry_tls": schema.SingleNestedAttribute{
								Description:         "driver image registry TLS setting for the container image",
								MarkdownDescription: "driver image registry TLS setting for the container image",
								Attributes: map[string]schema.Attribute{
									"insecure": schema.BoolAttribute{
										Description:         "If true, check if the container image already exists using plain HTTP.",
										MarkdownDescription: "If true, check if the container image already exists using plain HTTP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_tls_verify": schema.BoolAttribute{
										Description:         "If true, skip any TLS server certificate validation",
										MarkdownDescription: "If true, skip any TLS server certificate validation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_sign": schema.SingleNestedAttribute{
								Description:         "image signing config to sign the driver image when building driver image on the fly image signing is required for installing driver on secure boot enabled system",
								MarkdownDescription: "image signing config to sign the driver image when building driver image on the fly image signing is required for installing driver on secure boot enabled system",
								Attributes: map[string]schema.Attribute{
									"cert_secret": schema.SingleNestedAttribute{
										Description:         "ImageSignCertSecret the public key used to sign kernel modules within image necessary for secure boot enabled system",
										MarkdownDescription: "ImageSignCertSecret the public key used to sign kernel modules within image necessary for secure boot enabled system",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": schema.SingleNestedAttribute{
										Description:         "ImageSignKeySecret the private key used to sign kernel modules within image necessary for secure boot enabled system",
										MarkdownDescription: "ImageSignKeySecret the private key used to sign kernel modules within image necessary for secure boot enabled system",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"kernel_module_config": schema.SingleNestedAttribute{
								Description:         "advanced arguments, parameters and more configs to manage tne driver",
								MarkdownDescription: "advanced arguments, parameters and more configs to manage tne driver",
								Attributes: map[string]schema.Attribute{
									"load_args": schema.ListAttribute{
										Description:         "LoadArg are the arguments when modprobe is executed to load the kernel module. The command will be 'modprobe ${Args} module_name'.",
										MarkdownDescription: "LoadArg are the arguments when modprobe is executed to load the kernel module. The command will be 'modprobe ${Args} module_name'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parameters": schema.ListAttribute{
										Description:         "Parameters is being used for modprobe commands. The command will be 'modprobe ${Args} module_name ${Parameters}'.",
										MarkdownDescription: "Parameters is being used for modprobe commands. The command will be 'modprobe ${Args} module_name ${Parameters}'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unload_args": schema.ListAttribute{
										Description:         "UnloadArg are the arguments when modprobe is executed to unload the kernel module. The command will be 'modprobe -r ${Args} module_name'.",
										MarkdownDescription: "UnloadArg are the arguments when modprobe is executed to unload the kernel module. The command will be 'modprobe -r ${Args} module_name'.",
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
								Description:         "tolerations for kmm module object",
								MarkdownDescription: "tolerations for kmm module object",
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

							"upgrade_policy": schema.SingleNestedAttribute{
								Description:         "policy to upgrade the drivers",
								MarkdownDescription: "policy to upgrade the drivers",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "enable upgrade policy, disabled by default If disabled, user has to manually upgrade all the nodes.",
										MarkdownDescription: "enable upgrade policy, disabled by default If disabled, user has to manually upgrade all the nodes.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_parallel_upgrades": schema.Int64Attribute{
										Description:         "MaxParallelUpgrades indicates how many nodes can be upgraded in parallel 0 means no limit, all nodes will be upgraded in parallel",
										MarkdownDescription: "MaxParallelUpgrades indicates how many nodes can be upgraded in parallel 0 means no limit, all nodes will be upgraded in parallel",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"max_unavailable_nodes": schema.StringAttribute{
										Description:         "MaxUnavailableNodes indicates maximum number of nodes that can be in a failed upgrade state beyond which upgrades will stop to keep cluster at a minimal healthy state Value can be an integer (ex: 2) which would mean atmost 2 nodes can be in failed state after which new upgrades will not start. Or it can be a percentage string(ex: '50%') from which absolute number will be calculated and round up",
										MarkdownDescription: "MaxUnavailableNodes indicates maximum number of nodes that can be in a failed upgrade state beyond which upgrades will stop to keep cluster at a minimal healthy state Value can be an integer (ex: 2) which would mean atmost 2 nodes can be in failed state after which new upgrades will not start. Or it can be a percentage string(ex: '50%') from which absolute number will be calculated and round up",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_drain_policy": schema.SingleNestedAttribute{
										Description:         "Node draining policy",
										MarkdownDescription: "Node draining policy",
										Attributes: map[string]schema.Attribute{
											"force": schema.BoolAttribute{
												Description:         "Force indicates if force draining is allowed",
												MarkdownDescription: "Force indicates if force draining is allowed",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grace_period_seconds": schema.Int64Attribute{
												Description:         "GracePeriodSeconds indicates the time kubernetes waits for a pod to shut down gracefully after receiving a termination signal",
												MarkdownDescription: "GracePeriodSeconds indicates the time kubernetes waits for a pod to shut down gracefully after receiving a termination signal",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "TimeoutSecond specifies the length of time in seconds to wait before giving up drain, zero means infinite",
												MarkdownDescription: "TimeoutSecond specifies the length of time in seconds to wait before giving up drain, zero means infinite",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_deletion_policy": schema.SingleNestedAttribute{
										Description:         "Pod Deletion policy. If both NodeDrainPolicy and PodDeletionPolicy config is available, NodeDrainPolicy(if enabled) will take precedence.",
										MarkdownDescription: "Pod Deletion policy. If both NodeDrainPolicy and PodDeletionPolicy config is available, NodeDrainPolicy(if enabled) will take precedence.",
										Attributes: map[string]schema.Attribute{
											"force": schema.BoolAttribute{
												Description:         "Force indicates if force deletion is allowed",
												MarkdownDescription: "Force indicates if force deletion is allowed",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grace_period_seconds": schema.Int64Attribute{
												Description:         "GracePeriodSeconds indicates the time kubernetes waits for a pod to shut down gracefully after receiving a termination signal",
												MarkdownDescription: "GracePeriodSeconds indicates the time kubernetes waits for a pod to shut down gracefully after receiving a termination signal",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "TimeoutSecond specifies the length of time in seconds to wait before giving up on pod deletion, zero means infinite",
												MarkdownDescription: "TimeoutSecond specifies the length of time in seconds to wait before giving up on pod deletion, zero means infinite",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"reboot_required": schema.BoolAttribute{
										Description:         "reboot between driver upgrades, enabled by default, if enabled spec.commonConfig.utilsContainer will be used to perform reboot on worker nodes",
										MarkdownDescription: "reboot between driver upgrades, enabled by default, if enabled spec.commonConfig.utilsContainer will be used to perform reboot on worker nodes",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": schema.StringAttribute{
								Description:         "version of the drivers source code, can be used as part of image of dockerfile source image default value for different OS is: ubuntu: 6.1.3, coreOS: 6.2.2",
								MarkdownDescription: "version of the drivers source code, can be used as part of image of dockerfile source image default value for different OS is: ubuntu: 6.1.3, coreOS: 6.2.2",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vfio_config": schema.SingleNestedAttribute{
								Description:         "vfio config specify the specific configs for binding PCI devices to vfio-pci kernel module, applies for driver type vf-passthrough and pf-passthrough",
								MarkdownDescription: "vfio config specify the specific configs for binding PCI devices to vfio-pci kernel module, applies for driver type vf-passthrough and pf-passthrough",
								Attributes: map[string]schema.Attribute{
									"device_i_ds": schema.ListAttribute{
										Description:         "list of PCI device IDs to load into vfio-pci driver. default is the list of AMD GPU PF/VF PCI device IDs based on driver type vf-passthrough/pf-passthrough.",
										MarkdownDescription: "list of PCI device IDs to load into vfio-pci driver. default is the list of AMD GPU PF/VF PCI device IDs based on driver type vf-passthrough/pf-passthrough.",
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

					"metrics_exporter": schema.SingleNestedAttribute{
						Description:         "metrics exporter",
						MarkdownDescription: "metrics exporter",
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Description:         "optional configuration for metrics",
								MarkdownDescription: "optional configuration for metrics",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the configMap that defines the list of metrics default list:[]",
										MarkdownDescription: "Name of the configMap that defines the list of metrics default list:[]",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "enable metrics exporter, disabled by default",
								MarkdownDescription: "enable metrics exporter, disabled by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "metrics exporter image",
								MarkdownDescription: "metrics exporter image",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "image pull policy for metrics exporter",
								MarkdownDescription: "image pull policy for metrics exporter",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"image_registry_secret": schema.SingleNestedAttribute{
								Description:         "metrics exporter image registry secret used to pull/push images",
								MarkdownDescription: "metrics exporter image registry secret used to pull/push images",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_port": schema.Int64Attribute{
								Description:         "NodePort is the external port for pulling metrics from outside the cluster, in the range 30000-32767 (assigned automatically by default)",
								MarkdownDescription: "NodePort is the external port for pulling metrics from outside the cluster, in the range 30000-32767 (assigned automatically by default)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(30000),
									int64validator.AtMost(32767),
								},
							},

							"pod_resource_api_socket_path": schema.StringAttribute{
								Description:         "Set the host path for pod-resource kubelet.socket, vanila kubernetes path is /var/lib/kubelet/pod-resources microk8s path is /var/snap/microk8s/common/var/lib/kubelet/pod-resources/ path is an absolute unix path that allows a trailing slash",
								MarkdownDescription: "Set the host path for pod-resource kubelet.socket, vanila kubernetes path is /var/lib/kubelet/pod-resources microk8s path is /var/snap/microk8s/common/var/lib/kubelet/pod-resources/ path is an absolute unix path that allows a trailing slash",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(/[^/\0]+)*(/)?$`), ""),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the internal port used for in-cluster and node access to pull metrics from the metrics-exporter (default 5000).",
								MarkdownDescription: "Port is the internal port used for in-cluster and node access to pull metrics from the metrics-exporter (default 5000).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prometheus": schema.SingleNestedAttribute{
								Description:         "Prometheus configuration for metrics exporter",
								MarkdownDescription: "Prometheus configuration for metrics exporter",
								Attributes: map[string]schema.Attribute{
									"service_monitor": schema.SingleNestedAttribute{
										Description:         "ServiceMonitor configuration for Prometheus integration",
										MarkdownDescription: "ServiceMonitor configuration for Prometheus integration",
										Attributes: map[string]schema.Attribute{
											"attach_metadata": schema.SingleNestedAttribute{
												Description:         "AttachMetadata defines if Prometheus should attach node metadata to the target",
												MarkdownDescription: "AttachMetadata defines if Prometheus should attach node metadata to the target",
												Attributes: map[string]schema.Attribute{
													"node": schema.BoolAttribute{
														Description:         "When set to true, Prometheus attaches node metadata to the discovered targets. The Prometheus service account must have the 'list' and 'watch' permissions on the 'Nodes' objects.",
														MarkdownDescription: "When set to true, Prometheus attaches node metadata to the discovered targets. The Prometheus service account must have the 'list' and 'watch' permissions on the 'Nodes' objects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"authorization": schema.SingleNestedAttribute{
												Description:         "Optional Prometheus authorization configuration for accessing the endpoint",
												MarkdownDescription: "Optional Prometheus authorization configuration for accessing the endpoint",
												Attributes: map[string]schema.Attribute{
													"credentials": schema.SingleNestedAttribute{
														Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
														MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from. Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

													"type": schema.StringAttribute{
														Description:         "Defines the authentication type. The value is case-insensitive. 'Basic' is not a supported value. Default: 'Bearer'",
														MarkdownDescription: "Defines the authentication type. The value is case-insensitive. 'Basic' is not a supported value. Default: 'Bearer'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_file": schema.StringAttribute{
												Description:         "Path to bearer token file to be used by Prometheus (e.g., service account token path) Deprecated: Use Authorization instead. This field is kept for backward compatibility.",
												MarkdownDescription: "Path to bearer token file to be used by Prometheus (e.g., service account token path) Deprecated: Use Authorization instead. This field is kept for backward compatibility.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable or disable ServiceMonitor creation (default false)",
												MarkdownDescription: "Enable or disable ServiceMonitor creation (default false)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"honor_labels": schema.BoolAttribute{
												Description:         "HonorLabels chooses the metric's labels on collisions with target labels (default true)",
												MarkdownDescription: "HonorLabels chooses the metric's labels on collisions with target labels (default true)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"honor_timestamps": schema.BoolAttribute{
												Description:         "HonorTimestamps controls whether the scrape endpoints honor timestamps (default false)",
												MarkdownDescription: "HonorTimestamps controls whether the scrape endpoints honor timestamps (default false)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"interval": schema.StringAttribute{
												Description:         "How frequently to scrape metrics. Accepts values with time unit suffix: '30s', '1m', '2h', '500ms'",
												MarkdownDescription: "How frequently to scrape metrics. Accepts values with time unit suffix: '30s', '1m', '2h', '500ms'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+)(ms|s|m|h)$`), ""),
												},
											},

											"labels": schema.MapAttribute{
												Description:         "Additional labels to add to the ServiceMonitor (default release: prometheus)",
												MarkdownDescription: "Additional labels to add to the ServiceMonitor (default release: prometheus)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metric_relabelings": schema.ListNestedAttribute{
												Description:         "Relabeling rules applied to individual scraped metrics",
												MarkdownDescription: "Relabeling rules applied to individual scraped metrics",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.StringAttribute{
															Description:         "Action to perform based on the regex matching. 'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0. Default: 'Replace'",
															MarkdownDescription: "Action to perform based on the regex matching. 'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0. Default: 'Replace'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
															},
														},

														"modulus": schema.Int64Attribute{
															Description:         "Modulus to take of the hash of the source label values. Only applicable when the action is 'HashMod'.",
															MarkdownDescription: "Modulus to take of the hash of the source label values. Only applicable when the action is 'HashMod'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"regex": schema.StringAttribute{
															Description:         "Regular expression against which the extracted value is matched.",
															MarkdownDescription: "Regular expression against which the extracted value is matched.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replacement": schema.StringAttribute{
															Description:         "Replacement value against which a Replace action is performed if the regular expression matches. Regex capture groups are available.",
															MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches. Regex capture groups are available.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"separator": schema.StringAttribute{
															Description:         "Separator is the string between concatenated SourceLabels.",
															MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_labels": schema.ListAttribute{
															Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
															MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_label": schema.StringAttribute{
															Description:         "Label to which the resulting string is written in a replacement. It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions. Regex capture groups are available.",
															MarkdownDescription: "Label to which the resulting string is written in a replacement. It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions. Regex capture groups are available.",
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

											"relabelings": schema.ListNestedAttribute{
												Description:         "RelabelConfigs to apply to samples before ingestion",
												MarkdownDescription: "RelabelConfigs to apply to samples before ingestion",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.StringAttribute{
															Description:         "Action to perform based on the regex matching. 'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0. Default: 'Replace'",
															MarkdownDescription: "Action to perform based on the regex matching. 'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0. Default: 'Replace'",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
															},
														},

														"modulus": schema.Int64Attribute{
															Description:         "Modulus to take of the hash of the source label values. Only applicable when the action is 'HashMod'.",
															MarkdownDescription: "Modulus to take of the hash of the source label values. Only applicable when the action is 'HashMod'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"regex": schema.StringAttribute{
															Description:         "Regular expression against which the extracted value is matched.",
															MarkdownDescription: "Regular expression against which the extracted value is matched.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replacement": schema.StringAttribute{
															Description:         "Replacement value against which a Replace action is performed if the regular expression matches. Regex capture groups are available.",
															MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches. Regex capture groups are available.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"separator": schema.StringAttribute{
															Description:         "Separator is the string between concatenated SourceLabels.",
															MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_labels": schema.ListAttribute{
															Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
															MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_label": schema.StringAttribute{
															Description:         "Label to which the resulting string is written in a replacement. It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions. Regex capture groups are available.",
															MarkdownDescription: "Label to which the resulting string is written in a replacement. It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions. Regex capture groups are available.",
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

											"tls_config": schema.SingleNestedAttribute{
												Description:         "TLS settings used by Prometheus to connect to the metrics endpoint",
												MarkdownDescription: "TLS settings used by Prometheus to connect to the metrics endpoint",
												Attributes: map[string]schema.Attribute{
													"ca": schema.SingleNestedAttribute{
														Description:         "Certificate authority used when verifying server certificates.",
														MarkdownDescription: "Certificate authority used when verifying server certificates.",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																		MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

															"secret": schema.SingleNestedAttribute{
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from. Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																		MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

													"ca_file": schema.StringAttribute{
														Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
														MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert": schema.SingleNestedAttribute{
														Description:         "Client certificate to present when doing client-authentication.",
														MarkdownDescription: "Client certificate to present when doing client-authentication.",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																		MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

															"secret": schema.SingleNestedAttribute{
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from. Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																		MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

													"cert_file": schema.StringAttribute{
														Description:         "Path to the client cert file in the Prometheus container for the targets.",
														MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_file": schema.StringAttribute{
														Description:         "Path to the client key file in the Prometheus container for the targets.",
														MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_secret": schema.SingleNestedAttribute{
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from. Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

													"max_version": schema.StringAttribute{
														Description:         "Maximum acceptable TLS version. It requires Prometheus >= v2.41.0.",
														MarkdownDescription: "Maximum acceptable TLS version. It requires Prometheus >= v2.41.0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
														},
													},

													"min_version": schema.StringAttribute{
														Description:         "Minimum acceptable TLS version. It requires Prometheus >= v2.35.0.",
														MarkdownDescription: "Minimum acceptable TLS version. It requires Prometheus >= v2.35.0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
														},
													},

													"server_name": schema.StringAttribute{
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",
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

							"rbac_config": schema.SingleNestedAttribute{
								Description:         "optional kube-rbac-proxy config to provide rbac services",
								MarkdownDescription: "optional kube-rbac-proxy config to provide rbac services",
								Attributes: map[string]schema.Attribute{
									"client_ca_config_map": schema.SingleNestedAttribute{
										Description:         "Reference to a configmap containing the client CA (key: ca.crt) for mTLS client validation",
										MarkdownDescription: "Reference to a configmap containing the client CA (key: ca.crt) for mTLS client validation",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_https": schema.BoolAttribute{
										Description:         "disable https protecting the proxy endpoint",
										MarkdownDescription: "disable https protecting the proxy endpoint",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable": schema.BoolAttribute{
										Description:         "enable kube-rbac-proxy, disabled by default",
										MarkdownDescription: "enable kube-rbac-proxy, disabled by default",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "kube-rbac-proxy image",
										MarkdownDescription: "kube-rbac-proxy image",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
										},
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "certificate secret to mount in kube-rbac container for TLS, self signed certificates will be generated by default",
										MarkdownDescription: "certificate secret to mount in kube-rbac container for TLS, self signed certificates will be generated by default",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"static_authorization": schema.SingleNestedAttribute{
										Description:         "Optional static RBAC rules based on client certificate Common Name (CN)",
										MarkdownDescription: "Optional static RBAC rules based on client certificate Common Name (CN)",
										Attributes: map[string]schema.Attribute{
											"client_name": schema.StringAttribute{
												Description:         "Expected CN (Common Name) from client cert (e.g., Prometheus SA identity)",
												MarkdownDescription: "Expected CN (Common Name) from client cert (e.g., Prometheus SA identity)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enables static authorization using client certificate CN",
												MarkdownDescription: "Enables static authorization using client certificate CN",
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

							"selector": schema.MapAttribute{
								Description:         "Selector describes on which nodes to enable metrics exporter",
								MarkdownDescription: "Selector describes on which nodes to enable metrics exporter",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_type": schema.StringAttribute{
								Description:         "ServiceType service type for metrics, clusterIP/NodePort, clusterIP by default",
								MarkdownDescription: "ServiceType service type for metrics, clusterIP/NodePort, clusterIP by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ClusterIP", "NodePort"),
								},
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "tolerations for metrics exporter",
								MarkdownDescription: "tolerations for metrics exporter",
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

							"upgrade_policy": schema.SingleNestedAttribute{
								Description:         "upgrade policy for metrics exporter daemons",
								MarkdownDescription: "upgrade policy for metrics exporter daemons",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.Int64Attribute{
										Description:         "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										MarkdownDescription: "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"upgrade_strategy": schema.StringAttribute{
										Description:         "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										MarkdownDescription: "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("RollingUpdate", "OnDelete"),
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

					"remediation_workflow": schema.SingleNestedAttribute{
						Description:         "remediation workflow",
						MarkdownDescription: "remediation workflow",
						Attributes: map[string]schema.Attribute{
							"conditional_workflows": schema.SingleNestedAttribute{
								Description:         "Name of the ConfigMap that holds condition-to-workflow mappings.",
								MarkdownDescription: "Name of the ConfigMap that holds condition-to-workflow mappings.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "enable remediation workflows. disabled by default enable if operator should automatically handle remediation of node incase of gpu issues",
								MarkdownDescription: "enable remediation workflows. disabled by default enable if operator should automatically handle remediation of node incase of gpu issues",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ttl_for_failed_workflows": schema.Int64Attribute{
								Description:         "Time to live for argo workflow object and its pods for a failed workflow in hours. By default, it is set to 24 hours",
								MarkdownDescription: "Time to live for argo workflow object and its pods for a failed workflow in hours. By default, it is set to 24 hours",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.MapAttribute{
						Description:         "Selector describes on which nodes the GPU Operator should enable the GPU device.",
						MarkdownDescription: "Selector describes on which nodes the GPU Operator should enable the GPU device.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"test_runner": schema.SingleNestedAttribute{
						Description:         "test runner",
						MarkdownDescription: "test runner",
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Description:         "config map to customize the config for test runner, if not specified default test config will be aplied",
								MarkdownDescription: "config map to customize the config for test runner, if not specified default test config will be aplied",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "enable test runner, disabled by default",
								MarkdownDescription: "enable test runner, disabled by default",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "test runner image",
								MarkdownDescription: "test runner image",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]+(?:[._-][a-z0-9]+)*(:[0-9]+)?)(/[a-z0-9]+(?:[._-][a-z0-9]+)*)*(?::[a-z0-9._-]+)?(?:@[a-zA-Z0-9]+:[a-f0-9]+)?$`), ""),
								},
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "image pull policy for test runner",
								MarkdownDescription: "image pull policy for test runner",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"image_registry_secret": schema.SingleNestedAttribute{
								Description:         "test runner image registry secret used to pull/push images",
								MarkdownDescription: "test runner image registry secret used to pull/push images",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logs_location": schema.SingleNestedAttribute{
								Description:         "captures logs location and export config for test runner logs",
								MarkdownDescription: "captures logs location and export config for test runner logs",
								Attributes: map[string]schema.Attribute{
									"host_path": schema.StringAttribute{
										Description:         "host path to store test runner internal status db in order to persist test running status",
										MarkdownDescription: "host path to store test runner internal status db in order to persist test running status",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"logs_export_secrets": schema.ListNestedAttribute{
										Description:         "LogsExportSecrets is a list of secrets that contain connectivity info to multiple cloud providers",
										MarkdownDescription: "LogsExportSecrets is a list of secrets that contain connectivity info to multiple cloud providers",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"mount_path": schema.StringAttribute{
										Description:         "volume mount destination within test runner container",
										MarkdownDescription: "volume mount destination within test runner container",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": schema.MapAttribute{
								Description:         "Selector describes on which nodes to enable test runner",
								MarkdownDescription: "Selector describes on which nodes to enable test runner",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "tolerations for test runner",
								MarkdownDescription: "tolerations for test runner",
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

							"upgrade_policy": schema.SingleNestedAttribute{
								Description:         "upgrade policy for test runner daemonset",
								MarkdownDescription: "upgrade policy for test runner daemonset",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.Int64Attribute{
										Description:         "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										MarkdownDescription: "MaxUnavailable specifies the maximum number of Pods that can be unavailable during the update process. Applicable for RollingUpdate only. Default value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"upgrade_strategy": schema.StringAttribute{
										Description:         "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										MarkdownDescription: "UpgradeStrategy specifies the type of the DaemonSet update. Valid values are 'RollingUpdate' (default) or 'OnDelete'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("RollingUpdate", "OnDelete"),
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
		},
	}
}

func (r *AmdComDeviceConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_amd_com_device_config_v1alpha1_manifest")

	var model AmdComDeviceConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("amd.com/v1alpha1")
	model.Kind = pointer.String("DeviceConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
