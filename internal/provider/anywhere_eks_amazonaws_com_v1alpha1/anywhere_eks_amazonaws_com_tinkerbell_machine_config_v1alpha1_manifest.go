/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1ManifestData struct {
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
		HardwareSelector    *map[string]string `tfsdk:"hardware_selector" json:"hardwareSelector,omitempty"`
		HostOSConfiguration *struct {
			BottlerocketConfiguration *struct {
				Boot *struct {
					BootKernelParameters *map[string][]string `tfsdk:"boot_kernel_parameters" json:"bootKernelParameters,omitempty"`
				} `tfsdk:"boot" json:"boot,omitempty"`
				Kernel *struct {
					SysctlSettings *map[string]string `tfsdk:"sysctl_settings" json:"sysctlSettings,omitempty"`
				} `tfsdk:"kernel" json:"kernel,omitempty"`
				Kubernetes *struct {
					AllowedUnsafeSysctls            *[]string          `tfsdk:"allowed_unsafe_sysctls" json:"allowedUnsafeSysctls,omitempty"`
					ClusterDNSIPs                   *[]string          `tfsdk:"cluster_dnsi_ps" json:"clusterDNSIPs,omitempty"`
					ClusterDomain                   *string            `tfsdk:"cluster_domain" json:"clusterDomain,omitempty"`
					ContainerLogMaxFiles            *int64             `tfsdk:"container_log_max_files" json:"containerLogMaxFiles,omitempty"`
					ContainerLogMaxSize             *string            `tfsdk:"container_log_max_size" json:"containerLogMaxSize,omitempty"`
					CpuCFSQuota                     *bool              `tfsdk:"cpu_cfs_quota" json:"cpuCFSQuota,omitempty"`
					CpuManagerPolicy                *string            `tfsdk:"cpu_manager_policy" json:"cpuManagerPolicy,omitempty"`
					CpuManagerPolicyOptions         *map[string]string `tfsdk:"cpu_manager_policy_options" json:"cpuManagerPolicyOptions,omitempty"`
					CpuManagerReconcilePeriod       *string            `tfsdk:"cpu_manager_reconcile_period" json:"cpuManagerReconcilePeriod,omitempty"`
					EventBurst                      *int64             `tfsdk:"event_burst" json:"eventBurst,omitempty"`
					EventRecordQPS                  *int64             `tfsdk:"event_record_qps" json:"eventRecordQPS,omitempty"`
					EvictionHard                    *map[string]string `tfsdk:"eviction_hard" json:"evictionHard,omitempty"`
					EvictionMaxPodGracePeriod       *int64             `tfsdk:"eviction_max_pod_grace_period" json:"evictionMaxPodGracePeriod,omitempty"`
					EvictionSoft                    *map[string]string `tfsdk:"eviction_soft" json:"evictionSoft,omitempty"`
					EvictionSoftGracePeriod         *map[string]string `tfsdk:"eviction_soft_grace_period" json:"evictionSoftGracePeriod,omitempty"`
					ImageGCHighThresholdPercent     *int64             `tfsdk:"image_gc_high_threshold_percent" json:"imageGCHighThresholdPercent,omitempty"`
					ImageGCLowThresholdPercent      *int64             `tfsdk:"image_gc_low_threshold_percent" json:"imageGCLowThresholdPercent,omitempty"`
					KubeAPIBurst                    *int64             `tfsdk:"kube_api_burst" json:"kubeAPIBurst,omitempty"`
					KubeAPIQPS                      *int64             `tfsdk:"kube_apiqps" json:"kubeAPIQPS,omitempty"`
					KubeReserved                    *map[string]string `tfsdk:"kube_reserved" json:"kubeReserved,omitempty"`
					MaxPods                         *int64             `tfsdk:"max_pods" json:"maxPods,omitempty"`
					MemoryManagerPolicy             *string            `tfsdk:"memory_manager_policy" json:"memoryManagerPolicy,omitempty"`
					PodPidsLimit                    *int64             `tfsdk:"pod_pids_limit" json:"podPidsLimit,omitempty"`
					ProviderID                      *string            `tfsdk:"provider_id" json:"providerID,omitempty"`
					RegistryBurst                   *int64             `tfsdk:"registry_burst" json:"registryBurst,omitempty"`
					RegistryPullQPS                 *int64             `tfsdk:"registry_pull_qps" json:"registryPullQPS,omitempty"`
					ShutdownGracePeriod             *string            `tfsdk:"shutdown_grace_period" json:"shutdownGracePeriod,omitempty"`
					ShutdownGracePeriodCriticalPods *string            `tfsdk:"shutdown_grace_period_critical_pods" json:"shutdownGracePeriodCriticalPods,omitempty"`
					SystemReserved                  *map[string]string `tfsdk:"system_reserved" json:"systemReserved,omitempty"`
					TopologyManagerPolicy           *string            `tfsdk:"topology_manager_policy" json:"topologyManagerPolicy,omitempty"`
					TopologyManagerScope            *string            `tfsdk:"topology_manager_scope" json:"topologyManagerScope,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			} `tfsdk:"bottlerocket_configuration" json:"bottlerocketConfiguration,omitempty"`
			CertBundles *[]struct {
				Data *string `tfsdk:"data" json:"data,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"cert_bundles" json:"certBundles,omitempty"`
			NtpConfiguration *struct {
				Servers *[]string `tfsdk:"servers" json:"servers,omitempty"`
			} `tfsdk:"ntp_configuration" json:"ntpConfiguration,omitempty"`
		} `tfsdk:"host_os_configuration" json:"hostOSConfiguration,omitempty"`
		OsFamily    *string `tfsdk:"os_family" json:"osFamily,omitempty"`
		OsImageURL  *string `tfsdk:"os_image_url" json:"osImageURL,omitempty"`
		TemplateRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"template_ref" json:"templateRef,omitempty"`
		Users *[]struct {
			Name              *string   `tfsdk:"name" json:"name,omitempty"`
			SshAuthorizedKeys *[]string `tfsdk:"ssh_authorized_keys" json:"sshAuthorizedKeys,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TinkerbellMachineConfig is the Schema for the tinkerbellmachineconfigs API.",
		MarkdownDescription: "TinkerbellMachineConfig is the Schema for the tinkerbellmachineconfigs API.",
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
				Description:         "TinkerbellMachineConfigSpec defines the desired state of TinkerbellMachineConfig.",
				MarkdownDescription: "TinkerbellMachineConfigSpec defines the desired state of TinkerbellMachineConfig.",
				Attributes: map[string]schema.Attribute{
					"hardware_selector": schema.MapAttribute{
						Description:         "HardwareSelector models a simple key-value selector used in Tinkerbell provisioning.",
						MarkdownDescription: "HardwareSelector models a simple key-value selector used in Tinkerbell provisioning.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"host_os_configuration": schema.SingleNestedAttribute{
						Description:         "HostOSConfiguration defines the configuration settings on the host OS.",
						MarkdownDescription: "HostOSConfiguration defines the configuration settings on the host OS.",
						Attributes: map[string]schema.Attribute{
							"bottlerocket_configuration": schema.SingleNestedAttribute{
								Description:         "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								MarkdownDescription: "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								Attributes: map[string]schema.Attribute{
									"boot": schema.SingleNestedAttribute{
										Description:         "Boot defines the boot settings for bottlerocket.",
										MarkdownDescription: "Boot defines the boot settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"boot_kernel_parameters": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kernel": schema.SingleNestedAttribute{
										Description:         "Kernel defines the kernel settings for bottlerocket.",
										MarkdownDescription: "Kernel defines the kernel settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"sysctl_settings": schema.MapAttribute{
												Description:         "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
												MarkdownDescription: "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
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

									"kubernetes": schema.SingleNestedAttribute{
										Description:         "Kubernetes defines the Kubernetes settings on the host OS.",
										MarkdownDescription: "Kubernetes defines the Kubernetes settings on the host OS.",
										Attributes: map[string]schema.Attribute{
											"allowed_unsafe_sysctls": schema.ListAttribute{
												Description:         "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												MarkdownDescription: "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_dnsi_ps": schema.ListAttribute{
												Description:         "ClusterDNSIPs defines IP addresses of the DNS servers.",
												MarkdownDescription: "ClusterDNSIPs defines IP addresses of the DNS servers.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_domain": schema.StringAttribute{
												Description:         "ClusterDomain defines the DNS domain for the cluster, allowing all Kubernetes-run containers to search this domain before the host’s search domains",
												MarkdownDescription: "ClusterDomain defines the DNS domain for the cluster, allowing all Kubernetes-run containers to search this domain before the host’s search domains",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container_log_max_files": schema.Int64Attribute{
												Description:         "ContainerLogMaxFiles specifies the maximum number of container log files that can be present for a container",
												MarkdownDescription: "ContainerLogMaxFiles specifies the maximum number of container log files that can be present for a container",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container_log_max_size": schema.StringAttribute{
												Description:         "ContainerLogMaxSize is a quantity defining the maximum size of the container log file before it is rotated",
												MarkdownDescription: "ContainerLogMaxSize is a quantity defining the maximum size of the container log file before it is rotated",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cpu_cfs_quota": schema.BoolAttribute{
												Description:         "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits",
												MarkdownDescription: "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cpu_manager_policy": schema.StringAttribute{
												Description:         "CPUManagerPolicy is the name of the policy to use.",
												MarkdownDescription: "CPUManagerPolicy is the name of the policy to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cpu_manager_policy_options": schema.MapAttribute{
												Description:         "CPUManagerPolicyOptions is a set of key=value which allows to set extra options to fine tune the behaviour of the cpu manager policies",
												MarkdownDescription: "CPUManagerPolicyOptions is a set of key=value which allows to set extra options to fine tune the behaviour of the cpu manager policies",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cpu_manager_reconcile_period": schema.StringAttribute{
												Description:         "CPUManagerReconcilePeriod is the reconciliation period for the CPU Manager.",
												MarkdownDescription: "CPUManagerReconcilePeriod is the reconciliation period for the CPU Manager.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"event_burst": schema.Int64Attribute{
												Description:         "EventBurst is the maximum size of a burst of event creations.",
												MarkdownDescription: "EventBurst is the maximum size of a burst of event creations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"event_record_qps": schema.Int64Attribute{
												Description:         "EventRecordQPS is the maximum event creations per second.",
												MarkdownDescription: "EventRecordQPS is the maximum event creations per second.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_hard": schema.MapAttribute{
												Description:         "EvictionHard is a map of signal names to quantities that defines hard eviction thresholds.",
												MarkdownDescription: "EvictionHard is a map of signal names to quantities that defines hard eviction thresholds.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_max_pod_grace_period": schema.Int64Attribute{
												Description:         "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.",
												MarkdownDescription: "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_soft": schema.MapAttribute{
												Description:         "EvictionSoft is a map of signal names to quantities that defines soft eviction thresholds.",
												MarkdownDescription: "EvictionSoft is a map of signal names to quantities that defines soft eviction thresholds.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_soft_grace_period": schema.MapAttribute{
												Description:         "EvictionSoftGracePeriod is a map of signal names to quantities that defines grace periods for each soft eviction signal.",
												MarkdownDescription: "EvictionSoftGracePeriod is a map of signal names to quantities that defines grace periods for each soft eviction signal.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_gc_high_threshold_percent": schema.Int64Attribute{
												Description:         "ImageGCHighThresholdPercent is the percent of disk usage after which image garbage collection is always run.",
												MarkdownDescription: "ImageGCHighThresholdPercent is the percent of disk usage after which image garbage collection is always run.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_gc_low_threshold_percent": schema.Int64Attribute{
												Description:         "ImageGCLowThresholdPercent is the percent of disk usage before which image garbage collection is never run.",
												MarkdownDescription: "ImageGCLowThresholdPercent is the percent of disk usage before which image garbage collection is never run.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kube_api_burst": schema.Int64Attribute{
												Description:         "KubeAPIBurst is the burst to allow while talking with kubernetes API server.",
												MarkdownDescription: "KubeAPIBurst is the burst to allow while talking with kubernetes API server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kube_apiqps": schema.Int64Attribute{
												Description:         "KubeAPIQPS is the QPS to use while talking with kubernetes apiserver.",
												MarkdownDescription: "KubeAPIQPS is the QPS to use while talking with kubernetes apiserver.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kube_reserved": schema.MapAttribute{
												Description:         "KubeReserved is a set of ResourceName=ResourceQuantity pairs that describe resources reserved for kubernetes system components",
												MarkdownDescription: "KubeReserved is a set of ResourceName=ResourceQuantity pairs that describe resources reserved for kubernetes system components",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_pods": schema.Int64Attribute{
												Description:         "MaxPods defines the maximum number of pods that can run on a node.",
												MarkdownDescription: "MaxPods defines the maximum number of pods that can run on a node.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memory_manager_policy": schema.StringAttribute{
												Description:         "MemoryManagerPolicy is the name of the policy to use by memory manager.",
												MarkdownDescription: "MemoryManagerPolicy is the name of the policy to use by memory manager.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_pids_limit": schema.Int64Attribute{
												Description:         "PodPidsLimit is the maximum number of PIDs in any pod.",
												MarkdownDescription: "PodPidsLimit is the maximum number of PIDs in any pod.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"provider_id": schema.StringAttribute{
												Description:         "ProviderID sets the unique ID of the instance that an external provider.",
												MarkdownDescription: "ProviderID sets the unique ID of the instance that an external provider.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"registry_burst": schema.Int64Attribute{
												Description:         "RegistryBurst is the maximum size of bursty pulls.",
												MarkdownDescription: "RegistryBurst is the maximum size of bursty pulls.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"registry_pull_qps": schema.Int64Attribute{
												Description:         "RegistryPullQPS is the limit of registry pulls per second.",
												MarkdownDescription: "RegistryPullQPS is the limit of registry pulls per second.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"shutdown_grace_period": schema.StringAttribute{
												Description:         "ShutdownGracePeriod specifies the total duration that the node should delay the shutdown and total grace period for pod termination during a node shutdown.",
												MarkdownDescription: "ShutdownGracePeriod specifies the total duration that the node should delay the shutdown and total grace period for pod termination during a node shutdown.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"shutdown_grace_period_critical_pods": schema.StringAttribute{
												Description:         "ShutdownGracePeriodCriticalPods specifies the duration used to terminate critical pods during a node shutdown.",
												MarkdownDescription: "ShutdownGracePeriodCriticalPods specifies the duration used to terminate critical pods during a node shutdown.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"system_reserved": schema.MapAttribute{
												Description:         "SystemReserved is a set of ResourceName=ResourceQuantity pairs that describe resources reserved for non-kubernetes components.",
												MarkdownDescription: "SystemReserved is a set of ResourceName=ResourceQuantity pairs that describe resources reserved for non-kubernetes components.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topology_manager_policy": schema.StringAttribute{
												Description:         "TopologyManagerPolicy is the name of the topology manager policy to use.",
												MarkdownDescription: "TopologyManagerPolicy is the name of the topology manager policy to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topology_manager_scope": schema.StringAttribute{
												Description:         "TopologyManagerScope represents the scope of topology hint generation that topology manager requests and hint providers generate.",
												MarkdownDescription: "TopologyManagerScope represents the scope of topology hint generation that topology manager requests and hint providers generate.",
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

							"cert_bundles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"data": schema.StringAttribute{
											Description:         "Data defines the cert bundle data.",
											MarkdownDescription: "Data defines the cert bundle data.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name defines the cert bundle name.",
											MarkdownDescription: "Name defines the cert bundle name.",
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

							"ntp_configuration": schema.SingleNestedAttribute{
								Description:         "NTPConfiguration defines the NTP configuration on the host OS.",
								MarkdownDescription: "NTPConfiguration defines the NTP configuration on the host OS.",
								Attributes: map[string]schema.Attribute{
									"servers": schema.ListAttribute{
										Description:         "Servers defines a list of NTP servers to be configured on the host OS.",
										MarkdownDescription: "Servers defines a list of NTP servers to be configured on the host OS.",
										ElementType:         types.StringType,
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

					"os_family": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"os_image_url": schema.StringAttribute{
						Description:         "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						MarkdownDescription: "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ssh_authorized_keys": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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
		},
	}
}

func (r *AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_tinkerbell_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComTinkerbellMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("TinkerbellMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
