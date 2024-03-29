/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComClusterV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComClusterV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComClusterV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComClusterV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComClusterV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		BundlesRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"bundles_ref" json:"bundlesRef,omitempty"`
		ClusterNetwork *struct {
			Cni       *string `tfsdk:"cni" json:"cni,omitempty"`
			CniConfig *struct {
				Cilium *struct {
					EgressMasqueradeInterfaces *string `tfsdk:"egress_masquerade_interfaces" json:"egressMasqueradeInterfaces,omitempty"`
					Ipv4NativeRoutingCIDR      *string `tfsdk:"ipv4_native_routing_cidr" json:"ipv4NativeRoutingCIDR,omitempty"`
					Ipv6NativeRoutingCIDR      *string `tfsdk:"ipv6_native_routing_cidr" json:"ipv6NativeRoutingCIDR,omitempty"`
					PolicyEnforcementMode      *string `tfsdk:"policy_enforcement_mode" json:"policyEnforcementMode,omitempty"`
					RoutingMode                *string `tfsdk:"routing_mode" json:"routingMode,omitempty"`
					SkipUpgrade                *bool   `tfsdk:"skip_upgrade" json:"skipUpgrade,omitempty"`
				} `tfsdk:"cilium" json:"cilium,omitempty"`
				Kindnetd *map[string]string `tfsdk:"kindnetd" json:"kindnetd,omitempty"`
			} `tfsdk:"cni_config" json:"cniConfig,omitempty"`
			Dns *struct {
				ResolvConf *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"resolv_conf" json:"resolvConf,omitempty"`
			} `tfsdk:"dns" json:"dns,omitempty"`
			Nodes *struct {
				CidrMaskSize *int64 `tfsdk:"cidr_mask_size" json:"cidrMaskSize,omitempty"`
			} `tfsdk:"nodes" json:"nodes,omitempty"`
			Pods *struct {
				CidrBlocks *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
			Services *struct {
				CidrBlocks *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"cluster_network" json:"clusterNetwork,omitempty"`
		ControlPlaneConfiguration *struct {
			CertSans *[]string `tfsdk:"cert_sans" json:"certSans,omitempty"`
			Count    *int64    `tfsdk:"count" json:"count,omitempty"`
			Endpoint *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
			} `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Labels          *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MachineGroupRef *struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"machine_group_ref" json:"machineGroupRef,omitempty"`
			MachineHealthCheck *struct {
				MaxUnhealthy            *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
				NodeStartupTimeout      *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
				UnhealthyMachineTimeout *string `tfsdk:"unhealthy_machine_timeout" json:"unhealthyMachineTimeout,omitempty"`
			} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
			SkipLoadBalancerDeployment *bool `tfsdk:"skip_load_balancer_deployment" json:"skipLoadBalancerDeployment,omitempty"`
			Taints                     *[]struct {
				Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"taints" json:"taints,omitempty"`
			UpgradeRolloutStrategy *struct {
				RollingUpdate *struct {
					MaxSurge *int64 `tfsdk:"max_surge" json:"maxSurge,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"upgrade_rollout_strategy" json:"upgradeRolloutStrategy,omitempty"`
		} `tfsdk:"control_plane_configuration" json:"controlPlaneConfiguration,omitempty"`
		DatacenterRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"datacenter_ref" json:"datacenterRef,omitempty"`
		EksaVersion    *string `tfsdk:"eksa_version" json:"eksaVersion,omitempty"`
		EtcdEncryption *[]struct {
			Providers *[]struct {
				Kms *struct {
					Cachesize           *int64  `tfsdk:"cachesize" json:"cachesize,omitempty"`
					Name                *string `tfsdk:"name" json:"name,omitempty"`
					SocketListenAddress *string `tfsdk:"socket_listen_address" json:"socketListenAddress,omitempty"`
					Timeout             *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"kms" json:"kms,omitempty"`
			} `tfsdk:"providers" json:"providers,omitempty"`
			Resources *[]string `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"etcd_encryption" json:"etcdEncryption,omitempty"`
		ExternalEtcdConfiguration *struct {
			Count           *int64 `tfsdk:"count" json:"count,omitempty"`
			MachineGroupRef *struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"machine_group_ref" json:"machineGroupRef,omitempty"`
		} `tfsdk:"external_etcd_configuration" json:"externalEtcdConfiguration,omitempty"`
		GitOpsRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"git_ops_ref" json:"gitOpsRef,omitempty"`
		IdentityProviderRefs *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"identity_provider_refs" json:"identityProviderRefs,omitempty"`
		KubernetesVersion  *string `tfsdk:"kubernetes_version" json:"kubernetesVersion,omitempty"`
		MachineHealthCheck *struct {
			MaxUnhealthy            *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
			NodeStartupTimeout      *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
			UnhealthyMachineTimeout *string `tfsdk:"unhealthy_machine_timeout" json:"unhealthyMachineTimeout,omitempty"`
		} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
		ManagementCluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"management_cluster" json:"managementCluster,omitempty"`
		Packages *struct {
			Controller *struct {
				Digest          *string   `tfsdk:"digest" json:"digest,omitempty"`
				DisableWebhooks *bool     `tfsdk:"disable_webhooks" json:"disableWebhooks,omitempty"`
				Env             *[]string `tfsdk:"env" json:"env,omitempty"`
				Repository      *string   `tfsdk:"repository" json:"repository,omitempty"`
				Resources       *struct {
					Limits *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"limits" json:"limits,omitempty"`
					Requests *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tag *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"controller" json:"controller,omitempty"`
			Cronjob *struct {
				Digest     *string `tfsdk:"digest" json:"digest,omitempty"`
				Disable    *bool   `tfsdk:"disable" json:"disable,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"cronjob" json:"cronjob,omitempty"`
			Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
		} `tfsdk:"packages" json:"packages,omitempty"`
		PodIamConfig *struct {
			ServiceAccountIssuer *string `tfsdk:"service_account_issuer" json:"serviceAccountIssuer,omitempty"`
		} `tfsdk:"pod_iam_config" json:"podIamConfig,omitempty"`
		ProxyConfiguration *struct {
			HttpProxy  *string   `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
			HttpsProxy *string   `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
			NoProxy    *[]string `tfsdk:"no_proxy" json:"noProxy,omitempty"`
		} `tfsdk:"proxy_configuration" json:"proxyConfiguration,omitempty"`
		RegistryMirrorConfiguration *struct {
			Authenticate       *bool   `tfsdk:"authenticate" json:"authenticate,omitempty"`
			CaCertContent      *string `tfsdk:"ca_cert_content" json:"caCertContent,omitempty"`
			Endpoint           *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			OciNamespaces      *[]struct {
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Registry  *string `tfsdk:"registry" json:"registry,omitempty"`
			} `tfsdk:"oci_namespaces" json:"ociNamespaces,omitempty"`
			Port *string `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"registry_mirror_configuration" json:"registryMirrorConfiguration,omitempty"`
		WorkerNodeGroupConfigurations *[]struct {
			AutoscalingConfiguration *struct {
				MaxCount *int64 `tfsdk:"max_count" json:"maxCount,omitempty"`
				MinCount *int64 `tfsdk:"min_count" json:"minCount,omitempty"`
			} `tfsdk:"autoscaling_configuration" json:"autoscalingConfiguration,omitempty"`
			Count             *int64             `tfsdk:"count" json:"count,omitempty"`
			KubernetesVersion *string            `tfsdk:"kubernetes_version" json:"kubernetesVersion,omitempty"`
			Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MachineGroupRef   *struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"machine_group_ref" json:"machineGroupRef,omitempty"`
			MachineHealthCheck *struct {
				MaxUnhealthy            *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
				NodeStartupTimeout      *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
				UnhealthyMachineTimeout *string `tfsdk:"unhealthy_machine_timeout" json:"unhealthyMachineTimeout,omitempty"`
			} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Taints *[]struct {
				Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"taints" json:"taints,omitempty"`
			UpgradeRolloutStrategy *struct {
				RollingUpdate *struct {
					MaxSurge       *int64 `tfsdk:"max_surge" json:"maxSurge,omitempty"`
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"upgrade_rollout_strategy" json:"upgradeRolloutStrategy,omitempty"`
		} `tfsdk:"worker_node_group_configurations" json:"workerNodeGroupConfigurations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the clusters API.",
		MarkdownDescription: "Cluster is the Schema for the clusters API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "ClusterSpec defines the desired state of Cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.",
				Attributes: map[string]schema.Attribute{
					"bundles_ref": schema.SingleNestedAttribute{
						Description:         "BundlesRef contains a reference to the Bundles containing the desired dependencies for the cluster. DEPRECATED: Use EksaVersion instead.",
						MarkdownDescription: "BundlesRef contains a reference to the Bundles containing the desired dependencies for the cluster. DEPRECATED: Use EksaVersion instead.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion refers to the Bundles APIVersion",
								MarkdownDescription: "APIVersion refers to the Bundles APIVersion",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name refers to the name of the Bundles object in the cluster",
								MarkdownDescription: "Name refers to the name of the Bundles object in the cluster",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace refers to the Bundles's namespace",
								MarkdownDescription: "Namespace refers to the Bundles's namespace",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_network": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cni": schema.StringAttribute{
								Description:         "Deprecated. Use CNIConfig",
								MarkdownDescription: "Deprecated. Use CNIConfig",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cni_config": schema.SingleNestedAttribute{
								Description:         "CNIConfig specifies the CNI plugin to be installed in the cluster",
								MarkdownDescription: "CNIConfig specifies the CNI plugin to be installed in the cluster",
								Attributes: map[string]schema.Attribute{
									"cilium": schema.SingleNestedAttribute{
										Description:         "CiliumConfig contains configuration specific to the Cilium CNI.",
										MarkdownDescription: "CiliumConfig contains configuration specific to the Cilium CNI.",
										Attributes: map[string]schema.Attribute{
											"egress_masquerade_interfaces": schema.StringAttribute{
												Description:         "EgressMasquaradeInterfaces determines which network interfaces are used for masquerading. Accepted values are a valid interface name or interface prefix.",
												MarkdownDescription: "EgressMasquaradeInterfaces determines which network interfaces are used for masquerading. Accepted values are a valid interface name or interface prefix.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ipv4_native_routing_cidr": schema.StringAttribute{
												Description:         "IPv4NativeRoutingCIDR specifies the CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true",
												MarkdownDescription: "IPv4NativeRoutingCIDR specifies the CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ipv6_native_routing_cidr": schema.StringAttribute{
												Description:         "IPv6NativeRoutingCIDR specifies the IPv6 CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true",
												MarkdownDescription: "IPv6NativeRoutingCIDR specifies the IPv6 CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"policy_enforcement_mode": schema.StringAttribute{
												Description:         "PolicyEnforcementMode determines communication allowed between pods. Accepted values are default, always, never.",
												MarkdownDescription: "PolicyEnforcementMode determines communication allowed between pods. Accepted values are default, always, never.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"routing_mode": schema.StringAttribute{
												Description:         "RoutingMode indicates the routing tunnel mode to use for Cilium. Accepted values are overlay (geneve tunnel with overlay) or direct (tunneling disabled with direct routing) Defaults to overlay.",
												MarkdownDescription: "RoutingMode indicates the routing tunnel mode to use for Cilium. Accepted values are overlay (geneve tunnel with overlay) or direct (tunneling disabled with direct routing) Defaults to overlay.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"skip_upgrade": schema.BoolAttribute{
												Description:         "SkipUpgrade indicicates that Cilium maintenance should be skipped during upgrades. This can be used when operators wish to self manage the Cilium installation.",
												MarkdownDescription: "SkipUpgrade indicicates that Cilium maintenance should be skipped during upgrades. This can be used when operators wish to self manage the Cilium installation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kindnetd": schema.MapAttribute{
										Description:         "KindnetdConfig contains configuration specific to the Kindnetd CNI.",
										MarkdownDescription: "KindnetdConfig contains configuration specific to the Kindnetd CNI.",
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

							"dns": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"resolv_conf": schema.SingleNestedAttribute{
										Description:         "ResolvConf refers to the DNS resolver configuration",
										MarkdownDescription: "ResolvConf refers to the DNS resolver configuration",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path defines the path to the file that contains the DNS resolver configuration",
												MarkdownDescription: "Path defines the path to the file that contains the DNS resolver configuration",
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

							"nodes": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cidr_mask_size": schema.Int64Attribute{
										Description:         "CIDRMaskSize defines the mask size for node cidr in the cluster, default for ipv4 is 24. This is an optional field",
										MarkdownDescription: "CIDRMaskSize defines the mask size for node cidr in the cluster, default for ipv4 is 24. This is an optional field",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pods": schema.SingleNestedAttribute{
								Description:         "Comma-separated list of CIDR blocks to use for pod and service subnets. Defaults to 192.168.0.0/16 for pod subnet.",
								MarkdownDescription: "Comma-separated list of CIDR blocks to use for pod and service subnets. Defaults to 192.168.0.0/16 for pod subnet.",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"services": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cidr_blocks": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"control_plane_configuration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cert_sans": schema.ListAttribute{
								Description:         "CertSANs is a slice of domain names or IPs to be added as Subject Name Alternatives of the Kube API Servers Certificate.",
								MarkdownDescription: "CertSANs is a slice of domain names or IPs to be added as Subject Name Alternatives of the Kube API Servers Certificate.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"count": schema.Int64Attribute{
								Description:         "Count defines the number of desired control plane nodes. Defaults to 1.",
								MarkdownDescription: "Count defines the number of desired control plane nodes. Defaults to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.SingleNestedAttribute{
								Description:         "Endpoint defines the host ip and port to use for the control plane.",
								MarkdownDescription: "Endpoint defines the host ip and port to use for the control plane.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host defines the ip that you want to use to connect to the control plane",
										MarkdownDescription: "Host defines the ip that you want to use to connect to the control plane",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels define the labels to assign to the node",
								MarkdownDescription: "Labels define the labels to assign to the node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"machine_group_ref": schema.SingleNestedAttribute{
								Description:         "MachineGroupRef defines the machine group configuration for the control plane.",
								MarkdownDescription: "MachineGroupRef defines the machine group configuration for the control plane.",
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

							"machine_health_check": schema.SingleNestedAttribute{
								Description:         "MachineHealthCheck is a control-plane level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.",
								MarkdownDescription: "MachineHealthCheck is a control-plane level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.",
								Attributes: map[string]schema.Attribute{
									"max_unhealthy": schema.StringAttribute{
										Description:         "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
										MarkdownDescription: "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_startup_timeout": schema.StringAttribute{
										Description:         "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
										MarkdownDescription: "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"unhealthy_machine_timeout": schema.StringAttribute{
										Description:         "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
										MarkdownDescription: "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"skip_load_balancer_deployment": schema.BoolAttribute{
								Description:         "SkipLoadBalancerDeployment skip deploying control plane load balancer. Make sure your infrastructure can handle control plane load balancing when you set this field to true.",
								MarkdownDescription: "SkipLoadBalancerDeployment skip deploying control plane load balancer. Make sure your infrastructure can handle control plane load balancing when you set this field to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"taints": schema.ListNestedAttribute{
								Description:         "Taints define the set of taints to be applied on control plane nodes",
								MarkdownDescription: "Taints define the set of taints to be applied on control plane nodes",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Required. The taint key to be applied to a node.",
											MarkdownDescription: "Required. The taint key to be applied to a node.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"time_added": schema.StringAttribute{
											Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
											MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												validators.DateTime64Validator(),
											},
										},

										"value": schema.StringAttribute{
											Description:         "The taint value corresponding to the taint key.",
											MarkdownDescription: "The taint value corresponding to the taint key.",
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

							"upgrade_rollout_strategy": schema.SingleNestedAttribute{
								Description:         "UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs",
								MarkdownDescription: "UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs",
								Attributes: map[string]schema.Attribute{
									"rolling_update": schema.SingleNestedAttribute{
										Description:         "ControlPlaneRollingUpdateParams is API for rolling update strategy knobs.",
										MarkdownDescription: "ControlPlaneRollingUpdateParams is API for rolling update strategy knobs.",
										Attributes: map[string]schema.Attribute{
											"max_surge": schema.Int64Attribute{
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

									"type": schema.StringAttribute{
										Description:         "UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.",
										MarkdownDescription: "UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.",
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

					"datacenter_ref": schema.SingleNestedAttribute{
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

					"eksa_version": schema.StringAttribute{
						Description:         "EksaVersion is the semver identifying the release of eks-a used to populate the cluster components.",
						MarkdownDescription: "EksaVersion is the semver identifying the release of eks-a used to populate the cluster components.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"etcd_encryption": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"providers": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"kms": schema.SingleNestedAttribute{
												Description:         "KMS defines the configuration for KMS Encryption provider.",
												MarkdownDescription: "KMS defines the configuration for KMS Encryption provider.",
												Attributes: map[string]schema.Attribute{
													"cachesize": schema.Int64Attribute{
														Description:         "CacheSize defines the maximum number of encrypted objects to be cached in memory. The default value is 1000. You can set this to a negative value to disable caching.",
														MarkdownDescription: "CacheSize defines the maximum number of encrypted objects to be cached in memory. The default value is 1000. You can set this to a negative value to disable caching.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name defines the name of KMS plugin to be used.",
														MarkdownDescription: "Name defines the name of KMS plugin to be used.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"socket_listen_address": schema.StringAttribute{
														Description:         "SocketListenAddress defines a UNIX socket address that the KMS provider listens on.",
														MarkdownDescription: "SocketListenAddress defines a UNIX socket address that the KMS provider listens on.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"timeout": schema.StringAttribute{
														Description:         "Timeout for kube-apiserver to wait for KMS plugin. Default is 3s.",
														MarkdownDescription: "Timeout for kube-apiserver to wait for KMS plugin. Default is 3s.",
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"resources": schema.ListAttribute{
									Description:         "Resources defines a list of objects and custom resources definitions that should be encrypted.",
									MarkdownDescription: "Resources defines a list of objects and custom resources definitions that should be encrypted.",
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

					"external_etcd_configuration": schema.SingleNestedAttribute{
						Description:         "ExternalEtcdConfiguration defines the configuration options for using unstacked etcd topology.",
						MarkdownDescription: "ExternalEtcdConfiguration defines the configuration options for using unstacked etcd topology.",
						Attributes: map[string]schema.Attribute{
							"count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"machine_group_ref": schema.SingleNestedAttribute{
								Description:         "MachineGroupRef defines the machine group configuration for the etcd machines.",
								MarkdownDescription: "MachineGroupRef defines the machine group configuration for the etcd machines.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_ops_ref": schema.SingleNestedAttribute{
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

					"identity_provider_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes_version": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"machine_health_check": schema.SingleNestedAttribute{
						Description:         "MachineHealthCheck allows to configure timeouts for machine health checks. Machine Health Checks are responsible for remediating unhealthy Machines. Configuring these values will decide how long to wait to remediate unhealthy machine or determine health of nodes' machines.",
						MarkdownDescription: "MachineHealthCheck allows to configure timeouts for machine health checks. Machine Health Checks are responsible for remediating unhealthy Machines. Configuring these values will decide how long to wait to remediate unhealthy machine or determine health of nodes' machines.",
						Attributes: map[string]schema.Attribute{
							"max_unhealthy": schema.StringAttribute{
								Description:         "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
								MarkdownDescription: "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_startup_timeout": schema.StringAttribute{
								Description:         "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
								MarkdownDescription: "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"unhealthy_machine_timeout": schema.StringAttribute{
								Description:         "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
								MarkdownDescription: "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"management_cluster": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
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

					"packages": schema.SingleNestedAttribute{
						Description:         "PackageConfiguration for installing EKS Anywhere curated packages.",
						MarkdownDescription: "PackageConfiguration for installing EKS Anywhere curated packages.",
						Attributes: map[string]schema.Attribute{
							"controller": schema.SingleNestedAttribute{
								Description:         "Controller package controller configuration",
								MarkdownDescription: "Controller package controller configuration",
								Attributes: map[string]schema.Attribute{
									"digest": schema.StringAttribute{
										Description:         "Digest package controller digest",
										MarkdownDescription: "Digest package controller digest",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_webhooks": schema.BoolAttribute{
										Description:         "DisableWebhooks on package controller",
										MarkdownDescription: "DisableWebhooks on package controller",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.ListAttribute{
										Description:         "Env of package controller in the format 'key=value'",
										MarkdownDescription: "Env of package controller in the format 'key=value'",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository package controller repository",
										MarkdownDescription: "Repository package controller repository",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources of package controller",
										MarkdownDescription: "Resources of package controller",
										Attributes: map[string]schema.Attribute{
											"limits": schema.SingleNestedAttribute{
												Description:         "ImageResource resources for container image.",
												MarkdownDescription: "ImageResource resources for container image.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU image cpu",
														MarkdownDescription: "CPU image cpu",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory image memory",
														MarkdownDescription: "Memory image memory",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": schema.SingleNestedAttribute{
												Description:         "Requests for image resources",
												MarkdownDescription: "Requests for image resources",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU image cpu",
														MarkdownDescription: "CPU image cpu",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory image memory",
														MarkdownDescription: "Memory image memory",
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

									"tag": schema.StringAttribute{
										Description:         "Tag package controller tag",
										MarkdownDescription: "Tag package controller tag",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cronjob": schema.SingleNestedAttribute{
								Description:         "Cronjob for ecr token refresher",
								MarkdownDescription: "Cronjob for ecr token refresher",
								Attributes: map[string]schema.Attribute{
									"digest": schema.StringAttribute{
										Description:         "Digest ecr token refresher digest",
										MarkdownDescription: "Digest ecr token refresher digest",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable": schema.BoolAttribute{
										Description:         "Disable on cron job",
										MarkdownDescription: "Disable on cron job",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository ecr token refresher repository",
										MarkdownDescription: "Repository ecr token refresher repository",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag ecr token refresher tag",
										MarkdownDescription: "Tag ecr token refresher tag",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable": schema.BoolAttribute{
								Description:         "Disable package controller on cluster",
								MarkdownDescription: "Disable package controller on cluster",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_iam_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"service_account_issuer": schema.StringAttribute{
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

					"proxy_configuration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"http_proxy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"https_proxy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_proxy": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"registry_mirror_configuration": schema.SingleNestedAttribute{
						Description:         "RegistryMirrorConfiguration defines the settings for image registry mirror.",
						MarkdownDescription: "RegistryMirrorConfiguration defines the settings for image registry mirror.",
						Attributes: map[string]schema.Attribute{
							"authenticate": schema.BoolAttribute{
								Description:         "Authenticate defines if registry requires authentication",
								MarkdownDescription: "Authenticate defines if registry requires authentication",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_cert_content": schema.StringAttribute{
								Description:         "CACertContent defines the contents registry mirror CA certificate",
								MarkdownDescription: "CACertContent defines the contents registry mirror CA certificate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Endpoint defines the registry mirror endpoint to use for pulling images",
								MarkdownDescription: "Endpoint defines the registry mirror endpoint to use for pulling images",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "InsecureSkipVerify skips the registry certificate verification. Only use this solution for isolated testing or in a tightly controlled, air-gapped environment.",
								MarkdownDescription: "InsecureSkipVerify skips the registry certificate verification. Only use this solution for isolated testing or in a tightly controlled, air-gapped environment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oci_namespaces": schema.ListNestedAttribute{
								Description:         "OCINamespaces defines the mapping from an upstream registry to a local namespace where upstream artifacts are placed into",
								MarkdownDescription: "OCINamespaces defines the mapping from an upstream registry to a local namespace where upstream artifacts are placed into",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Namespace refers to the name of a namespace in the local registry",
											MarkdownDescription: "Namespace refers to the name of a namespace in the local registry",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"registry": schema.StringAttribute{
											Description:         "Name refers to the name of the upstream registry",
											MarkdownDescription: "Name refers to the name of the upstream registry",
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

							"port": schema.StringAttribute{
								Description:         "Port defines the port exposed for registry mirror endpoint",
								MarkdownDescription: "Port defines the port exposed for registry mirror endpoint",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"worker_node_group_configurations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"autoscaling_configuration": schema.SingleNestedAttribute{
									Description:         "AutoScalingConfiguration defines the auto scaling configuration",
									MarkdownDescription: "AutoScalingConfiguration defines the auto scaling configuration",
									Attributes: map[string]schema.Attribute{
										"max_count": schema.Int64Attribute{
											Description:         "MaxCount defines the maximum number of nodes for the associated resource group.",
											MarkdownDescription: "MaxCount defines the maximum number of nodes for the associated resource group.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_count": schema.Int64Attribute{
											Description:         "MinCount defines the minimum number of nodes for the associated resource group.",
											MarkdownDescription: "MinCount defines the minimum number of nodes for the associated resource group.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"count": schema.Int64Attribute{
									Description:         "Count defines the number of desired worker nodes. Defaults to 1.",
									MarkdownDescription: "Count defines the number of desired worker nodes. Defaults to 1.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kubernetes_version": schema.StringAttribute{
									Description:         "KuberenetesVersion defines the version for worker nodes. If not set, the top level spec kubernetesVersion will be used.",
									MarkdownDescription: "KuberenetesVersion defines the version for worker nodes. If not set, the top level spec kubernetesVersion will be used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels define the labels to assign to the node",
									MarkdownDescription: "Labels define the labels to assign to the node",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"machine_group_ref": schema.SingleNestedAttribute{
									Description:         "MachineGroupRef defines the machine group configuration for the worker nodes.",
									MarkdownDescription: "MachineGroupRef defines the machine group configuration for the worker nodes.",
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

								"machine_health_check": schema.SingleNestedAttribute{
									Description:         "MachineHealthCheck is a worker node level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.",
									MarkdownDescription: "MachineHealthCheck is a worker node level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.",
									Attributes: map[string]schema.Attribute{
										"max_unhealthy": schema.StringAttribute{
											Description:         "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
											MarkdownDescription: "MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_startup_timeout": schema.StringAttribute{
											Description:         "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
											MarkdownDescription: "NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"unhealthy_machine_timeout": schema.StringAttribute{
											Description:         "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
											MarkdownDescription: "UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name refers to the name of the worker node group",
									MarkdownDescription: "Name refers to the name of the worker node group",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"taints": schema.ListNestedAttribute{
									Description:         "Taints define the set of taints to be applied on worker nodes",
									MarkdownDescription: "Taints define the set of taints to be applied on worker nodes",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"effect": schema.StringAttribute{
												Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "Required. The taint key to be applied to a node.",
												MarkdownDescription: "Required. The taint key to be applied to a node.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"time_added": schema.StringAttribute{
												Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
												MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.DateTime64Validator(),
												},
											},

											"value": schema.StringAttribute{
												Description:         "The taint value corresponding to the taint key.",
												MarkdownDescription: "The taint value corresponding to the taint key.",
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

								"upgrade_rollout_strategy": schema.SingleNestedAttribute{
									Description:         "UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs",
									MarkdownDescription: "UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs",
									Attributes: map[string]schema.Attribute{
										"rolling_update": schema.SingleNestedAttribute{
											Description:         "WorkerNodesRollingUpdateParams is API for rolling update strategy knobs.",
											MarkdownDescription: "WorkerNodesRollingUpdateParams is API for rolling update strategy knobs.",
											Attributes: map[string]schema.Attribute{
												"max_surge": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"max_unavailable": schema.Int64Attribute{
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

										"type": schema.StringAttribute{
											Description:         "UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.",
											MarkdownDescription: "UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AnywhereEksAmazonawsComClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
