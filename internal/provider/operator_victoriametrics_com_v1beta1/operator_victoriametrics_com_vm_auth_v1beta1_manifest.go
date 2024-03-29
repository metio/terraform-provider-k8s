/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

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
	_ datasource.DataSource = &OperatorVictoriametricsComVmauthV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmauthV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmauthV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmauthV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmauthV1Beta1ManifestData struct {
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
		Affinity     *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
		ConfigMaps   *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
		ConfigSecret *string              `tfsdk:"config_secret" json:"configSecret,omitempty"`
		Containers   *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
		DnsConfig    *struct {
			Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
			Options     *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
		} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
		DnsPolicy   *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
		ExtraArgs   *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
		ExtraEnvs   *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
		HostAliases *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
		HostNetwork *bool `tfsdk:"host_network" json:"hostNetwork,omitempty"`
		Image       *struct {
			PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		Ingress *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Class_name  *string            `tfsdk:"class_name" json:"class_name,omitempty"`
			ExtraRules  *[]struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Http *struct {
					Paths *[]struct {
						Backend *struct {
							Resource *struct {
								ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
								Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"resource" json:"resource,omitempty"`
							Service *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Port *struct {
									Name   *string `tfsdk:"name" json:"name,omitempty"`
									Number *int64  `tfsdk:"number" json:"number,omitempty"`
								} `tfsdk:"port" json:"port,omitempty"`
							} `tfsdk:"service" json:"service,omitempty"`
						} `tfsdk:"backend" json:"backend,omitempty"`
						Path     *string `tfsdk:"path" json:"path,omitempty"`
						PathType *string `tfsdk:"path_type" json:"pathType,omitempty"`
					} `tfsdk:"paths" json:"paths,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
			} `tfsdk:"extra_rules" json:"extraRules,omitempty"`
			ExtraTls *[]struct {
				Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"extra_tls" json:"extraTls,omitempty"`
			Host          *string            `tfsdk:"host" json:"host,omitempty"`
			Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name          *string            `tfsdk:"name" json:"name,omitempty"`
			TlsHosts      *[]string          `tfsdk:"tls_hosts" json:"tlsHosts,omitempty"`
			TlsSecretName *string            `tfsdk:"tls_secret_name" json:"tlsSecretName,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
		License        *struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			KeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		} `tfsdk:"license" json:"license,omitempty"`
		LivenessProbe       *map[string]string `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		LogFormat           *string            `tfsdk:"log_format" json:"logFormat,omitempty"`
		LogLevel            *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		MinReadySeconds     *int64             `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
		NodeSelector        *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PodDisruptionBudget *struct {
			MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
			SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
		} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		PodSecurityPolicyName *string `tfsdk:"pod_security_policy_name" json:"podSecurityPolicyName,omitempty"`
		Port                  *string `tfsdk:"port" json:"port,omitempty"`
		PriorityClassName     *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ReadinessGates        *[]struct {
			ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
		} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
		ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
		Resources      *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RevisionHistoryLimitCount *int64             `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
		RuntimeClassName          *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
		SchedulerName             *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Secrets                   *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
		SecurityContext           *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
		SelectAllByDefault        *bool              `tfsdk:"select_all_by_default" json:"selectAllByDefault,omitempty"`
		ServiceAccountName        *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		ServiceScrapeSpec         *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
		ServiceSpec               *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
		StartupProbe                  *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
		TerminationGracePeriodSeconds *int64             `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
		Tolerations                   *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		UnauthorizedAccessConfig  *[]struct {
			Drop_src_path_prefix_parts *int64    `tfsdk:"drop_src_path_prefix_parts" json:"drop_src_path_prefix_parts,omitempty"`
			Headers                    *[]string `tfsdk:"headers" json:"headers,omitempty"`
			Ip_filters                 *struct {
				Allow_list *[]string `tfsdk:"allow_list" json:"allow_list,omitempty"`
				Deny_list  *[]string `tfsdk:"deny_list" json:"deny_list,omitempty"`
			} `tfsdk:"ip_filters" json:"ip_filters,omitempty"`
			Load_balancing_policy *string   `tfsdk:"load_balancing_policy" json:"load_balancing_policy,omitempty"`
			Response_headers      *[]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
			Retry_status_codes    *[]string `tfsdk:"retry_status_codes" json:"retry_status_codes,omitempty"`
			Src_hosts             *[]string `tfsdk:"src_hosts" json:"src_hosts,omitempty"`
			Src_paths             *[]string `tfsdk:"src_paths" json:"src_paths,omitempty"`
			Url_prefix            *[]string `tfsdk:"url_prefix" json:"url_prefix,omitempty"`
		} `tfsdk:"unauthorized_access_config" json:"unauthorizedAccessConfig,omitempty"`
		UseStrictSecurity     *bool `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		UserNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"user_namespace_selector" json:"userNamespaceSelector,omitempty"`
		UserSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"user_selector" json:"userSelector,omitempty"`
		VolumeMounts *[]struct {
			MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmauthV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_auth_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmauthV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMAuth is the Schema for the vmauths API",
		MarkdownDescription: "VMAuth is the Schema for the vmauths API",
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
				Description:         "VMAuthSpec defines the desired state of VMAuth",
				MarkdownDescription: "VMAuthSpec defines the desired state of VMAuth",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.MapAttribute{
						Description:         "Affinity If specified, the pod's scheduling constraints.",
						MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_maps": schema.ListAttribute{
						Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the VMAuthobject, which shall be mounted into the VMAuth Pods.",
						MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the VMAuthobject, which shall be mounted into the VMAuth Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_secret": schema.StringAttribute{
						Description:         "ConfigSecret is the name of a Kubernetes Secret in the same namespace as theVMAuth object, which contains auth configuration for vmauth,configuration must be inside secret key: config.yaml.It must be created and managed manually.If it's defined, configuration for vmauth becomes unmanaged and operator'll not create any related secrets/config-reloaders",
						MarkdownDescription: "ConfigSecret is the name of a Kubernetes Secret in the same namespace as theVMAuth object, which contains auth configuration for vmauth,configuration must be inside secret key: config.yaml.It must be created and managed manually.If it's defined, configuration for vmauth becomes unmanaged and operator'll not create any related secrets/config-reloaders",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"containers": schema.ListAttribute{
						Description:         "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
						MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers.It can be useful for proxies, backup, etc.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_config": schema.SingleNestedAttribute{
						Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
						MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
						Attributes: map[string]schema.Attribute{
							"nameservers": schema.ListAttribute{
								Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
								MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"options": schema.ListNestedAttribute{
								Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
								MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Required.",
											MarkdownDescription: "Required.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
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

							"searches": schema.ListAttribute{
								Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
								MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

					"dns_policy": schema.StringAttribute{
						Description:         "DNSPolicy sets DNS policy for the pod",
						MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_args": schema.MapAttribute{
						Description:         "ExtraArgs that will be passed to  VMAuth podfor example remoteWrite.tmpDataPath: /tmp",
						MarkdownDescription: "ExtraArgs that will be passed to  VMAuth podfor example remoteWrite.tmpDataPath: /tmp",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_envs": schema.ListAttribute{
						Description:         "ExtraEnvs that will be added to VMAuth pod",
						MarkdownDescription: "ExtraEnvs that will be added to VMAuth pod",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "HostAliases provides mapping for ip and hostname,that would be propagated to pod,cannot be used with HostNetwork.",
						MarkdownDescription: "HostAliases provides mapping for ip and hostname,that would be propagated to pod,cannot be used with HostNetwork.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostnames": schema.ListAttribute{
									Description:         "Hostnames for the above IP address.",
									MarkdownDescription: "Hostnames for the above IP address.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ip": schema.StringAttribute{
									Description:         "IP address of the host file entry.",
									MarkdownDescription: "IP address of the host file entry.",
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

					"host_network": schema.BoolAttribute{
						Description:         "HostNetwork controls whether the pod may use the node network namespace",
						MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image - docker image settings for VMAuthif no specified operator uses default config version",
						MarkdownDescription: "Image - docker image settings for VMAuthif no specified operator uses default config version",
						Attributes: map[string]schema.Attribute{
							"pull_policy": schema.StringAttribute{
								Description:         "PullPolicy describes how to pull docker image",
								MarkdownDescription: "PullPolicy describes how to pull docker image",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "Repository contains name of docker image + it's repository if needed",
								MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag contains desired docker image version",
								MarkdownDescription: "Tag contains desired docker image version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespaceto use for pulling images from registriessee https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress enables ingress configuration for VMAuth.",
						MarkdownDescription: "Ingress enables ingress configuration for VMAuth.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"class_name": schema.StringAttribute{
								Description:         "ClassName defines ingress class name for VMAuth",
								MarkdownDescription: "ClassName defines ingress class name for VMAuth",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_rules": schema.ListNestedAttribute{
								Description:         "ExtraRules - additional rules for ingress,must be checked for correctness by user.",
								MarkdownDescription: "ExtraRules - additional rules for ingress,must be checked for correctness by user.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "host is the fully qualified domain name of a network host, as defined by RFC 3986.Note the following deviations from the 'host' part of theURI as defined in RFC 3986:1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future.Incoming requests are matched against the host before theIngressRuleValue. If the host is unspecified, the Ingress routes alltraffic based on the specified IngressRuleValue.host can be 'precise' which is a domain name without the terminating dot ofa network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain nameprefixed with a single wildcard label (e.g. '*.foo.com').The wildcard character '*' must appear by itself as the first DNS label andmatches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*').Requests will be matched against the Host field in the following way:1. If host is precise, the request matches this rule if the http host header is equal to Host.2. If host is a wildcard, then the request matches this rule if the http host headeris to equal to the suffix (removing the first label) of the wildcard rule.",
											MarkdownDescription: "host is the fully qualified domain name of a network host, as defined by RFC 3986.Note the following deviations from the 'host' part of theURI as defined in RFC 3986:1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future.Incoming requests are matched against the host before theIngressRuleValue. If the host is unspecified, the Ingress routes alltraffic based on the specified IngressRuleValue.host can be 'precise' which is a domain name without the terminating dot ofa network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain nameprefixed with a single wildcard label (e.g. '*.foo.com').The wildcard character '*' must appear by itself as the first DNS label andmatches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*').Requests will be matched against the Host field in the following way:1. If host is precise, the request matches this rule if the http host header is equal to Host.2. If host is a wildcard, then the request matches this rule if the http host headeris to equal to the suffix (removing the first label) of the wildcard rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "HTTPIngressRuleValue is a list of http selectors pointing to backends.In the example: http://<host>/<path>?<searchpart> -> backend wherewhere parts of the url correspond to RFC 3986, this resource will be usedto match against everything after the last '/' and before the first '?'or '#'.",
											MarkdownDescription: "HTTPIngressRuleValue is a list of http selectors pointing to backends.In the example: http://<host>/<path>?<searchpart> -> backend wherewhere parts of the url correspond to RFC 3986, this resource will be usedto match against everything after the last '/' and before the first '?'or '#'.",
											Attributes: map[string]schema.Attribute{
												"paths": schema.ListNestedAttribute{
													Description:         "paths is a collection of paths that map requests to backends.",
													MarkdownDescription: "paths is a collection of paths that map requests to backends.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"backend": schema.SingleNestedAttribute{
																Description:         "backend defines the referenced service endpoint to which the trafficwill be forwarded to.",
																MarkdownDescription: "backend defines the referenced service endpoint to which the trafficwill be forwarded to.",
																Attributes: map[string]schema.Attribute{
																	"resource": schema.SingleNestedAttribute{
																		Description:         "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
																		MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
																		Attributes: map[string]schema.Attribute{
																			"api_group": schema.StringAttribute{
																				Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind is the type of resource being referenced",
																				MarkdownDescription: "Kind is the type of resource being referenced",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name is the name of resource being referenced",
																				MarkdownDescription: "Name is the name of resource being referenced",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service": schema.SingleNestedAttribute{
																		Description:         "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
																		MarkdownDescription: "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
																				MarkdownDescription: "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"port": schema.SingleNestedAttribute{
																				Description:         "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
																				MarkdownDescription: "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
																						MarkdownDescription: "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"number": schema.Int64Attribute{
																						Description:         "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
																						MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
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
																Required: true,
																Optional: false,
																Computed: false,
															},

															"path": schema.StringAttribute{
																Description:         "path is matched against the path of an incoming request. Currently it cancontain characters disallowed from the conventional 'path' part of a URLas defined by RFC 3986. Paths must begin with a '/' and must be presentwhen using PathType with value 'Exact' or 'Prefix'.",
																MarkdownDescription: "path is matched against the path of an incoming request. Currently it cancontain characters disallowed from the conventional 'path' part of a URLas defined by RFC 3986. Paths must begin with a '/' and must be presentwhen using PathType with value 'Exact' or 'Prefix'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path_type": schema.StringAttribute{
																Description:         "pathType determines the interpretation of the path matching. PathType canbe one of the following values:* Exact: Matches the URL path exactly.* Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",
																MarkdownDescription: "pathType determines the interpretation of the path matching. PathType canbe one of the following values:* Exact: Matches the URL path exactly.* Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",
																Required:            true,
																Optional:            false,
																Computed:            false,
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"extra_tls": schema.ListNestedAttribute{
								Description:         "ExtraTLS - additional TLS configuration for ingressmust be checked for correctness by user.",
								MarkdownDescription: "ExtraTLS - additional TLS configuration for ingressmust be checked for correctness by user.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "hosts is a list of hosts included in the TLS certificate. The values inthis list must match the name/s used in the tlsSecret. Defaults to thewildcard host setting for the loadbalancer controller fulfilling thisIngress, if left unspecified.",
											MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values inthis list must match the name/s used in the tlsSecret. Defaults to thewildcard host setting for the loadbalancer controller fulfilling thisIngress, if left unspecified.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "secretName is the name of the secret used to terminate TLS traffic onport 443. Field is left optional to allow TLS routing based on SNIhostname alone. If the SNI host in a listener conflicts with the 'Host'header field used by an IngressRule, the SNI host is used for terminationand value of the 'Host' header is used for routing.",
											MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic onport 443. Field is left optional to allow TLS routing based on SNIhostname alone. If the SNI host in a listener conflicts with the 'Host'header field used by an IngressRule, the SNI host is used for terminationand value of the 'Host' header is used for routing.",
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

							"host": schema.StringAttribute{
								Description:         "Host defines ingress host parameter for default ruleIt will be used, only if TlsHosts is empty",
								MarkdownDescription: "Host defines ingress host parameter for default ruleIt will be used, only if TlsHosts is empty",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_hosts": schema.ListAttribute{
								Description:         "TlsHosts configures TLS access for ingress, tlsSecretName must be defined for it.",
								MarkdownDescription: "TlsHosts configures TLS access for ingress, tlsSecretName must be defined for it.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret_name": schema.StringAttribute{
								Description:         "TlsSecretName defines secretname at the VMAuth namespace with cert and keyhttps://kubernetes.io/docs/concepts/services-networking/ingress/#tls",
								MarkdownDescription: "TlsSecretName defines secretname at the VMAuth namespace with cert and keyhttps://kubernetes.io/docs/concepts/services-networking/ingress/#tls",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"init_containers": schema.ListAttribute{
						Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmSingle configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmSingle configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"license": schema.SingleNestedAttribute{
						Description:         "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						MarkdownDescription: "License allows to configure license key to be used for enterprise features.Using license key is supported starting from VictoriaMetrics v1.94.0.See: https://docs.victoriametrics.com/enterprise.html",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								MarkdownDescription: "Enterprise license key. This flag is available only in VictoriaMetrics enterprise.Documentation - https://docs.victoriametrics.com/enterprise.htmlfor more information, visit https://victoriametrics.com/products/enterprise/ .To request a trial license, go to https://victoriametrics.com/products/enterprise/trial/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_ref": schema.SingleNestedAttribute{
								Description:         "KeyRef is reference to secret with license key for enterprise features.",
								MarkdownDescription: "KeyRef is reference to secret with license key for enterprise features.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"liveness_probe": schema.MapAttribute{
						Description:         "LivenessProbe that will be added CRD pod",
						MarkdownDescription: "LivenessProbe that will be added CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_format": schema.StringAttribute{
						Description:         "LogFormat for VMAuth to be configured with.",
						MarkdownDescription: "LogFormat for VMAuth to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("default", "json"),
						},
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel for victoria metrics single to be configured with.",
						MarkdownDescription: "LogLevel for victoria metrics single to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
						},
					},

					"min_ready_seconds": schema.Int64Attribute{
						Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
						MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
						MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_disruption_budget": schema.SingleNestedAttribute{
						Description:         "PodDisruptionBudget created by operator",
						MarkdownDescription: "PodDisruptionBudget created by operator",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.StringAttribute{
								Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by'selector' are unavailable after the eviction, i.e. even in absence ofthe evicted pod. For example, one can prevent all voluntary evictionsby specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_available": schema.StringAttribute{
								Description:         "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
								MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by'selector' will still be available after the eviction, i.e. even in theabsence of the evicted pod.  So for example you can prevent all voluntaryevictions by specifying '100%'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector_labels": schema.MapAttribute{
								Description:         "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
								MarkdownDescription: "replaces default labels selector generated by operatorit's useful when you need to create custom budget",
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

					"pod_metadata": schema.SingleNestedAttribute{
						Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMAuth pods.",
						MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMAuth pods.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_security_policy_name": schema.StringAttribute{
						Description:         "PodSecurityPolicyName - defines name for podSecurityPolicyin case of empty value, prefixedName will be used.",
						MarkdownDescription: "PodSecurityPolicyName - defines name for podSecurityPolicyin case of empty value, prefixedName will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.StringAttribute{
						Description:         "Port listen port",
						MarkdownDescription: "Port listen port",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName assigned to the Pods",
						MarkdownDescription: "PriorityClassName assigned to the Pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_gates": schema.ListNestedAttribute{
						Description:         "ReadinessGates defines pod readiness gates",
						MarkdownDescription: "ReadinessGates defines pod readiness gates",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"condition_type": schema.StringAttribute{
									Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
									MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

					"readiness_probe": schema.MapAttribute{
						Description:         "ReadinessProbe that will be added CRD pod",
						MarkdownDescription: "ReadinessProbe that will be added CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_count": schema.Int64Attribute{
						Description:         "ReplicaCount is the expected size of the VMAuth",
						MarkdownDescription: "ReplicaCount is the expected size of the VMAuth",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
						MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not defined default resources from operator config will be used",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
								Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"revision_history_limit_count": schema.Int64Attribute{
						Description:         "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
						MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment ormaximum number of revisions that will be maintained in the StatefulSet's revision history.Defaults to 10.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"runtime_class_name": schema.StringAttribute{
						Description:         "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
						MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod.https://kubernetes.io/docs/concepts/containers/runtime-class/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scheduler_name": schema.StringAttribute{
						Description:         "SchedulerName - defines kubernetes scheduler name",
						MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secrets": schema.ListAttribute{
						Description:         "Secrets is a list of Secrets in the same namespace as the VMAuthobject, which shall be mounted into the VMAuth Pods.",
						MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the VMAuthobject, which shall be mounted into the VMAuth Pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_context": schema.MapAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.This defaults to the default PodSecurityContext.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"select_all_by_default": schema.BoolAttribute{
						Description:         "SelectAllByDefault changes default behavior for empty CRD selectors, such userSelector.with selectAllByDefault: true and empty userSelector and userNamespaceSelectorOperator selects all exist userswith selectAllByDefault: false - selects nothing",
						MarkdownDescription: "SelectAllByDefault changes default behavior for empty CRD selectors, such userSelector.with selectAllByDefault: true and empty userSelector and userNamespaceSelectorOperator selects all exist userswith selectAllByDefault: false - selects nothing",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run theVMAuth Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run theVMAuth Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_scrape_spec": schema.MapAttribute{
						Description:         "ServiceScrapeSpec that will be added to vmauth VMServiceScrape spec",
						MarkdownDescription: "ServiceScrapeSpec that will be added to vmauth VMServiceScrape spec",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_spec": schema.SingleNestedAttribute{
						Description:         "ServiceSpec that will be added to vmsingle service spec",
						MarkdownDescription: "ServiceSpec that will be added to vmsingle service spec",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
								MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, althoughsome resources may allow a client to request the generation of an appropriate nameautomatically. Name is primarily intended for creation idempotence and configurationdefinition.Cannot be updated.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": schema.MapAttribute{
								Description:         "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
								MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
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

					"startup_probe": schema.MapAttribute{
						Description:         "StartupProbe that will be added to CRD pod",
						MarkdownDescription: "StartupProbe that will be added to CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "TerminationGracePeriodSeconds period for container graceful termination",
						MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations If specified, the pod's tolerations.",
						MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"topology_spread_constraints": schema.ListAttribute{
						Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option,controls how pods are spread across your cluster among failure-domainssuch as regions, zones, nodes, and other user-defined topology domainshttps://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"unauthorized_access_config": schema.ListNestedAttribute{
						Description:         "UnauthorizedAccessConfig configures access for un authorized users",
						MarkdownDescription: "UnauthorizedAccessConfig configures access for un authorized users",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"drop_src_path_prefix_parts": schema.Int64Attribute{
									Description:         "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
									MarkdownDescription: "DropSrcPathPrefixParts is the number of '/'-delimited request path prefix parts to drop before proxying the request to backend.See https://docs.victoriametrics.com/vmauth.html#dropping-request-path-prefix for more details.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"headers": schema.ListAttribute{
									Description:         "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
									MarkdownDescription: "Headers represent additional http headers, that vmauth usesin form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.68.0 version of vmauth",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ip_filters": schema.SingleNestedAttribute{
									Description:         "IPFilters defines filter for src ip addressenterprise only",
									MarkdownDescription: "IPFilters defines filter for src ip addressenterprise only",
									Attributes: map[string]schema.Attribute{
										"allow_list": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"deny_list": schema.ListAttribute{
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

								"load_balancing_policy": schema.StringAttribute{
									Description:         "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
									MarkdownDescription: "LoadBalancingPolicy defines load balancing policy to use for backend urls.Supported policies: least_loaded, first_available.See https://docs.victoriametrics.com/vmauth.html#load-balancing for more details (default 'least_loaded')",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("least_loaded", "first_available"),
									},
								},

								"response_headers": schema.ListAttribute{
									Description:         "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
									MarkdownDescription: "ResponseHeaders represent additional http headers, that vmauth adds for request responsein form of ['header_key: header_value']multiple values for header key:['header_key: value1,value2']it's available since 1.93.0 version of vmauth",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"retry_status_codes": schema.ListAttribute{
									Description:         "RetryStatusCodes defines http status codes in numeric format for request retriese.g. [429,503]",
									MarkdownDescription: "RetryStatusCodes defines http status codes in numeric format for request retriese.g. [429,503]",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"src_hosts": schema.ListAttribute{
									Description:         "SrcHosts is the list of regular expressions, which match the request hostname.",
									MarkdownDescription: "SrcHosts is the list of regular expressions, which match the request hostname.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"src_paths": schema.ListAttribute{
									Description:         "Paths src request paths",
									MarkdownDescription: "Paths src request paths",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url_prefix": schema.ListAttribute{
									Description:         "URLs defines url_prefix for dst routing",
									MarkdownDescription: "URLs defines url_prefix for dst routing",
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

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_namespace_selector": schema.SingleNestedAttribute{
						Description:         "UserNamespaceSelector Namespaces to be selected for  VMAuth discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAuth namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "UserNamespaceSelector Namespaces to be selected for  VMAuth discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAuth namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"user_selector": schema.SingleNestedAttribute{
						Description:         "UserSelector defines VMUser to be selected for config file generation.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAuth namespace.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "UserSelector defines VMUser to be selected for config file generation.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAuth namespace.If both nil - behaviour controlled by selectAllByDefault",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMAuth container,that are generated as a result of StorageSpec objects.",
						MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment definition.VolumeMounts specified will be appended to other VolumeMounts in the VMAuth container,that are generated as a result of StorageSpec objects.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
									MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"mount_propagation": schema.StringAttribute{
									Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
									MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
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
									Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
									MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
									Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
									MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path_expr": schema.StringAttribute{
									Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
									MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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

					"volumes": schema.ListAttribute{
						Description:         "Volumes allows configuration of additional volumes on the output deploy definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
						MarkdownDescription: "Volumes allows configuration of additional volumes on the output deploy definition.Volumes specified will be appended to other volumes that are generated as a result ofStorageSpec objects.",
						ElementType:         types.MapType{ElemType: types.StringType},
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

func (r *OperatorVictoriametricsComVmauthV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_auth_v1beta1_manifest")

	var model OperatorVictoriametricsComVmauthV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMAuth")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
