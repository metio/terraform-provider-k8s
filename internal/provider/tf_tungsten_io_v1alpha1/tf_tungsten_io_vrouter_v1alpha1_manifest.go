/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tf_tungsten_io_v1alpha1

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
	_ datasource.DataSource = &TfTungstenIoVrouterV1Alpha1Manifest{}
)

func NewTfTungstenIoVrouterV1Alpha1Manifest() datasource.DataSource {
	return &TfTungstenIoVrouterV1Alpha1Manifest{}
}

type TfTungstenIoVrouterV1Alpha1Manifest struct{}

type TfTungstenIoVrouterV1Alpha1ManifestData struct {
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
		CommonConfiguration *struct {
			AuthParameters *struct {
				AuthMode               *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
				KeystoneAuthParameters *struct {
					Address           *string `tfsdk:"address" json:"address,omitempty"`
					AdminPassword     *string `tfsdk:"admin_password" json:"adminPassword,omitempty"`
					AdminPort         *int64  `tfsdk:"admin_port" json:"adminPort,omitempty"`
					AdminTenant       *string `tfsdk:"admin_tenant" json:"adminTenant,omitempty"`
					AdminUsername     *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
					AuthProtocol      *string `tfsdk:"auth_protocol" json:"authProtocol,omitempty"`
					Insecure          *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					ProjectDomainName *string `tfsdk:"project_domain_name" json:"projectDomainName,omitempty"`
					Region            *string `tfsdk:"region" json:"region,omitempty"`
					UserDomainName    *string `tfsdk:"user_domain_name" json:"userDomainName,omitempty"`
				} `tfsdk:"keystone_auth_parameters" json:"keystoneAuthParameters,omitempty"`
				KeystoneSecretName *string `tfsdk:"keystone_secret_name" json:"keystoneSecretName,omitempty"`
			} `tfsdk:"auth_parameters" json:"authParameters,omitempty"`
			Distribution     *string            `tfsdk:"distribution" json:"distribution,omitempty"`
			ImagePullSecrets *[]string          `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
			NodeSelector     *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations      *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"common_configuration" json:"commonConfiguration,omitempty"`
		ServiceConfiguration *struct {
			AgentMode                 *string `tfsdk:"agent_mode" json:"agentMode,omitempty"`
			BarbicanPassword          *string `tfsdk:"barbican_password" json:"barbicanPassword,omitempty"`
			BarbicanTenantName        *string `tfsdk:"barbican_tenant_name" json:"barbicanTenantName,omitempty"`
			BarbicanUser              *string `tfsdk:"barbican_user" json:"barbicanUser,omitempty"`
			CloudOrchestrator         *string `tfsdk:"cloud_orchestrator" json:"cloudOrchestrator,omitempty"`
			CniMTU                    *int64  `tfsdk:"cni_mtu" json:"cniMTU,omitempty"`
			CollectorPort             *string `tfsdk:"collector_port" json:"collectorPort,omitempty"`
			ConfigApiPort             *string `tfsdk:"config_api_port" json:"configApiPort,omitempty"`
			ConfigApiServerCaCertfile *string `tfsdk:"config_api_server_ca_certfile" json:"configApiServerCaCertfile,omitempty"`
			ConfigApiSslEnable        *bool   `tfsdk:"config_api_ssl_enable" json:"configApiSslEnable,omitempty"`
			Containers                *[]struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Image   *string   `tfsdk:"image" json:"image,omitempty"`
				Name    *string   `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"containers" json:"containers,omitempty"`
			ControlInstance                 *string            `tfsdk:"control_instance" json:"controlInstance,omitempty"`
			DataSubnet                      *string            `tfsdk:"data_subnet" json:"dataSubnet,omitempty"`
			DnsServerPort                   *string            `tfsdk:"dns_server_port" json:"dnsServerPort,omitempty"`
			DpdkUioDriver                   *string            `tfsdk:"dpdk_uio_driver" json:"dpdkUioDriver,omitempty"`
			EnvVariablesConfig              *map[string]string `tfsdk:"env_variables_config" json:"envVariablesConfig,omitempty"`
			FabricSntHashTableSize          *string            `tfsdk:"fabric_snt_hash_table_size" json:"fabricSntHashTableSize,omitempty"`
			HugePages1G                     *int64             `tfsdk:"huge_pages1_g" json:"hugePages1G,omitempty"`
			HugePages2M                     *int64             `tfsdk:"huge_pages2_m" json:"hugePages2M,omitempty"`
			HypervisorType                  *string            `tfsdk:"hypervisor_type" json:"hypervisorType,omitempty"`
			IntrospectSslEnable             *bool              `tfsdk:"introspect_ssl_enable" json:"introspectSslEnable,omitempty"`
			K8sToken                        *string            `tfsdk:"k8s_token" json:"k8sToken,omitempty"`
			K8sTokenFile                    *string            `tfsdk:"k8s_token_file" json:"k8sTokenFile,omitempty"`
			KeystoneAuthAdminPassword       *string            `tfsdk:"keystone_auth_admin_password" json:"keystoneAuthAdminPassword,omitempty"`
			KeystoneAuthAdminPort           *string            `tfsdk:"keystone_auth_admin_port" json:"keystoneAuthAdminPort,omitempty"`
			KeystoneAuthCaCertfile          *string            `tfsdk:"keystone_auth_ca_certfile" json:"keystoneAuthCaCertfile,omitempty"`
			KeystoneAuthCertfile            *string            `tfsdk:"keystone_auth_certfile" json:"keystoneAuthCertfile,omitempty"`
			KeystoneAuthHost                *string            `tfsdk:"keystone_auth_host" json:"keystoneAuthHost,omitempty"`
			KeystoneAuthInsecure            *bool              `tfsdk:"keystone_auth_insecure" json:"keystoneAuthInsecure,omitempty"`
			KeystoneAuthKeyfile             *string            `tfsdk:"keystone_auth_keyfile" json:"keystoneAuthKeyfile,omitempty"`
			KeystoneAuthProjectDomainName   *string            `tfsdk:"keystone_auth_project_domain_name" json:"keystoneAuthProjectDomainName,omitempty"`
			KeystoneAuthProto               *string            `tfsdk:"keystone_auth_proto" json:"keystoneAuthProto,omitempty"`
			KeystoneAuthRegionName          *string            `tfsdk:"keystone_auth_region_name" json:"keystoneAuthRegionName,omitempty"`
			KeystoneAuthUrlTokens           *string            `tfsdk:"keystone_auth_url_tokens" json:"keystoneAuthUrlTokens,omitempty"`
			KeystoneAuthUrlVersion          *string            `tfsdk:"keystone_auth_url_version" json:"keystoneAuthUrlVersion,omitempty"`
			KeystoneAuthUserDomainName      *string            `tfsdk:"keystone_auth_user_domain_name" json:"keystoneAuthUserDomainName,omitempty"`
			KubernetesApiPort               *string            `tfsdk:"kubernetes_api_port" json:"kubernetesApiPort,omitempty"`
			KubernetesApiSecurePort         *string            `tfsdk:"kubernetes_api_secure_port" json:"kubernetesApiSecurePort,omitempty"`
			KubernetesPodSubnet             *string            `tfsdk:"kubernetes_pod_subnet" json:"kubernetesPodSubnet,omitempty"`
			L3mhCidr                        *string            `tfsdk:"l3mh_cidr" json:"l3mhCidr,omitempty"`
			LogDir                          *string            `tfsdk:"log_dir" json:"logDir,omitempty"`
			LogLocal                        *int64             `tfsdk:"log_local" json:"logLocal,omitempty"`
			MetadataProxySecret             *string            `tfsdk:"metadata_proxy_secret" json:"metadataProxySecret,omitempty"`
			MetadataSslCaCertfile           *string            `tfsdk:"metadata_ssl_ca_certfile" json:"metadataSslCaCertfile,omitempty"`
			MetadataSslCertType             *string            `tfsdk:"metadata_ssl_cert_type" json:"metadataSslCertType,omitempty"`
			MetadataSslCertfile             *string            `tfsdk:"metadata_ssl_certfile" json:"metadataSslCertfile,omitempty"`
			MetadataSslEnable               *string            `tfsdk:"metadata_ssl_enable" json:"metadataSslEnable,omitempty"`
			MetadataSslKeyfile              *string            `tfsdk:"metadata_ssl_keyfile" json:"metadataSslKeyfile,omitempty"`
			PhysicalInterface               *string            `tfsdk:"physical_interface" json:"physicalInterface,omitempty"`
			PriorityBandwidth               *string            `tfsdk:"priority_bandwidth" json:"priorityBandwidth,omitempty"`
			PriorityId                      *string            `tfsdk:"priority_id" json:"priorityId,omitempty"`
			PriorityScheduling              *string            `tfsdk:"priority_scheduling" json:"priorityScheduling,omitempty"`
			PriorityTagging                 *bool              `tfsdk:"priority_tagging" json:"priorityTagging,omitempty"`
			QosDefHwQueue                   *bool              `tfsdk:"qos_def_hw_queue" json:"qosDefHwQueue,omitempty"`
			QosLogicalQueues                *string            `tfsdk:"qos_logical_queues" json:"qosLogicalQueues,omitempty"`
			QosQueueId                      *string            `tfsdk:"qos_queue_id" json:"qosQueueId,omitempty"`
			RequiredKernelVrouterEncryption *string            `tfsdk:"required_kernel_vrouter_encryption" json:"requiredKernelVrouterEncryption,omitempty"`
			SampleDestination               *string            `tfsdk:"sample_destination" json:"sampleDestination,omitempty"`
			SandeshCaCertfile               *string            `tfsdk:"sandesh_ca_certfile" json:"sandeshCaCertfile,omitempty"`
			SandeshCertfile                 *string            `tfsdk:"sandesh_certfile" json:"sandeshCertfile,omitempty"`
			SandeshKeyfile                  *string            `tfsdk:"sandesh_keyfile" json:"sandeshKeyfile,omitempty"`
			SandeshServerCertfile           *string            `tfsdk:"sandesh_server_certfile" json:"sandeshServerCertfile,omitempty"`
			SandeshServerKeyfile            *string            `tfsdk:"sandesh_server_keyfile" json:"sandeshServerKeyfile,omitempty"`
			SandeshSslEnable                *bool              `tfsdk:"sandesh_ssl_enable" json:"sandeshSslEnable,omitempty"`
			ServerCaCertfile                *string            `tfsdk:"server_ca_certfile" json:"serverCaCertfile,omitempty"`
			ServerCertfile                  *string            `tfsdk:"server_certfile" json:"serverCertfile,omitempty"`
			ServerKeyfile                   *string            `tfsdk:"server_keyfile" json:"serverKeyfile,omitempty"`
			SloDestination                  *string            `tfsdk:"slo_destination" json:"sloDestination,omitempty"`
			SriovPhysicalInterface          *string            `tfsdk:"sriov_physical_interface" json:"sriovPhysicalInterface,omitempty"`
			SriovPhysicalNetwork            *string            `tfsdk:"sriov_physical_network" json:"sriovPhysicalNetwork,omitempty"`
			SriovVf                         *string            `tfsdk:"sriov_vf" json:"sriovVf,omitempty"`
			SslEnable                       *bool              `tfsdk:"ssl_enable" json:"sslEnable,omitempty"`
			SslInsecure                     *bool              `tfsdk:"ssl_insecure" json:"sslInsecure,omitempty"`
			StatsCollectorDestinationPath   *string            `tfsdk:"stats_collector_destination_path" json:"statsCollectorDestinationPath,omitempty"`
			Subcluster                      *string            `tfsdk:"subcluster" json:"subcluster,omitempty"`
			TsnAgentMode                    *string            `tfsdk:"tsn_agent_mode" json:"tsnAgentMode,omitempty"`
			VrouterCryptInterface           *string            `tfsdk:"vrouter_crypt_interface" json:"vrouterCryptInterface,omitempty"`
			VrouterDecryptInterface         *string            `tfsdk:"vrouter_decrypt_interface" json:"vrouterDecryptInterface,omitempty"`
			VrouterDecryptKey               *string            `tfsdk:"vrouter_decrypt_key" json:"vrouterDecryptKey,omitempty"`
			VrouterEncryption               *bool              `tfsdk:"vrouter_encryption" json:"vrouterEncryption,omitempty"`
			VrouterGateway                  *string            `tfsdk:"vrouter_gateway" json:"vrouterGateway,omitempty"`
			XmmpSslEnable                   *bool              `tfsdk:"xmmp_ssl_enable" json:"xmmpSslEnable,omitempty"`
			XmppServerCaCertfile            *string            `tfsdk:"xmpp_server_ca_certfile" json:"xmppServerCaCertfile,omitempty"`
			XmppServerCertfile              *string            `tfsdk:"xmpp_server_certfile" json:"xmppServerCertfile,omitempty"`
			XmppServerKeyfile               *string            `tfsdk:"xmpp_server_keyfile" json:"xmppServerKeyfile,omitempty"`
			XmppServerPort                  *string            `tfsdk:"xmpp_server_port" json:"xmppServerPort,omitempty"`
		} `tfsdk:"service_configuration" json:"serviceConfiguration,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TfTungstenIoVrouterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tf_tungsten_io_vrouter_v1alpha1_manifest"
}

func (r *TfTungstenIoVrouterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Vrouter is the Schema for the vrouters API.",
		MarkdownDescription: "Vrouter is the Schema for the vrouters API.",
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
				Description:         "VrouterSpec is the Spec for the vrouter API.",
				MarkdownDescription: "VrouterSpec is the Spec for the vrouter API.",
				Attributes: map[string]schema.Attribute{
					"common_configuration": schema.SingleNestedAttribute{
						Description:         "PodConfiguration is the common services struct.",
						MarkdownDescription: "PodConfiguration is the common services struct.",
						Attributes: map[string]schema.Attribute{
							"auth_parameters": schema.SingleNestedAttribute{
								Description:         "AuthParameters auth parameters",
								MarkdownDescription: "AuthParameters auth parameters",
								Attributes: map[string]schema.Attribute{
									"auth_mode": schema.StringAttribute{
										Description:         "AuthenticationMode auth mode",
										MarkdownDescription: "AuthenticationMode auth mode",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("noauth", "keystone"),
										},
									},

									"keystone_auth_parameters": schema.SingleNestedAttribute{
										Description:         "KeystoneAuthParameters keystone parameters",
										MarkdownDescription: "KeystoneAuthParameters keystone parameters",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_password": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_tenant": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"admin_username": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auth_protocol": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"project_domain_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_domain_name": schema.StringAttribute{
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

									"keystone_secret_name": schema.StringAttribute{
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

							"distribution": schema.StringAttribute{
								Description:         "OS family",
								MarkdownDescription: "OS family",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListAttribute{
								Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Kubernetes Cluster Configuration",
								MarkdownDescription: "Kubernetes Cluster Configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("info", "debug", "warning", "error", "critical", "none"),
								},
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "If specified, the pod's tolerations.",
								MarkdownDescription: "If specified, the pod's tolerations.",
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

					"service_configuration": schema.SingleNestedAttribute{
						Description:         "VrouterConfiguration is the Spec for the vrouter API.",
						MarkdownDescription: "VrouterConfiguration is the Spec for the vrouter API.",
						Attributes: map[string]schema.Attribute{
							"agent_mode": schema.StringAttribute{
								Description:         "vRouter",
								MarkdownDescription: "vRouter",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"barbican_password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"barbican_tenant_name": schema.StringAttribute{
								Description:         "Openstack",
								MarkdownDescription: "Openstack",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"barbican_user": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cloud_orchestrator": schema.StringAttribute{
								Description:         "New params for vrouter configuration",
								MarkdownDescription: "New params for vrouter configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cni_mtu": schema.Int64Attribute{
								Description:         "CniMTU - mtu for virtual tap devices",
								MarkdownDescription: "CniMTU - mtu for virtual tap devices",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"collector_port": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_api_port": schema.StringAttribute{
								Description:         "Config",
								MarkdownDescription: "Config",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_api_server_ca_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_api_ssl_enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
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

							"control_instance": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data_subnet": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_server_port": schema.StringAttribute{
								Description:         "DNS",
								MarkdownDescription: "DNS",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dpdk_uio_driver": schema.StringAttribute{
								Description:         "Host",
								MarkdownDescription: "Host",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_variables_config": schema.MapAttribute{
								Description:         "What is it doing? VrouterEncryption bool 'json:'vrouterEncryption,omitempty'' What is it doing? What is it doing?",
								MarkdownDescription: "What is it doing? VrouterEncryption bool 'json:'vrouterEncryption,omitempty'' What is it doing? What is it doing?",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fabric_snt_hash_table_size": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"huge_pages1_g": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"huge_pages2_m": schema.Int64Attribute{
								Description:         "HugePages",
								MarkdownDescription: "HugePages",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hypervisor_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"introspect_ssl_enable": schema.BoolAttribute{
								Description:         "Introspect",
								MarkdownDescription: "Introspect",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"k8s_token": schema.StringAttribute{
								Description:         "Kubernetes",
								MarkdownDescription: "Kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"k8s_token_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_admin_password": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_admin_port": schema.StringAttribute{
								Description:         "Keystone authentication",
								MarkdownDescription: "Keystone authentication",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_ca_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_host": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_insecure": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_project_domain_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_proto": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_region_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_url_tokens": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_url_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keystone_auth_user_domain_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_api_port": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_api_secure_port": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_pod_subnet": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"l3mh_cidr": schema.StringAttribute{
								Description:         "L3MH",
								MarkdownDescription: "L3MH",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_dir": schema.StringAttribute{
								Description:         "Logging",
								MarkdownDescription: "Logging",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_local": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_proxy_secret": schema.StringAttribute{
								Description:         "Metadata",
								MarkdownDescription: "Metadata",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_ssl_ca_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_ssl_cert_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_ssl_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_ssl_enable": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata_ssl_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"physical_interface": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_bandwidth": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_scheduling": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_tagging": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qos_def_hw_queue": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qos_logical_queues": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qos_queue_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"required_kernel_vrouter_encryption": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sample_destination": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_ca_certfile": schema.StringAttribute{
								Description:         "Sandesh",
								MarkdownDescription: "Sandesh",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_server_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_server_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sandesh_ssl_enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_ca_certfile": schema.StringAttribute{
								Description:         "Server SSL",
								MarkdownDescription: "Server SSL",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"slo_destination": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sriov_physical_interface": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sriov_physical_network": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sriov_vf": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_insecure": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stats_collector_destination_path": schema.StringAttribute{
								Description:         "Collector",
								MarkdownDescription: "Collector",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subcluster": schema.StringAttribute{
								Description:         "XMPP",
								MarkdownDescription: "XMPP",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tsn_agent_mode": schema.StringAttribute{
								Description:         "TSN",
								MarkdownDescription: "TSN",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vrouter_crypt_interface": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vrouter_decrypt_interface": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vrouter_decrypt_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vrouter_encryption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vrouter_gateway": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xmmp_ssl_enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xmpp_server_ca_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xmpp_server_certfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xmpp_server_keyfile": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xmpp_server_port": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *TfTungstenIoVrouterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tf_tungsten_io_vrouter_v1alpha1_manifest")

	var model TfTungstenIoVrouterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tf.tungsten.io/v1alpha1")
	model.Kind = pointer.String("Vrouter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
