/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package instana_io_v1

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
	_ datasource.DataSource = &InstanaIoInstanaAgentV1Manifest{}
)

func NewInstanaIoInstanaAgentV1Manifest() datasource.DataSource {
	return &InstanaIoInstanaAgentV1Manifest{}
}

type InstanaIoInstanaAgentV1Manifest struct{}

type InstanaIoInstanaAgentV1ManifestData struct {
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
		Agent *struct {
			AdditionalBackends *[]struct {
				EndpointHost *string `tfsdk:"endpoint_host" json:"endpointHost,omitempty"`
				EndpointPort *string `tfsdk:"endpoint_port" json:"endpointPort,omitempty"`
				Key          *string `tfsdk:"key" json:"key,omitempty"`
			} `tfsdk:"additional_backends" json:"additionalBackends,omitempty"`
			Charts_url    *string `tfsdk:"charts_url" json:"charts_url,omitempty"`
			Configuration *struct {
				AutoMountConfigEntries *bool `tfsdk:"auto_mount_config_entries" json:"autoMountConfigEntries,omitempty"`
			} `tfsdk:"configuration" json:"configuration,omitempty"`
			Configuration_yaml *string            `tfsdk:"configuration_yaml" json:"configuration_yaml,omitempty"`
			DownloadKey        *string            `tfsdk:"download_key" json:"downloadKey,omitempty"`
			EndpointHost       *string            `tfsdk:"endpoint_host" json:"endpointHost,omitempty"`
			EndpointPort       *string            `tfsdk:"endpoint_port" json:"endpointPort,omitempty"`
			Env                *map[string]string `tfsdk:"env" json:"env,omitempty"`
			Host               *struct {
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			} `tfsdk:"host" json:"host,omitempty"`
			Image *struct {
				Digest      *string `tfsdk:"digest" json:"digest,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				PullPolicy  *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				PullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
				Tag *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			InstanaMvnRepoUrl *string `tfsdk:"instana_mvn_repo_url" json:"instanaMvnRepoUrl,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			KeysSecret        *string `tfsdk:"keys_secret" json:"keysSecret,omitempty"`
			ListenAddress     *string `tfsdk:"listen_address" json:"listenAddress,omitempty"`
			MinReadySeconds   *int64  `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
			Pod               *struct {
				Affinity *struct {
					NodeAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							Preference *struct {
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
							} `tfsdk:"preference" json:"preference,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *struct {
							NodeSelectorTerms *[]struct {
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
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					PodAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Limits            *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
				Requests          *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				Tolerations       *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"pod" json:"pod,omitempty"`
			ProxyHost               *string `tfsdk:"proxy_host" json:"proxyHost,omitempty"`
			ProxyPassword           *string `tfsdk:"proxy_password" json:"proxyPassword,omitempty"`
			ProxyPort               *string `tfsdk:"proxy_port" json:"proxyPort,omitempty"`
			ProxyProtocol           *string `tfsdk:"proxy_protocol" json:"proxyProtocol,omitempty"`
			ProxyUseDNS             *bool   `tfsdk:"proxy_use_dns" json:"proxyUseDNS,omitempty"`
			ProxyUser               *string `tfsdk:"proxy_user" json:"proxyUser,omitempty"`
			RedactKubernetesSecrets *string `tfsdk:"redact_kubernetes_secrets" json:"redactKubernetesSecrets,omitempty"`
			Tls                     *struct {
				Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			UpdateStrategy *struct {
				RollingUpdate *struct {
					MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
					MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		} `tfsdk:"agent" json:"agent,omitempty"`
		Agent_clusterRoleBindingName *string            `tfsdk:"agent_cluster_role_binding_name" json:"agent.clusterRoleBindingName,omitempty"`
		Agent_clusterRoleName        *string            `tfsdk:"agent_cluster_role_name" json:"agent.clusterRoleName,omitempty"`
		Agent_configMapName          *string            `tfsdk:"agent_config_map_name" json:"agent.configMapName,omitempty"`
		Agent_cpuLimit               *string            `tfsdk:"agent_cpu_limit" json:"agent.cpuLimit,omitempty"`
		Agent_cpuReq                 *string            `tfsdk:"agent_cpu_req" json:"agent.cpuReq,omitempty"`
		Agent_daemonSetName          *string            `tfsdk:"agent_daemon_set_name" json:"agent.daemonSetName,omitempty"`
		Agent_downloadKey            *string            `tfsdk:"agent_download_key" json:"agent.downloadKey,omitempty"`
		Agent_endpoint_host          *string            `tfsdk:"agent_endpoint_host" json:"agent.endpoint.host,omitempty"`
		Agent_endpoint_port          *int64             `tfsdk:"agent_endpoint_port" json:"agent.endpoint.port,omitempty"`
		Agent_env                    *map[string]string `tfsdk:"agent_env" json:"agent.env,omitempty"`
		Agent_host_repository        *string            `tfsdk:"agent_host_repository" json:"agent.host.repository,omitempty"`
		Agent_image                  *string            `tfsdk:"agent_image" json:"agent.image,omitempty"`
		Agent_imagePullPolicy        *string            `tfsdk:"agent_image_pull_policy" json:"agent.imagePullPolicy,omitempty"`
		Agent_key                    *string            `tfsdk:"agent_key" json:"agent.key,omitempty"`
		Agent_memLimit               *string            `tfsdk:"agent_mem_limit" json:"agent.memLimit,omitempty"`
		Agent_memReq                 *string            `tfsdk:"agent_mem_req" json:"agent.memReq,omitempty"`
		Agent_rbac_create            *bool              `tfsdk:"agent_rbac_create" json:"agent.rbac.create,omitempty"`
		Agent_secretName             *string            `tfsdk:"agent_secret_name" json:"agent.secretName,omitempty"`
		Agent_serviceAccountName     *string            `tfsdk:"agent_service_account_name" json:"agent.serviceAccountName,omitempty"`
		Agent_tls_certificate        *string            `tfsdk:"agent_tls_certificate" json:"agent.tls.certificate,omitempty"`
		Agent_tls_key                *string            `tfsdk:"agent_tls_key" json:"agent.tls.key,omitempty"`
		Agent_tls_secretName         *string            `tfsdk:"agent_tls_secret_name" json:"agent.tls.secretName,omitempty"`
		Agent_zone_name              *string            `tfsdk:"agent_zone_name" json:"agent.zone.name,omitempty"`
		Cluster                      *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster" json:"cluster,omitempty"`
		Cluster_name *string            `tfsdk:"cluster_name" json:"cluster.name,omitempty"`
		Config_files *map[string]string `tfsdk:"config_files" json:"config.files,omitempty"`
		K8s_sensor   *struct {
			Deployment *struct {
				Enabled         *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
				MinReadySeconds *int64 `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				Pod             *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"pod" json:"pod,omitempty"`
				Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
			Image *struct {
				Digest      *string `tfsdk:"digest" json:"digest,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				PullPolicy  *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				PullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
				Tag *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
		} `tfsdk:"k8s_sensor" json:"k8s_sensor,omitempty"`
		Kubernetes *struct {
			Deployment *struct {
				Enabled         *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
				MinReadySeconds *int64 `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				Pod             *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"pod" json:"pod,omitempty"`
				Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Openshift     *bool `tfsdk:"openshift" json:"openshift,omitempty"`
		Opentelemetry *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Grpc    *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"grpc" json:"grpc,omitempty"`
			Http *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
		} `tfsdk:"opentelemetry" json:"opentelemetry,omitempty"`
		Opentelemetry_enabled *bool   `tfsdk:"opentelemetry_enabled" json:"opentelemetry.enabled,omitempty"`
		PinnedChartVersion    *string `tfsdk:"pinned_chart_version" json:"pinnedChartVersion,omitempty"`
		PodSecurityPolicy     *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pod_security_policy" json:"podSecurityPolicy,omitempty"`
		Prometheus *struct {
			RemoteWrite *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"remote_write" json:"remoteWrite,omitempty"`
		} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		Rbac *struct {
			Create *bool `tfsdk:"create" json:"create,omitempty"`
		} `tfsdk:"rbac" json:"rbac,omitempty"`
		Service *struct {
			Create *bool `tfsdk:"create" json:"create,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		ServiceAccount *struct {
			Create *bool   `tfsdk:"create" json:"create,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		Zone *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InstanaIoInstanaAgentV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_instana_io_instana_agent_v1_manifest"
}

func (r *InstanaIoInstanaAgentV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "InstanaAgent is the Schema for the agents API",
		MarkdownDescription: "InstanaAgent is the Schema for the agents API",
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
				Description:         "InstanaAgentSpec defines the desired state of the Instana Agent",
				MarkdownDescription: "InstanaAgentSpec defines the desired state of the Instana Agent",
				Attributes: map[string]schema.Attribute{
					"agent": schema.SingleNestedAttribute{
						Description:         "Agent deployment specific fields.",
						MarkdownDescription: "Agent deployment specific fields.",
						Attributes: map[string]schema.Attribute{
							"additional_backends": schema.ListNestedAttribute{
								Description:         "These are additional backends the Instana agent will report to besides the one configured via the 'agent.endpointHost', 'agent.endpointPort' and 'agent.key' setting.",
								MarkdownDescription: "These are additional backends the Instana agent will report to besides the one configured via the 'agent.endpointHost', 'agent.endpointPort' and 'agent.key' setting.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint_host": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"endpoint_port": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"charts_url": schema.StringAttribute{
								Description:         "Custom agent charts url.",
								MarkdownDescription: "Custom agent charts url.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"configuration": schema.SingleNestedAttribute{
								Description:         "Mount in a ConfigMap with Agent configuration. Alternative to the 'configuration_yaml' field.",
								MarkdownDescription: "Mount in a ConfigMap with Agent configuration. Alternative to the 'configuration_yaml' field.",
								Attributes: map[string]schema.Attribute{
									"auto_mount_config_entries": schema.BoolAttribute{
										Description:         "When setting this to true, the Helm chart will automatically look up the entries of the default instana-agent ConfigMap, and mount as agent configuration files under /opt/instana/agent/etc/instana all entries with keys that match the 'configuration-*.yaml' scheme",
										MarkdownDescription: "When setting this to true, the Helm chart will automatically look up the entries of the default instana-agent ConfigMap, and mount as agent configuration files under /opt/instana/agent/etc/instana all entries with keys that match the 'configuration-*.yaml' scheme",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"configuration_yaml": schema.StringAttribute{
								Description:         "Supply Agent configuration e.g. for configuring certain Sensors.",
								MarkdownDescription: "Supply Agent configuration e.g. for configuring certain Sensors.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"download_key": schema.StringAttribute{
								Description:         "The DownloadKey, sometimes known as 'sales key', that allows you to download software from Instana. It might be needed to specify this in addition to the Key.",
								MarkdownDescription: "The DownloadKey, sometimes known as 'sales key', that allows you to download software from Instana. It might be needed to specify this in addition to the Key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoint_host": schema.StringAttribute{
								Description:         "EndpointHost is the hostname of the Instana server your agents will connect to.",
								MarkdownDescription: "EndpointHost is the hostname of the Instana server your agents will connect to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"endpoint_port": schema.StringAttribute{
								Description:         "EndpointPort is the port number (as a String) of the Instana server your agents will connect to.",
								MarkdownDescription: "EndpointPort is the port number (as a String) of the Instana server your agents will connect to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"env": schema.MapAttribute{
								Description:         "Use the 'env' field to set additional environment variables for the Instana Agent, for example: env: INSTANA_AGENT_TAGS: dev",
								MarkdownDescription: "Use the 'env' field to set additional environment variables for the Instana Agent, for example: env: INSTANA_AGENT_TAGS: dev",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.SingleNestedAttribute{
								Description:         "Host sets a host path to be mounted as the Agent Maven repository (mainly for debugging or development purposes)",
								MarkdownDescription: "Host sets a host path to be mounted as the Agent Maven repository (mainly for debugging or development purposes)",
								Attributes: map[string]schema.Attribute{
									"repository": schema.StringAttribute{
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

							"image": schema.SingleNestedAttribute{
								Description:         "Override the container image used for the Instana Agent pods.",
								MarkdownDescription: "Override the container image used for the Instana Agent pods.",
								Attributes: map[string]schema.Attribute{
									"digest": schema.StringAttribute{
										Description:         "Digest (a.k.a. Image ID) of the agent container image. If specified, it has priority over 'agent.image.tag', which will then be ignored.",
										MarkdownDescription: "Digest (a.k.a. Image ID) of the agent container image. If specified, it has priority over 'agent.image.tag', which will then be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the container image of the Instana agent.",
										MarkdownDescription: "Name is the name of the container image of the Instana agent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy specifies when to pull the image container.",
										MarkdownDescription: "PullPolicy specifies when to pull the image container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_secrets": schema.ListNestedAttribute{
										Description:         "PullSecrets allows you to override the default pull secret that is created when 'agent.image.name' starts with 'containers.instana.io'. Setting 'agent.image.pullSecrets' prevents the creation of the default 'containers-instana-io' secret.",
										MarkdownDescription: "PullSecrets allows you to override the default pull secret that is created when 'agent.image.name' starts with 'containers.instana.io'. Setting 'agent.image.pullSecrets' prevents the creation of the default 'containers-instana-io' secret.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
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

									"tag": schema.StringAttribute{
										Description:         "Tag is the name of the agent container image; if 'agent.image.digest' is specified, this property is ignored.",
										MarkdownDescription: "Tag is the name of the agent container image; if 'agent.image.digest' is specified, this property is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"instana_mvn_repo_url": schema.StringAttribute{
								Description:         "Override for the Maven repository URL when the Agent needs to connect to a locally provided Maven repository 'proxy' Alternative to 'Host' for referencing a different Maven repo.",
								MarkdownDescription: "Override for the Maven repository URL when the Agent needs to connect to a locally provided Maven repository 'proxy' Alternative to 'Host' for referencing a different Maven repo.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "Key is the secret token which your agent uses to authenticate to Instana's servers.",
								MarkdownDescription: "Key is the secret token which your agent uses to authenticate to Instana's servers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keys_secret": schema.StringAttribute{
								Description:         "Rather than specifying the Key and optionally the DownloadKey, you can 'bring your own secret' creating it in the namespace in which you install the 'instana-agent' and specify its name in the 'KeysSecret' field. The secret you create must contain a field called 'key' and optionally one called 'downloadKey', which contain, respectively, the values you'd otherwise set in '.agent.key' and 'agent.downloadKey'.",
								MarkdownDescription: "Rather than specifying the Key and optionally the DownloadKey, you can 'bring your own secret' creating it in the namespace in which you install the 'instana-agent' and specify its name in the 'KeysSecret' field. The secret you create must contain a field called 'key' and optionally one called 'downloadKey', which contain, respectively, the values you'd otherwise set in '.agent.key' and 'agent.downloadKey'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"listen_address": schema.StringAttribute{
								Description:         "ListenAddress is the IP addresses the Agent HTTP server will listen on. Normally this will just be localhost ('127.0.0.1'), the pod public IP and any container runtime bridge interfaces. Set 'listenAddress: *' for making the Agent listen on all network interfaces.",
								MarkdownDescription: "ListenAddress is the IP addresses the Agent HTTP server will listen on. Normally this will just be localhost ('127.0.0.1'), the pod public IP and any container runtime bridge interfaces. Set 'listenAddress: *' for making the Agent listen on all network interfaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
								MarkdownDescription: "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "Set agent mode, possible options are APM, INFRASTRUCTURE or AWS. KUBERNETES should not be used but instead enabled via 'kubernetes.deployment.enabled: true'.",
								MarkdownDescription: "Set agent mode, possible options are APM, INFRASTRUCTURE or AWS. KUBERNETES should not be used but instead enabled via 'kubernetes.deployment.enabled: true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod": schema.SingleNestedAttribute{
								Description:         "Override Agent Pod specific settings such as annotations, labels and resources.",
								MarkdownDescription: "Override Agent Pod specific settings such as annotations, labels and resources.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "agent.pod.affinity are affinities to influence agent pod assignment. https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/",
										MarkdownDescription: "agent.pod.affinity are affinities to influence agent pod assignment. https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.SingleNestedAttribute{
												Description:         "Describes node affinity scheduling rules for the pod.",
												MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"preference": schema.SingleNestedAttribute{
																	Description:         "A node selector term, associated with the corresponding weight.",
																	MarkdownDescription: "A node selector term, associated with the corresponding weight.",
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
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																	MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

													"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
														Attributes: map[string]schema.Attribute{
															"node_selector_terms": schema.ListNestedAttribute{
																Description:         "Required. A list of node selector terms. The terms are ORed.",
																MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_affinity": schema.SingleNestedAttribute{
												Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
												MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																	MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																								Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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
																								Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																	MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

													"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

											"pod_anti_affinity": schema.SingleNestedAttribute{
												Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
												MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																	MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																								Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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
																								Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																	MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

													"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces. This field is beta-level and is only honored when PodAffinityNamespaceSelector feature is enabled.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"annotations": schema.MapAttribute{
										Description:         "agent.pod.annotations are additional annotations to be added to the agent pods.",
										MarkdownDescription: "agent.pod.annotations are additional annotations to be added to the agent pods.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "agent.pod.labels are additional labels to be added to the agent pods.",
										MarkdownDescription: "agent.pod.labels are additional labels to be added to the agent pods.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority_class_name": schema.StringAttribute{
										Description:         "agent.pod.priorityClassName is the name of an existing PriorityClass that should be set on the agent pods https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/",
										MarkdownDescription: "agent.pod.priorityClassName is the name of an existing PriorityClass that should be set on the agent pods https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "agent.pod.tolerations are tolerations to influence agent pod assignment.",
										MarkdownDescription: "agent.pod.tolerations are tolerations to influence agent pod assignment.",
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

							"proxy_host": schema.StringAttribute{
								Description:         "proxyHost sets the INSTANA_AGENT_PROXY_HOST environment variable.",
								MarkdownDescription: "proxyHost sets the INSTANA_AGENT_PROXY_HOST environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_password": schema.StringAttribute{
								Description:         "proxyPassword sets the INSTANA_AGENT_PROXY_PASSWORD environment variable.",
								MarkdownDescription: "proxyPassword sets the INSTANA_AGENT_PROXY_PASSWORD environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_port": schema.StringAttribute{
								Description:         "proxyPort sets the INSTANA_AGENT_PROXY_PORT environment variable.",
								MarkdownDescription: "proxyPort sets the INSTANA_AGENT_PROXY_PORT environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_protocol": schema.StringAttribute{
								Description:         "proxyProtocol sets the INSTANA_AGENT_PROXY_PROTOCOL environment variable.",
								MarkdownDescription: "proxyProtocol sets the INSTANA_AGENT_PROXY_PROTOCOL environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_use_dns": schema.BoolAttribute{
								Description:         "proxyUseDNS sets the INSTANA_AGENT_PROXY_USE_DNS environment variable.",
								MarkdownDescription: "proxyUseDNS sets the INSTANA_AGENT_PROXY_USE_DNS environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_user": schema.StringAttribute{
								Description:         "proxyUser sets the INSTANA_AGENT_PROXY_USER environment variable.",
								MarkdownDescription: "proxyUser sets the INSTANA_AGENT_PROXY_USER environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redact_kubernetes_secrets": schema.StringAttribute{
								Description:         "RedactKubernetesSecrets sets the INSTANA_KUBERNETES_REDACT_SECRETS environment variable.",
								MarkdownDescription: "RedactKubernetesSecrets sets the INSTANA_KUBERNETES_REDACT_SECRETS environment variable.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS for end-to-end encryption between the Instana Agent and clients accessing the Agent. The Instana Agent does not yet allow enforcing TLS encryption, enabling makes it possible for clients to 'opt-in'. So TLS is only enabled on a connection when requested by the client.",
								MarkdownDescription: "TLS for end-to-end encryption between the Instana Agent and clients accessing the Agent. The Instana Agent does not yet allow enforcing TLS encryption, enabling makes it possible for clients to 'opt-in'. So TLS is only enabled on a connection when requested by the client.",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.StringAttribute{
										Description:         "certificate (together with key) is the alternative to an existing Secret. Must be base64 encoded.",
										MarkdownDescription: "certificate (together with key) is the alternative to an existing Secret. Must be base64 encoded.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "key (together with certificate) is the alternative to an existing Secret. Must be base64 encoded.",
										MarkdownDescription: "key (together with certificate) is the alternative to an existing Secret. Must be base64 encoded.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "secretName is the name of the secret that has the relevant files.",
										MarkdownDescription: "secretName is the name of the secret that has the relevant files.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"update_strategy": schema.SingleNestedAttribute{
								Description:         "Control how to update the Agent DaemonSet",
								MarkdownDescription: "Control how to update the Agent DaemonSet",
								Attributes: map[string]schema.Attribute{
									"rolling_update": schema.SingleNestedAttribute{
										Description:         "Rolling update config params. Present only if type = 'RollingUpdate'. --- TODO: Update this to follow our convention for oneOf, whatever we decide it to be. Same as Deployment 'strategy.rollingUpdate'. See https://github.com/kubernetes/kubernetes/issues/35345",
										MarkdownDescription: "Rolling update config params. Present only if type = 'RollingUpdate'. --- TODO: Update this to follow our convention for oneOf, whatever we decide it to be. Same as Deployment 'strategy.rollingUpdate'. See https://github.com/kubernetes/kubernetes/issues/35345",
										Attributes: map[string]schema.Attribute{
											"max_surge": schema.StringAttribute{
												Description:         "The maximum number of nodes with an existing available DaemonSet pod that can have an updated DaemonSet pod during during an update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up to a minimum of 1. Default value is 0. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their a new pod created before the old pod is marked as deleted. The update starts by launching new pods on 30% of nodes. Once an updated pod is available (Ready for at least minReadySeconds) the old DaemonSet pod on that node is marked deleted. If the old pod becomes unavailable for any reason (Ready transitions to false, is evicted, or is drained) an updated pod is immediatedly created on that node without considering surge limits. Allowing surge implies the possibility that the resources consumed by the daemonset on any given node can double if the readiness check fails, and so resource intensive daemonsets should take into account that they may cause evictions during disruption. This is beta field and enabled/disabled by DaemonSetUpdateSurge feature gate.",
												MarkdownDescription: "The maximum number of nodes with an existing available DaemonSet pod that can have an updated DaemonSet pod during during an update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up to a minimum of 1. Default value is 0. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their a new pod created before the old pod is marked as deleted. The update starts by launching new pods on 30% of nodes. Once an updated pod is available (Ready for at least minReadySeconds) the old DaemonSet pod on that node is marked deleted. If the old pod becomes unavailable for any reason (Ready transitions to false, is evicted, or is drained) an updated pod is immediatedly created on that node without considering surge limits. Allowing surge implies the possibility that the resources consumed by the daemonset on any given node can double if the readiness check fails, and so resource intensive daemonsets should take into account that they may cause evictions during disruption. This is beta field and enabled/disabled by DaemonSetUpdateSurge feature gate.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.StringAttribute{
												Description:         "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0 if MaxSurge is 0 Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
												MarkdownDescription: "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0 if MaxSurge is 0 Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
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
										Description:         "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
										MarkdownDescription: "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"agent_cluster_role_binding_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_cluster_role_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_config_map_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_cpu_limit": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_cpu_req": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_daemon_set_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_download_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_endpoint_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_endpoint_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_env": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_host_repository": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_image_pull_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_mem_limit": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_mem_req": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_rbac_create": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_secret_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_service_account_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_tls_certificate": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_tls_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_tls_secret_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"agent_zone_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster": schema.SingleNestedAttribute{
						Description:         "Name of the cluster, that will be assigned to this cluster in Instana. Either specifying the 'cluster.name' or 'zone.name' is mandatory.",
						MarkdownDescription: "Name of the cluster, that will be assigned to this cluster in Instana. Either specifying the 'cluster.name' or 'zone.name' is mandatory.",
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

					"cluster_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_files": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"k8s_sensor": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"deployment": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_ready_seconds": schema.Int64Attribute{
										Description:         "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
										MarkdownDescription: "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "Override pod resource requirements for the Kubernetes Sensor pods.",
										MarkdownDescription: "Override pod resource requirements for the Kubernetes Sensor pods.",
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

									"replicas": schema.Int64Attribute{
										Description:         "Specify the number of replicas for the Kubernetes Sensor.",
										MarkdownDescription: "Specify the number of replicas for the Kubernetes Sensor.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"digest": schema.StringAttribute{
										Description:         "Digest (a.k.a. Image ID) of the agent container image. If specified, it has priority over 'agent.image.tag', which will then be ignored.",
										MarkdownDescription: "Digest (a.k.a. Image ID) of the agent container image. If specified, it has priority over 'agent.image.tag', which will then be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of the container image of the Instana agent.",
										MarkdownDescription: "Name is the name of the container image of the Instana agent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy specifies when to pull the image container.",
										MarkdownDescription: "PullPolicy specifies when to pull the image container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_secrets": schema.ListNestedAttribute{
										Description:         "PullSecrets allows you to override the default pull secret that is created when 'agent.image.name' starts with 'containers.instana.io'. Setting 'agent.image.pullSecrets' prevents the creation of the default 'containers-instana-io' secret.",
										MarkdownDescription: "PullSecrets allows you to override the default pull secret that is created when 'agent.image.name' starts with 'containers.instana.io'. Setting 'agent.image.pullSecrets' prevents the creation of the default 'containers-instana-io' secret.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
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

									"tag": schema.StringAttribute{
										Description:         "Tag is the name of the agent container image; if 'agent.image.digest' is specified, this property is ignored.",
										MarkdownDescription: "Tag is the name of the agent container image; if 'agent.image.digest' is specified, this property is ignored.",
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

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "Allows for installment of the Kubernetes Sensor as separate pod. Which allows for better tailored resource settings (mainly memory) both for the Agent pods and the Kubernetes Sensor pod.",
						MarkdownDescription: "Allows for installment of the Kubernetes Sensor as separate pod. Which allows for better tailored resource settings (mainly memory) both for the Agent pods and the Kubernetes Sensor pod.",
						Attributes: map[string]schema.Attribute{
							"deployment": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_ready_seconds": schema.Int64Attribute{
										Description:         "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
										MarkdownDescription: "The minimum number of seconds for which a newly created Pod should be ready without any of its containers crashing, for it to be considered available",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "Override pod resource requirements for the Kubernetes Sensor pods.",
										MarkdownDescription: "Override pod resource requirements for the Kubernetes Sensor pods.",
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

									"replicas": schema.Int64Attribute{
										Description:         "Specify the number of replicas for the Kubernetes Sensor.",
										MarkdownDescription: "Specify the number of replicas for the Kubernetes Sensor.",
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

					"openshift": schema.BoolAttribute{
						Description:         "Set to 'True' to indicate the Operator is being deployed in a OpenShift cluster. Provides a hint so that RBAC etc is configured correctly.",
						MarkdownDescription: "Set to 'True' to indicate the Operator is being deployed in a OpenShift cluster. Provides a hint so that RBAC etc is configured correctly.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"opentelemetry": schema.SingleNestedAttribute{
						Description:         "Enables the OpenTelemetry gRPC endpoint on the Agent. If true, it will also apply 'service.create: true'.",
						MarkdownDescription: "Enables the OpenTelemetry gRPC endpoint on the Agent. If true, it will also apply 'service.create: true'.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grpc": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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

							"http": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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

					"opentelemetry_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pinned_chart_version": schema.StringAttribute{
						Description:         "Specifying the PinnedChartVersion allows for 'pinning' the Helm Chart used by the Operator for installing the Agent DaemonSet. Normally the Operator will always install and update to the latest Helm Chart version. The Operator will check and make sure no 'unsupported' Chart versions can be selected.",
						MarkdownDescription: "Specifying the PinnedChartVersion allows for 'pinning' the Helm Chart used by the Operator for installing the Agent DaemonSet. Normally the Operator will always install and update to the latest Helm Chart version. The Operator will check and make sure no 'unsupported' Chart versions can be selected.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_security_policy": schema.SingleNestedAttribute{
						Description:         "Specify a PodSecurityPolicy for the Instana Agent Pods. If enabled requires 'rbac.create: true'.",
						MarkdownDescription: "Specify a PodSecurityPolicy for the Instana Agent Pods. If enabled requires 'rbac.create: true'.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
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

					"prometheus": schema.SingleNestedAttribute{
						Description:         "Enables the Prometheus endpoint on the Agent. If true, it will also apply 'service.create: true'.",
						MarkdownDescription: "Enables the Prometheus endpoint on the Agent. If true, it will also apply 'service.create: true'.",
						Attributes: map[string]schema.Attribute{
							"remote_write": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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

					"rbac": schema.SingleNestedAttribute{
						Description:         "Specifies whether RBAC resources should be created.",
						MarkdownDescription: "Specifies whether RBAC resources should be created.",
						Attributes: map[string]schema.Attribute{
							"create": schema.BoolAttribute{
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

					"service": schema.SingleNestedAttribute{
						Description:         "Specifies whether to create the instana-agent 'Service' to expose within the cluster. The Service can then be used e.g. for the Prometheus remote-write, OpenTelemetry GRCP endpoint and other APIs. Note: Requires Kubernetes 1.17+, as it uses topologyKeys.",
						MarkdownDescription: "Specifies whether to create the instana-agent 'Service' to expose within the cluster. The Service can then be used e.g. for the Prometheus remote-write, OpenTelemetry GRCP endpoint and other APIs. Note: Requires Kubernetes 1.17+, as it uses topologyKeys.",
						Attributes: map[string]schema.Attribute{
							"create": schema.BoolAttribute{
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

					"service_account": schema.SingleNestedAttribute{
						Description:         "Specifies whether a ServiceAccount should be created (default 'true'), and possibly the name to use.",
						MarkdownDescription: "Specifies whether a ServiceAccount should be created (default 'true'), and possibly the name to use.",
						Attributes: map[string]schema.Attribute{
							"create": schema.BoolAttribute{
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

					"zone": schema.SingleNestedAttribute{
						Description:         "Name of the zone in which the host(s) will be displayed on the map. Optional, but then 'cluster.name' must be specified.",
						MarkdownDescription: "Name of the zone in which the host(s) will be displayed on the map. Optional, but then 'cluster.name' must be specified.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InstanaIoInstanaAgentV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_instana_io_instana_agent_v1_manifest")

	var model InstanaIoInstanaAgentV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("instana.io/v1")
	model.Kind = pointer.String("InstanaAgent")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
