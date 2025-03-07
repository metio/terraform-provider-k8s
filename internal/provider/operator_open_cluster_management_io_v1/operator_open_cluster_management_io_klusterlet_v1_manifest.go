/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_open_cluster_management_io_v1

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
	_ datasource.DataSource = &OperatorOpenClusterManagementIoKlusterletV1Manifest{}
)

func NewOperatorOpenClusterManagementIoKlusterletV1Manifest() datasource.DataSource {
	return &OperatorOpenClusterManagementIoKlusterletV1Manifest{}
}

type OperatorOpenClusterManagementIoKlusterletV1Manifest struct{}

type OperatorOpenClusterManagementIoKlusterletV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterName  *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		DeployOption *struct {
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"deploy_option" json:"deployOption,omitempty"`
		ExternalServerURLs *[]struct {
			CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			Url      *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"external_server_urls" json:"externalServerURLs,omitempty"`
		HubApiServerHostAlias *struct {
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip       *string `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"hub_api_server_host_alias" json:"hubApiServerHostAlias,omitempty"`
		ImagePullSpec *string `tfsdk:"image_pull_spec" json:"imagePullSpec,omitempty"`
		Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
		NodePlacement *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
		PriorityClassName         *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		RegistrationConfiguration *struct {
			BootstrapKubeConfigs *struct {
				LocalSecretsConfig *struct {
					HubConnectionTimeoutSeconds *int64 `tfsdk:"hub_connection_timeout_seconds" json:"hubConnectionTimeoutSeconds,omitempty"`
					KubeConfigSecrets           *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"kube_config_secrets" json:"kubeConfigSecrets,omitempty"`
				} `tfsdk:"local_secrets_config" json:"localSecretsConfig,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"bootstrap_kube_configs" json:"bootstrapKubeConfigs,omitempty"`
			ClientCertExpirationSeconds *int64             `tfsdk:"client_cert_expiration_seconds" json:"clientCertExpirationSeconds,omitempty"`
			ClusterAnnotations          *map[string]string `tfsdk:"cluster_annotations" json:"clusterAnnotations,omitempty"`
			FeatureGates                *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
			KubeAPIBurst       *int64 `tfsdk:"kube_api_burst" json:"kubeAPIBurst,omitempty"`
			KubeAPIQPS         *int64 `tfsdk:"kube_apiqps" json:"kubeAPIQPS,omitempty"`
			RegistrationDriver *struct {
				AuthType *string `tfsdk:"auth_type" json:"authType,omitempty"`
				AwsIrsa  *struct {
					HubClusterArn     *string `tfsdk:"hub_cluster_arn" json:"hubClusterArn,omitempty"`
					ManagedClusterArn *string `tfsdk:"managed_cluster_arn" json:"managedClusterArn,omitempty"`
				} `tfsdk:"aws_irsa" json:"awsIrsa,omitempty"`
			} `tfsdk:"registration_driver" json:"registrationDriver,omitempty"`
		} `tfsdk:"registration_configuration" json:"registrationConfiguration,omitempty"`
		RegistrationImagePullSpec *string `tfsdk:"registration_image_pull_spec" json:"registrationImagePullSpec,omitempty"`
		ResourceRequirement       *struct {
			ResourceRequirements *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"resource_requirement" json:"resourceRequirement,omitempty"`
		WorkConfiguration *struct {
			AppliedManifestWorkEvictionGracePeriod *string `tfsdk:"applied_manifest_work_eviction_grace_period" json:"appliedManifestWorkEvictionGracePeriod,omitempty"`
			FeatureGates                           *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
			KubeAPIBurst *int64 `tfsdk:"kube_api_burst" json:"kubeAPIBurst,omitempty"`
			KubeAPIQPS   *int64 `tfsdk:"kube_apiqps" json:"kubeAPIQPS,omitempty"`
		} `tfsdk:"work_configuration" json:"workConfiguration,omitempty"`
		WorkImagePullSpec *string `tfsdk:"work_image_pull_spec" json:"workImagePullSpec,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_open_cluster_management_io_klusterlet_v1_manifest"
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Klusterlet represents controllers to install the resources for a managed cluster. When configured, the Klusterlet requires a secret named bootstrap-hub-kubeconfig in the agent namespace to allow API requests to the hub for the registration protocol. In Hosted mode, the Klusterlet requires an additional secret named external-managed-kubeconfig in the agent namespace to allow API requests to the managed cluster for resources installation.",
		MarkdownDescription: "Klusterlet represents controllers to install the resources for a managed cluster. When configured, the Klusterlet requires a secret named bootstrap-hub-kubeconfig in the agent namespace to allow API requests to the hub for the registration protocol. In Hosted mode, the Klusterlet requires an additional secret named external-managed-kubeconfig in the agent namespace to allow API requests to the managed cluster for resources installation.",
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
				Description:         "Spec represents the desired deployment configuration of Klusterlet agent.",
				MarkdownDescription: "Spec represents the desired deployment configuration of Klusterlet agent.",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the name of the managed cluster to be created on hub. The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.",
						MarkdownDescription: "ClusterName is the name of the managed cluster to be created on hub. The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
						},
					},

					"deploy_option": schema.SingleNestedAttribute{
						Description:         "DeployOption contains the options of deploying a klusterlet",
						MarkdownDescription: "DeployOption contains the options of deploying a klusterlet",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode can be Default, Hosted, Singleton or SingletonHosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). In Singleton mode, registration/work agent is started as a single deployment. In SingletonHosted mode, agent is started as a single deployment in hosted mode. Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default, Hosted, Singleton or SingletonHosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). In Singleton mode, registration/work agent is started as a single deployment. In SingletonHosted mode, agent is started as a single deployment in hosted mode. Note: Do not modify the Mode field once it's applied.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_server_urls": schema.ListNestedAttribute{
						Description:         "ExternalServerURLs represents a list of apiserver urls and ca bundles that is accessible externally If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.",
						MarkdownDescription: "ExternalServerURLs represents a list of apiserver urls and ca bundles that is accessible externally If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ca_bundle": schema.StringAttribute{
									Description:         "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
									MarkdownDescription: "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"url": schema.StringAttribute{
									Description:         "URL is the url of apiserver endpoint of the managed cluster.",
									MarkdownDescription: "URL is the url of apiserver endpoint of the managed cluster.",
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

					"hub_api_server_host_alias": schema.SingleNestedAttribute{
						Description:         "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						MarkdownDescription: "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "Hostname for the above IP address.",
								MarkdownDescription: "Hostname for the above IP address.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
								},
							},

							"ip": schema.StringAttribute{
								Description:         "IP address of the host file entry.",
								MarkdownDescription: "IP address of the host file entry.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_spec": schema.StringAttribute{
						Description:         "ImagePullSpec represents the desired image configuration of agent, it takes effect only when singleton mode is set. quay.io/open-cluster-management.io/registration-operator:latest will be used if unspecified",
						MarkdownDescription: "ImagePullSpec represents the desired image configuration of agent, it takes effect only when singleton mode is set. quay.io/open-cluster-management.io/registration-operator:latest will be used if unspecified",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						MarkdownDescription: "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(57),
							stringvalidator.RegexMatches(regexp.MustCompile(`^open-cluster-management-[-a-z0-9]*[a-z0-9]$`), ""),
						},
					},

					"node_placement": schema.SingleNestedAttribute{
						Description:         "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						MarkdownDescription: "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								MarkdownDescription: "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations are attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
								MarkdownDescription: "Tolerations are attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName is the name of the PriorityClass that will be used by the deployed klusterlet agent. It will be ignored when the PriorityClass/v1 API is not available on the managed cluster.",
						MarkdownDescription: "PriorityClassName is the name of the PriorityClass that will be used by the deployed klusterlet agent. It will be ignored when the PriorityClass/v1 API is not available on the managed cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"registration_configuration": schema.SingleNestedAttribute{
						Description:         "RegistrationConfiguration contains the configuration of registration",
						MarkdownDescription: "RegistrationConfiguration contains the configuration of registration",
						Attributes: map[string]schema.Attribute{
							"bootstrap_kube_configs": schema.SingleNestedAttribute{
								Description:         "BootstrapKubeConfigs defines the ordered list of bootstrap kubeconfigs. The order decides which bootstrap kubeconfig to use first when rebootstrap. When the agent loses the connection to the current hub over HubConnectionTimeoutSeconds, or the managedcluster CR is set 'hubAcceptsClient=false' on the hub, the controller marks the related bootstrap kubeconfig as 'failed'. A failed bootstrapkubeconfig won't be used for the duration specified by SkipFailedBootstrapKubeConfigSeconds. But if the user updates the content of a failed bootstrapkubeconfig, the 'failed' mark will be cleared.",
								MarkdownDescription: "BootstrapKubeConfigs defines the ordered list of bootstrap kubeconfigs. The order decides which bootstrap kubeconfig to use first when rebootstrap. When the agent loses the connection to the current hub over HubConnectionTimeoutSeconds, or the managedcluster CR is set 'hubAcceptsClient=false' on the hub, the controller marks the related bootstrap kubeconfig as 'failed'. A failed bootstrapkubeconfig won't be used for the duration specified by SkipFailedBootstrapKubeConfigSeconds. But if the user updates the content of a failed bootstrapkubeconfig, the 'failed' mark will be cleared.",
								Attributes: map[string]schema.Attribute{
									"local_secrets_config": schema.SingleNestedAttribute{
										Description:         "LocalSecretsConfig include a list of secrets that contains the kubeconfigs for ordered bootstrap kubeconifigs. The secrets must be in the same namespace where the agent controller runs.",
										MarkdownDescription: "LocalSecretsConfig include a list of secrets that contains the kubeconfigs for ordered bootstrap kubeconifigs. The secrets must be in the same namespace where the agent controller runs.",
										Attributes: map[string]schema.Attribute{
											"hub_connection_timeout_seconds": schema.Int64Attribute{
												Description:         "HubConnectionTimeoutSeconds is used to set the timeout of connecting to the hub cluster. When agent loses the connection to the hub over the timeout seconds, the agent do a rebootstrap. By default is 10 mins.",
												MarkdownDescription: "HubConnectionTimeoutSeconds is used to set the timeout of connecting to the hub cluster. When agent loses the connection to the hub over the timeout seconds, the agent do a rebootstrap. By default is 10 mins.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(180),
												},
											},

											"kube_config_secrets": schema.ListNestedAttribute{
												Description:         "KubeConfigSecrets is a list of secret names. The secrets are in the same namespace where the agent controller runs.",
												MarkdownDescription: "KubeConfigSecrets is a list of secret names. The secrets are in the same namespace where the agent controller runs.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the secret.",
															MarkdownDescription: "Name is the name of the secret.",
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

									"type": schema.StringAttribute{
										Description:         "Type specifies the type of priority bootstrap kubeconfigs. By default, it is set to None, representing no priority bootstrap kubeconfigs are set.",
										MarkdownDescription: "Type specifies the type of priority bootstrap kubeconfigs. By default, it is set to None, representing no priority bootstrap kubeconfigs are set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("None", "LocalSecrets"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"client_cert_expiration_seconds": schema.Int64Attribute{
								Description:         "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								MarkdownDescription: "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_annotations": schema.MapAttribute{
								Description:         "ClusterAnnotations is annotations with the reserve prefix 'agent.open-cluster-management.io' set on ManagedCluster when creating only, other actors can update it afterwards.",
								MarkdownDescription: "ClusterAnnotations is annotations with the reserve prefix 'agent.open-cluster-management.io' set on ManagedCluster when creating only, other actors can update it afterwards.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kube_api_burst": schema.Int64Attribute{
								Description:         "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
								MarkdownDescription: "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_apiqps": schema.Int64Attribute{
								Description:         "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
								MarkdownDescription: "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"registration_driver": schema.SingleNestedAttribute{
								Description:         "This provides driver details required to register with hub",
								MarkdownDescription: "This provides driver details required to register with hub",
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Description:         "Type of the authentication used by managedcluster to register as well as pull work from hub. Possible values are csr and awsirsa.",
										MarkdownDescription: "Type of the authentication used by managedcluster to register as well as pull work from hub. Possible values are csr and awsirsa.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("csr", "awsirsa"),
										},
									},

									"aws_irsa": schema.SingleNestedAttribute{
										Description:         "Contain the details required for registering with hub cluster (ie: an EKS cluster) using AWS IAM roles for service account. This is required only when the authType is awsirsa.",
										MarkdownDescription: "Contain the details required for registering with hub cluster (ie: an EKS cluster) using AWS IAM roles for service account. This is required only when the authType is awsirsa.",
										Attributes: map[string]schema.Attribute{
											"hub_cluster_arn": schema.StringAttribute{
												Description:         "The arn of the hub cluster (ie: an EKS cluster). This will be required to pass information to hub, which hub will use to create IAM identities for this klusterlet. Example - arn:eks:us-west-2:12345678910:cluster/hub-cluster1.",
												MarkdownDescription: "The arn of the hub cluster (ie: an EKS cluster). This will be required to pass information to hub, which hub will use to create IAM identities for this klusterlet. Example - arn:eks:us-west-2:12345678910:cluster/hub-cluster1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"managed_cluster_arn": schema.StringAttribute{
												Description:         "The arn of the managed cluster (ie: an EKS cluster). This will be required to generate the md5hash which will be used as a suffix to create IAM role on hub as well as used by kluslerlet-agent, to assume role suffixed with the md5hash, on startup. Example - arn:eks:us-west-2:12345678910:cluster/managed-cluster1.",
												MarkdownDescription: "The arn of the managed cluster (ie: an EKS cluster). This will be required to generate the md5hash which will be used as a suffix to create IAM role on hub as well as used by kluslerlet-agent, to assume role suffixed with the md5hash, on startup. Example - arn:eks:us-west-2:12345678910:cluster/managed-cluster1.",
												Required:            false,
												Optional:            true,
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

					"registration_image_pull_spec": schema.StringAttribute{
						Description:         "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_requirement": schema.SingleNestedAttribute{
						Description:         "ResourceRequirement specify QoS classes of deployments managed by klusterlet. It applies to all the containers in the deployments.",
						MarkdownDescription: "ResourceRequirement specify QoS classes of deployments managed by klusterlet. It applies to all the containers in the deployments.",
						Attributes: map[string]schema.Attribute{
							"resource_requirements": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements defines resource requests and limits when Type is ResourceQosClassResourceRequirement",
								MarkdownDescription: "ResourceRequirements defines resource requests and limits when Type is ResourceQosClassResourceRequirement",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
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

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Default", "BestEffort", "ResourceRequirement"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"work_configuration": schema.SingleNestedAttribute{
						Description:         "WorkConfiguration contains the configuration of work",
						MarkdownDescription: "WorkConfiguration contains the configuration of work",
						Attributes: map[string]schema.Attribute{
							"applied_manifest_work_eviction_grace_period": schema.StringAttribute{
								Description:         "AppliedManifestWorkEvictionGracePeriod is the eviction grace period the work agent will wait before evicting the AppliedManifestWorks, whose corresponding ManifestWorks are missing on the hub cluster, from the managed cluster. If not present, the default value of the work agent will be used.",
								MarkdownDescription: "AppliedManifestWorkEvictionGracePeriod is the eviction grace period the work agent will wait before evicting the AppliedManifestWorks, whose corresponding ManifestWorks are missing on the hub cluster, from the managed cluster. If not present, the default value of the work agent will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(s|m|h))+$`), ""),
								},
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kube_api_burst": schema.Int64Attribute{
								Description:         "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
								MarkdownDescription: "KubeAPIBurst indicates the maximum burst of the throttle while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 100",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_apiqps": schema.Int64Attribute{
								Description:         "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
								MarkdownDescription: "KubeAPIQPS indicates the maximum QPS while talking with apiserver of hub cluster from the spoke cluster. If it is set empty, use the default value: 50",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"work_image_pull_spec": schema.StringAttribute{
						Description:         "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
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

func (r *OperatorOpenClusterManagementIoKlusterletV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_open_cluster_management_io_klusterlet_v1_manifest")

	var model OperatorOpenClusterManagementIoKlusterletV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	model.Kind = pointer.String("Klusterlet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
