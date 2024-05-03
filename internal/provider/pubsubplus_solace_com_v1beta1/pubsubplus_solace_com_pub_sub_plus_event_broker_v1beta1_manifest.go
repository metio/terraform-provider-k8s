/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pubsubplus_solace_com_v1beta1

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
	_ datasource.DataSource = &PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest{}
)

func NewPubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest() datasource.DataSource {
	return &PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest{}
}

type PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest struct{}

type PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1ManifestData struct {
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
		AdminCredentialsSecret *string `tfsdk:"admin_credentials_secret" json:"adminCredentialsSecret,omitempty"`
		Developer              *bool   `tfsdk:"developer" json:"developer,omitempty"`
		ExtraEnvVars           *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"extra_env_vars" json:"extraEnvVars,omitempty"`
		ExtraEnvVarsCM     *string `tfsdk:"extra_env_vars_cm" json:"extraEnvVarsCM,omitempty"`
		ExtraEnvVarsSecret *string `tfsdk:"extra_env_vars_secret" json:"extraEnvVarsSecret,omitempty"`
		Image              *struct {
			PullPolicy  *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			PullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		Monitoring *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Image   *struct {
				PullPolicy  *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				PullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			IncludeRates    *bool `tfsdk:"include_rates" json:"includeRates,omitempty"`
			MetricsEndpoint *struct {
				ContainerPort                   *float64 `tfsdk:"container_port" json:"containerPort,omitempty"`
				EndpointTlsConfigPrivateKeyName *string  `tfsdk:"endpoint_tls_config_private_key_name" json:"endpointTlsConfigPrivateKeyName,omitempty"`
				EndpointTlsConfigSecret         *string  `tfsdk:"endpoint_tls_config_secret" json:"endpointTlsConfigSecret,omitempty"`
				EndpointTlsConfigServerCertName *string  `tfsdk:"endpoint_tls_config_server_cert_name" json:"endpointTlsConfigServerCertName,omitempty"`
				ListenTLS                       *bool    `tfsdk:"listen_tls" json:"listenTLS,omitempty"`
				Name                            *string  `tfsdk:"name" json:"name,omitempty"`
				Protocol                        *string  `tfsdk:"protocol" json:"protocol,omitempty"`
				ServicePort                     *float64 `tfsdk:"service_port" json:"servicePort,omitempty"`
				ServiceType                     *string  `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"metrics_endpoint" json:"metricsEndpoint,omitempty"`
			SslVerify *bool    `tfsdk:"ssl_verify" json:"sslVerify,omitempty"`
			TimeOut   *float64 `tfsdk:"time_out" json:"timeOut,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		MonitoringCredentialsSecret *string `tfsdk:"monitoring_credentials_secret" json:"monitoringCredentialsSecret,omitempty"`
		NodeAssignment              *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Spec *struct {
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
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
							MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
							MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
							MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
							MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"node_assignment" json:"nodeAssignment,omitempty"`
		PodAnnotations           *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
		PodDisruptionBudgetForHA *bool              `tfsdk:"pod_disruption_budget_for_ha" json:"podDisruptionBudgetForHA,omitempty"`
		PodLabels                *map[string]string `tfsdk:"pod_labels" json:"podLabels,omitempty"`
		PreSharedAuthKeySecret   *string            `tfsdk:"pre_shared_auth_key_secret" json:"preSharedAuthKeySecret,omitempty"`
		Redundancy               *bool              `tfsdk:"redundancy" json:"redundancy,omitempty"`
		SecurityContext          *struct {
			FsGroup   *float64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
			RunAsUser *float64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
		} `tfsdk:"security_context" json:"securityContext,omitempty"`
		Service *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Ports       *[]struct {
				ContainerPort *float64 `tfsdk:"container_port" json:"containerPort,omitempty"`
				Name          *string  `tfsdk:"name" json:"name,omitempty"`
				Protocol      *string  `tfsdk:"protocol" json:"protocol,omitempty"`
				ServicePort   *float64 `tfsdk:"service_port" json:"servicePort,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		ServiceAccount *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		Storage *struct {
			CustomVolumeMount *[]struct {
				Name                  *string `tfsdk:"name" json:"name,omitempty"`
				PersistentVolumeClaim *struct {
					ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
				} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
			} `tfsdk:"custom_volume_mount" json:"customVolumeMount,omitempty"`
			MessagingNodeStorageSize *string `tfsdk:"messaging_node_storage_size" json:"messagingNodeStorageSize,omitempty"`
			MonitorNodeStorageSize   *string `tfsdk:"monitor_node_storage_size" json:"monitorNodeStorageSize,omitempty"`
			Slow                     *bool   `tfsdk:"slow" json:"slow,omitempty"`
			UseStorageClass          *string `tfsdk:"use_storage_class" json:"useStorageClass,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		SystemScaling *struct {
			MaxConnections      *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
			MaxQueueMessages    *int64  `tfsdk:"max_queue_messages" json:"maxQueueMessages,omitempty"`
			MaxSpoolUsage       *int64  `tfsdk:"max_spool_usage" json:"maxSpoolUsage,omitempty"`
			MessagingNodeCpu    *string `tfsdk:"messaging_node_cpu" json:"messagingNodeCpu,omitempty"`
			MessagingNodeMemory *string `tfsdk:"messaging_node_memory" json:"messagingNodeMemory,omitempty"`
		} `tfsdk:"system_scaling" json:"systemScaling,omitempty"`
		Timezone *string `tfsdk:"timezone" json:"timezone,omitempty"`
		Tls      *struct {
			CertFilename          *string `tfsdk:"cert_filename" json:"certFilename,omitempty"`
			CertKeyFilename       *string `tfsdk:"cert_key_filename" json:"certKeyFilename,omitempty"`
			Enabled               *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			ServerTlsConfigSecret *string `tfsdk:"server_tls_config_secret" json:"serverTlsConfigSecret,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		UpdateStrategy *string `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest"
}

func (r *PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PubSub+ Event Broker",
		MarkdownDescription: "PubSub+ Event Broker",
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
				Description:         "EventBrokerSpec defines the desired state of PubSubPlusEventBroker",
				MarkdownDescription: "EventBrokerSpec defines the desired state of PubSubPlusEventBroker",
				Attributes: map[string]schema.Attribute{
					"admin_credentials_secret": schema.StringAttribute{
						Description:         "Defines the password for PubSubPlusEventBroker if provided. Random one will be generated if not provided.When provided, ensure the secret key name is 'username_admin_password'. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.",
						MarkdownDescription: "Defines the password for PubSubPlusEventBroker if provided. Random one will be generated if not provided.When provided, ensure the secret key name is 'username_admin_password'. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"developer": schema.BoolAttribute{
						Description:         "Developer true specifies a minimum footprint scaled-down deployment, not for production use.If set to true it overrides SystemScaling parameters.",
						MarkdownDescription: "Developer true specifies a minimum footprint scaled-down deployment, not for production use.If set to true it overrides SystemScaling parameters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_env_vars": schema.ListNestedAttribute{
						Description:         "List of extra environment variables to be added to the PubSubPlusEventBroker container. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.A primary use case is to specify configuration keys, although the variables defined here will not override the ones defined in ConfigMap",
						MarkdownDescription: "List of extra environment variables to be added to the PubSubPlusEventBroker container. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.A primary use case is to specify configuration keys, although the variables defined here will not override the ones defined in ConfigMap",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Specifies the Name of an environment variable to be added to the PubSubPlusEventBroker container",
									MarkdownDescription: "Specifies the Name of an environment variable to be added to the PubSubPlusEventBroker container",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Specifies the Value of an environment variable to be added to the PubSubPlusEventBroker container",
									MarkdownDescription: "Specifies the Value of an environment variable to be added to the PubSubPlusEventBroker container",
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

					"extra_env_vars_cm": schema.StringAttribute{
						Description:         "List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing ConfigMap. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.",
						MarkdownDescription: "List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing ConfigMap. Note: Do not configure Timezone or SystemScaling parameters here as it could cause unintended consequences.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_env_vars_secret": schema.StringAttribute{
						Description:         "List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing Secret",
						MarkdownDescription: "List of extra environment variables to be added to the PubSubPlusEventBroker container from an existing Secret",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image defines container image parameters for the event broker.",
						MarkdownDescription: "Image defines container image parameters for the event broker.",
						Attributes: map[string]schema.Attribute{
							"pull_policy": schema.StringAttribute{
								Description:         "Specifies ImagePullPolicy of the container image for the event broker.",
								MarkdownDescription: "Specifies ImagePullPolicy of the container image for the event broker.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pull_secrets": schema.ListNestedAttribute{
								Description:         "pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								MarkdownDescription: "pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
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

							"repository": schema.StringAttribute{
								Description:         "Defines the container image repo where the event broker image is pulled from",
								MarkdownDescription: "Defines the container image repo where the event broker image is pulled from",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Specifies the tag of the container image to be used for the event broker.",
								MarkdownDescription: "Specifies the tag of the container image to be used for the event broker.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "Monitoring specifies a Prometheus monitoring endpoint for the event broker",
						MarkdownDescription: "Monitoring specifies a Prometheus monitoring endpoint for the event broker",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled true enables the setup of the Prometheus Exporter.",
								MarkdownDescription: "Enabled true enables the setup of the Prometheus Exporter.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image defines container image parameters for the Prometheus Exporter.",
								MarkdownDescription: "Image defines container image parameters for the Prometheus Exporter.",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "Specifies ImagePullPolicy of the container image for the Prometheus Exporter.",
										MarkdownDescription: "Specifies ImagePullPolicy of the container image for the Prometheus Exporter.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pull_secrets": schema.ListNestedAttribute{
										Description:         "pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
										MarkdownDescription: "pullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
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

									"repository": schema.StringAttribute{
										Description:         "Defines the container image repo where the Prometheus Exporter image is pulled from",
										MarkdownDescription: "Defines the container image repo where the Prometheus Exporter image is pulled from",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Specifies the tag of the container image to be used for the Prometheus Exporter.",
										MarkdownDescription: "Specifies the tag of the container image to be used for the Prometheus Exporter.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"include_rates": schema.BoolAttribute{
								Description:         "Defines if Prometheus Exporter should include rates",
								MarkdownDescription: "Defines if Prometheus Exporter should include rates",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_endpoint": schema.SingleNestedAttribute{
								Description:         "MetricsEndpoint defines parameters to configure monitoring for the Prometheus Exporter.",
								MarkdownDescription: "MetricsEndpoint defines parameters to configure monitoring for the Prometheus Exporter.",
								Attributes: map[string]schema.Attribute{
									"container_port": schema.Float64Attribute{
										Description:         "ContainerPort is the port number to expose on the Prometheus Exporter pod.",
										MarkdownDescription: "ContainerPort is the port number to expose on the Prometheus Exporter pod.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_tls_config_private_key_name": schema.StringAttribute{
										Description:         "EndpointTlsConfigPrivateKeyName is the file name of the Private Key used to set up TLS configuration",
										MarkdownDescription: "EndpointTlsConfigPrivateKeyName is the file name of the Private Key used to set up TLS configuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_tls_config_secret": schema.StringAttribute{
										Description:         "EndpointTLSConfigSecret defines TLS secret name to set up TLS configuration",
										MarkdownDescription: "EndpointTLSConfigSecret defines TLS secret name to set up TLS configuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_tls_config_server_cert_name": schema.StringAttribute{
										Description:         "EndpointTlsConfigServerCertName is the file name of the Server Certificate used to set up TLS configuration",
										MarkdownDescription: "EndpointTlsConfigServerCertName is the file name of the Server Certificate used to set up TLS configuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"listen_tls": schema.BoolAttribute{
										Description:         "Defines if Metrics Service Endpoint uses TLS configuration",
										MarkdownDescription: "Defines if Metrics Service Endpoint uses TLS configuration",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is a unique name for the port that can be referred to by services.",
										MarkdownDescription: "Name is a unique name for the port that can be referred to by services.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"protocol": schema.StringAttribute{
										Description:         "Protocol for port. Must be UDP, TCP, or SCTP.",
										MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("TCP", "UDP", "SCTP"),
										},
									},

									"service_port": schema.Float64Attribute{
										Description:         "ServicePort is the port number to expose on the service",
										MarkdownDescription: "ServicePort is the port number to expose on the service",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_type": schema.StringAttribute{
										Description:         "Defines the service type for the Metrics Service Endpoint",
										MarkdownDescription: "Defines the service type for the Metrics Service Endpoint",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_verify": schema.BoolAttribute{
								Description:         "Defines if Prometheus Exporter verifies SSL",
								MarkdownDescription: "Defines if Prometheus Exporter verifies SSL",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_out": schema.Float64Attribute{
								Description:         "Timeout configuration for Prometheus Exporter scrapper",
								MarkdownDescription: "Timeout configuration for Prometheus Exporter scrapper",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring_credentials_secret": schema.StringAttribute{
						Description:         "Defines the password for PubSubPlusEventBroker to be used by the Exporter for monitoring.When provided, ensure the secret key name is 'username_monitor_password'. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.",
						MarkdownDescription: "Defines the password for PubSubPlusEventBroker to be used by the Exporter for monitoring.When provided, ensure the secret key name is 'username_monitor_password'. For valid values refer to the Solace documentation https://docs.solace.com/Admin/Configuring-Internal-CLI-User-Accounts.htm.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_assignment": schema.ListNestedAttribute{
						Description:         "NodeAssignment defines labels to constrain PubSubPlusEventBroker nodes to run on particular node(s), or to prefer to run on particular nodes.",
						MarkdownDescription: "NodeAssignment defines labels to constrain PubSubPlusEventBroker nodes to run on particular node(s), or to prefer to run on particular nodes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Defines the name of broker node type that has the nodeAssignment spec defined",
									MarkdownDescription: "Defines the name of broker node type that has the nodeAssignment spec defined",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Primary", "Backup", "Monitor"),
									},
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "If provided defines the labels to constrain the PubSubPlusEventBroker node to specific nodes",
									MarkdownDescription: "If provided defines the labels to constrain the PubSubPlusEventBroker node to specific nodes",
									Attributes: map[string]schema.Attribute{
										"affinity": schema.SingleNestedAttribute{
											Description:         "Affinity if provided defines the conditional approach to assign PubSubPlusEventBroker nodes to specific nodes to which they can be scheduled",
											MarkdownDescription: "Affinity if provided defines the conditional approach to assign PubSubPlusEventBroker nodes to specific nodes to which they can be scheduled",
											Attributes: map[string]schema.Attribute{
												"node_affinity": schema.SingleNestedAttribute{
													Description:         "Describes node affinity scheduling rules for the pod.",
													MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
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
																							Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																							Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
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
																							Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																							Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																			"match_label_keys": schema.ListAttribute{
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																			"namespaces": schema.ListAttribute{
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																		Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
															Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																	"match_label_keys": schema.ListAttribute{
																		Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"mismatch_label_keys": schema.ListAttribute{
																		Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																	"namespaces": schema.ListAttribute{
																		Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																			"match_label_keys": schema.ListAttribute{
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																			"namespaces": schema.ListAttribute{
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																		Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
															Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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

																	"match_label_keys": schema.ListAttribute{
																		Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"mismatch_label_keys": schema.ListAttribute{
																		Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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

																	"namespaces": schema.ListAttribute{
																		Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

										"node_selector": schema.MapAttribute{
											Description:         "NodeSelector if provided defines the exact labels of nodes to which PubSubPlusEventBroker nodes can be scheduled",
											MarkdownDescription: "NodeSelector if provided defines the exact labels of nodes to which PubSubPlusEventBroker nodes can be scheduled",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.ListNestedAttribute{
											Description:         "Toleration if provided defines the exact properties of the PubSubPlusEventBroker nodes can be scheduled on nodes with d matching taint.",
											MarkdownDescription: "Toleration if provided defines the exact properties of the PubSubPlusEventBroker nodes can be scheduled on nodes with d matching taint.",
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_annotations": schema.MapAttribute{
						Description:         "PodAnnotations allows adding provider-specific pod annotations to PubSubPlusEventBroker pods",
						MarkdownDescription: "PodAnnotations allows adding provider-specific pod annotations to PubSubPlusEventBroker pods",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_disruption_budget_for_ha": schema.BoolAttribute{
						Description:         "PodDisruptionBudgetForHA enables setting up PodDisruptionBudget for the broker pods in HA deployment.This parameter is ignored for non-HA deployments (if redundancy is false).",
						MarkdownDescription: "PodDisruptionBudgetForHA enables setting up PodDisruptionBudget for the broker pods in HA deployment.This parameter is ignored for non-HA deployments (if redundancy is false).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_labels": schema.MapAttribute{
						Description:         "PodLabels allows adding provider-specific pod labels to PubSubPlusEventBroker pods",
						MarkdownDescription: "PodLabels allows adding provider-specific pod labels to PubSubPlusEventBroker pods",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_shared_auth_key_secret": schema.StringAttribute{
						Description:         "PreSharedAuthKeySecret defines the PreSharedAuthKey Secret for PubSubPlusEventBroker. Random one will be generated if not provided.When provided, ensure the secret key name is 'preshared_auth_key'. For valid values refer to the Solace documentation https://docs.solace.com/Features/HA-Redundancy/Pre-Shared-Keys-SMB.htm?Highlight=pre%20shared.",
						MarkdownDescription: "PreSharedAuthKeySecret defines the PreSharedAuthKey Secret for PubSubPlusEventBroker. Random one will be generated if not provided.When provided, ensure the secret key name is 'preshared_auth_key'. For valid values refer to the Solace documentation https://docs.solace.com/Features/HA-Redundancy/Pre-Shared-Keys-SMB.htm?Highlight=pre%20shared.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redundancy": schema.BoolAttribute{
						Description:         "Redundancy true specifies HA deployment, false specifies Non-HA.",
						MarkdownDescription: "Redundancy true specifies HA deployment, false specifies Non-HA.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext defines the pod security context for the event broker.",
						MarkdownDescription: "SecurityContext defines the pod security context for the event broker.",
						Attributes: map[string]schema.Attribute{
							"fs_group": schema.Float64Attribute{
								Description:         "Specifies fsGroup in pod security context. 0 or unset defaults either to 1000002, or if OpenShift detected to unspecified (see documentation)",
								MarkdownDescription: "Specifies fsGroup in pod security context. 0 or unset defaults either to 1000002, or if OpenShift detected to unspecified (see documentation)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Float64Attribute{
								Description:         "Specifies runAsUser in pod security context. 0 or unset defaults either to 1000001, or if OpenShift detected to unspecified (see documentation)",
								MarkdownDescription: "Specifies runAsUser in pod security context. 0 or unset defaults either to 1000001, or if OpenShift detected to unspecified (see documentation)",
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
						Description:         "Service defines broker service details.",
						MarkdownDescription: "Service defines broker service details.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations allows adding provider-specific service annotations",
								MarkdownDescription: "Annotations allows adding provider-specific service annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ports": schema.ListNestedAttribute{
								Description:         "Ports specifies the ports to expose PubSubPlusEventBroker services.",
								MarkdownDescription: "Ports specifies the ports to expose PubSubPlusEventBroker services.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"container_port": schema.Float64Attribute{
											Description:         "Port number to expose on the pod.",
											MarkdownDescription: "Port number to expose on the pod.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Unique name for the port that can be referred to by services.",
											MarkdownDescription: "Unique name for the port that can be referred to by services.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol for port. Must be UDP, TCP, or SCTP.",
											MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", "UDP", "SCTP"),
											},
										},

										"service_port": schema.Float64Attribute{
											Description:         "Port number to expose on the service",
											MarkdownDescription: "Port number to expose on the service",
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

							"type": schema.StringAttribute{
								Description:         "ServiceType specifies how to expose the broker services. Options include ClusterIP, NodePort, LoadBalancer (default).",
								MarkdownDescription: "ServiceType specifies how to expose the broker services. Options include ClusterIP, NodePort, LoadBalancer (default).",
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
						Description:         "ServiceAccount defines a ServiceAccount dedicated to the PubSubPlusEventBroker",
						MarkdownDescription: "ServiceAccount defines a ServiceAccount dedicated to the PubSubPlusEventBroker",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name specifies the name of an existing ServiceAccount dedicated to the PubSubPlusEventBroker.If this value is missing a new ServiceAccount will be created.",
								MarkdownDescription: "Name specifies the name of an existing ServiceAccount dedicated to the PubSubPlusEventBroker.If this value is missing a new ServiceAccount will be created.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage defines storage details for the broker.",
						MarkdownDescription: "Storage defines storage details for the broker.",
						Attributes: map[string]schema.Attribute{
							"custom_volume_mount": schema.ListNestedAttribute{
								Description:         "CustomVolumeMount can be used to show the data volume should be mounted instead of using a storage class.",
								MarkdownDescription: "CustomVolumeMount can be used to show the data volume should be mounted instead of using a storage class.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Defines the name of PubSubPlusEventBroker node type that has the customVolumeMount spec defined",
											MarkdownDescription: "Defines the name of PubSubPlusEventBroker node type that has the customVolumeMount spec defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Primary", "Backup", "Monitor"),
											},
										},

										"persistent_volume_claim": schema.SingleNestedAttribute{
											Description:         "Defines the customVolumeMount that can be used mount the data volume instead of using a storage class",
											MarkdownDescription: "Defines the customVolumeMount that can be used mount the data volume instead of using a storage class",
											Attributes: map[string]schema.Attribute{
												"claim_name": schema.StringAttribute{
													Description:         "Defines the claimName of a custom PersistentVolumeClaim to be used instead",
													MarkdownDescription: "Defines the claimName of a custom PersistentVolumeClaim to be used instead",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"messaging_node_storage_size": schema.StringAttribute{
								Description:         "MessagingNodeStorageSize if provided will assign the minimum persistent storage to be used by the message nodes.",
								MarkdownDescription: "MessagingNodeStorageSize if provided will assign the minimum persistent storage to be used by the message nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"monitor_node_storage_size": schema.StringAttribute{
								Description:         "MonitorNodeStorageSize if provided this will create and assign the minimum recommended storage to Monitor pods.",
								MarkdownDescription: "MonitorNodeStorageSize if provided this will create and assign the minimum recommended storage to Monitor pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"slow": schema.BoolAttribute{
								Description:         "Slow indicate slow storage is in use, an example is NFS.",
								MarkdownDescription: "Slow indicate slow storage is in use, an example is NFS.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_storage_class": schema.StringAttribute{
								Description:         "UseStrorageClass Name of the StorageClass to be used to request persistent storage volumes. If undefined, the 'default' StorageClass will be used.",
								MarkdownDescription: "UseStrorageClass Name of the StorageClass to be used to request persistent storage volumes. If undefined, the 'default' StorageClass will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"system_scaling": schema.SingleNestedAttribute{
						Description:         "SystemScaling provides exact fine-grained specification of the event broker scaling parametersand the assigned CPU / memory resources to the Pod.",
						MarkdownDescription: "SystemScaling provides exact fine-grained specification of the event broker scaling parametersand the assigned CPU / memory resources to the Pod.",
						Attributes: map[string]schema.Attribute{
							"max_connections": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_queue_messages": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_spool_usage": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"messaging_node_cpu": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"messaging_node_memory": schema.StringAttribute{
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

					"timezone": schema.StringAttribute{
						Description:         "Defines the timezone for the event broker container, if undefined default is UTC. Valid values are tz database time zone names.",
						MarkdownDescription: "Defines the timezone for the event broker container, if undefined default is UTC. Valid values are tz database time zone names.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS provides TLS configuration for the event broker.",
						MarkdownDescription: "TLS provides TLS configuration for the event broker.",
						Attributes: map[string]schema.Attribute{
							"cert_filename": schema.StringAttribute{
								Description:         "Name of the Certificate file in the 'serverCertificatesSecret'",
								MarkdownDescription: "Name of the Certificate file in the 'serverCertificatesSecret'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_key_filename": schema.StringAttribute{
								Description:         "Name of the Key file in the 'serverCertificatesSecret'",
								MarkdownDescription: "Name of the Key file in the 'serverCertificatesSecret'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled true enables TLS for the broker.",
								MarkdownDescription: "Enabled true enables TLS for the broker.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_tls_config_secret": schema.StringAttribute{
								Description:         "Specifies the tls configuration secret to be used for the broker",
								MarkdownDescription: "Specifies the tls configuration secret to be used for the broker",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"update_strategy": schema.StringAttribute{
						Description:         "UpdateStrategy specifies how to update an existing deployment. manualPodRestart waits for user intervention.",
						MarkdownDescription: "UpdateStrategy specifies how to update an existing deployment. manualPodRestart waits for user intervention.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("automatedRolling", "manualPodRestart"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest")

	var model PubsubplusSolaceComPubSubPlusEventBrokerV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("pubsubplus.solace.com/v1beta1")
	model.Kind = pointer.String("PubSubPlusEventBroker")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
