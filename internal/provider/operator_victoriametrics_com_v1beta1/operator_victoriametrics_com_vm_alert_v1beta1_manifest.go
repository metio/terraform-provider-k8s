/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorVictoriametricsComVmalertV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmalertV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmalertV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmalertV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmalertV1Beta1ManifestData struct {
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
		Affinity                *map[string]string `tfsdk:"affinity" json:"affinity,omitempty"`
		ConfigMaps              *[]string          `tfsdk:"config_maps" json:"configMaps,omitempty"`
		ConfigReloaderExtraArgs *map[string]string `tfsdk:"config_reloader_extra_args" json:"configReloaderExtraArgs,omitempty"`
		ConfigReloaderImageTag  *string            `tfsdk:"config_reloader_image_tag" json:"configReloaderImageTag,omitempty"`
		ConfigReloaderResources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"config_reloader_resources" json:"configReloaderResources,omitempty"`
		Containers *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
		Datasource *struct {
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Headers   *[]string          `tfsdk:"headers" json:"headers,omitempty"`
			Oauth2    *map[string]string `tfsdk:"oauth2" json:"oauth2,omitempty"`
			TlsConfig *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url       *string            `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"datasource" json:"datasource,omitempty"`
		DisableSelfServiceScrape *bool `tfsdk:"disable_self_service_scrape" json:"disableSelfServiceScrape,omitempty"`
		DnsConfig                *struct {
			Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
			Options     *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
		} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
		DnsPolicy              *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
		EnforcedNamespaceLabel *string              `tfsdk:"enforced_namespace_label" json:"enforcedNamespaceLabel,omitempty"`
		EvaluationInterval     *string              `tfsdk:"evaluation_interval" json:"evaluationInterval,omitempty"`
		ExternalLabels         *map[string]string   `tfsdk:"external_labels" json:"externalLabels,omitempty"`
		ExtraArgs              *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
		ExtraEnvs              *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
		HostAliases            *[]struct {
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
		InitContainers *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
		License        *struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			KeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		} `tfsdk:"license" json:"license,omitempty"`
		LivenessProbe   *map[string]string `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		LogFormat       *string            `tfsdk:"log_format" json:"logFormat,omitempty"`
		LogLevel        *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		MinReadySeconds *int64             `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
		NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Notifier        *struct {
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Headers  *[]string          `tfsdk:"headers" json:"headers,omitempty"`
			Oauth2   *map[string]string `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Selector *struct {
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				NamespaceSelector *struct {
					Any        *bool     `tfsdk:"any" json:"any,omitempty"`
					MatchNames *[]string `tfsdk:"match_names" json:"matchNames,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			TlsConfig *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url       *string            `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"notifier" json:"notifier,omitempty"`
		NotifierConfigRef *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"notifier_config_ref" json:"notifierConfigRef,omitempty"`
		Notifiers *[]struct {
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Headers  *[]string          `tfsdk:"headers" json:"headers,omitempty"`
			Oauth2   *map[string]string `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Selector *struct {
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				NamespaceSelector *struct {
					Any        *bool     `tfsdk:"any" json:"any,omitempty"`
					MatchNames *[]string `tfsdk:"match_names" json:"matchNames,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			TlsConfig *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url       *string            `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"notifiers" json:"notifiers,omitempty"`
		Paused              *bool `tfsdk:"paused" json:"paused,omitempty"`
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
		Port              *string `tfsdk:"port" json:"port,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ReadinessGates    *[]struct {
			ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
		} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
		ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		RemoteRead     *struct {
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Headers   *[]string          `tfsdk:"headers" json:"headers,omitempty"`
			Lookback  *string            `tfsdk:"lookback" json:"lookback,omitempty"`
			Oauth2    *map[string]string `tfsdk:"oauth2" json:"oauth2,omitempty"`
			TlsConfig *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url       *string            `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"remote_read" json:"remoteRead,omitempty"`
		RemoteWrite *struct {
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
				Username      *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			BearerTokenFile   *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Concurrency   *int64             `tfsdk:"concurrency" json:"concurrency,omitempty"`
			FlushInterval *string            `tfsdk:"flush_interval" json:"flushInterval,omitempty"`
			Headers       *[]string          `tfsdk:"headers" json:"headers,omitempty"`
			MaxBatchSize  *int64             `tfsdk:"max_batch_size" json:"maxBatchSize,omitempty"`
			MaxQueueSize  *int64             `tfsdk:"max_queue_size" json:"maxQueueSize,omitempty"`
			Oauth2        *map[string]string `tfsdk:"oauth2" json:"oauth2,omitempty"`
			TlsConfig     *map[string]string `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url           *string            `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"remote_write" json:"remoteWrite,omitempty"`
		ReplicaCount *int64 `tfsdk:"replica_count" json:"replicaCount,omitempty"`
		Resources    *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RevisionHistoryLimitCount *int64 `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
		RollingUpdate             *struct {
			MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
			MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
		} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
		RuleNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"rule_namespace_selector" json:"ruleNamespaceSelector,omitempty"`
		RulePath     *[]string `tfsdk:"rule_path" json:"rulePath,omitempty"`
		RuleSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"rule_selector" json:"ruleSelector,omitempty"`
		RuntimeClassName   *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
		SchedulerName      *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		Secrets            *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
		SecurityContext    *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
		SelectAllByDefault *bool              `tfsdk:"select_all_by_default" json:"selectAllByDefault,omitempty"`
		ServiceAccountName *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		ServiceScrapeSpec  *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
		ServiceSpec        *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
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
		UpdateStrategy            *string              `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		UseDefaultResources       *bool                `tfsdk:"use_default_resources" json:"useDefaultResources,omitempty"`
		UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		UseVMConfigReloader       *bool                `tfsdk:"use_vm_config_reloader" json:"useVMConfigReloader,omitempty"`
		VolumeMounts              *[]struct {
			MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
			SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmalertV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_alert_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmalertV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMAlert executes a list of given alerting or recording rules against configured address.",
		MarkdownDescription: "VMAlert executes a list of given alerting or recording rules against configured address.",
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
				Description:         "VMAlertSpec defines the desired state of VMAlert",
				MarkdownDescription: "VMAlertSpec defines the desired state of VMAlert",
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
						Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
						MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_reloader_extra_args": schema.MapAttribute{
						Description:         "ConfigReloaderExtraArgs that will be passed to VMAuths config-reloader container for example resyncInterval: '30s'",
						MarkdownDescription: "ConfigReloaderExtraArgs that will be passed to VMAuths config-reloader container for example resyncInterval: '30s'",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_reloader_image_tag": schema.StringAttribute{
						Description:         "ConfigReloaderImageTag defines image:tag for config-reloader container",
						MarkdownDescription: "ConfigReloaderImageTag defines image:tag for config-reloader container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_reloader_resources": schema.SingleNestedAttribute{
						Description:         "ConfigReloaderResources config-reloader container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
						MarkdownDescription: "ConfigReloaderResources config-reloader container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
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

					"containers": schema.ListAttribute{
						Description:         "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
						MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datasource": schema.SingleNestedAttribute{
						Description:         "Datasource Victoria Metrics or VMSelect url. Required parameter. e.g. http://127.0.0.1:8428",
						MarkdownDescription: "Datasource Victoria Metrics or VMSelect url. Required parameter. e.g. http://127.0.0.1:8428",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
								Attributes: map[string]schema.Attribute{
									"password": schema.SingleNestedAttribute{
										Description:         "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"password_file": schema.StringAttribute{
										Description:         "PasswordFile defines path to password file at disk must be pre-mounted",
										MarkdownDescription: "PasswordFile defines path to password file at disk must be pre-mounted",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"bearer_token_file": schema.StringAttribute{
								Description:         "Path to bearer token file",
								MarkdownDescription: "Path to bearer token file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bearer_token_secret": schema.SingleNestedAttribute{
								Description:         "Optional bearer auth token to use for -remoteWrite.url",
								MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from. Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"headers": schema.ListAttribute{
								Description:         "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								MarkdownDescription: "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2": schema.MapAttribute{
								Description:         "OAuth2 defines OAuth2 configuration",
								MarkdownDescription: "OAuth2 defines OAuth2 configuration",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_config": schema.MapAttribute{
								Description:         "TLSConfig specifies TLSConfig configuration parameters.",
								MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "Victoria Metrics or VMSelect url. Required parameter. E.g. http://127.0.0.1:8428",
								MarkdownDescription: "Victoria Metrics or VMSelect url. Required parameter. E.g. http://127.0.0.1:8428",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"disable_self_service_scrape": schema.BoolAttribute{
						Description:         "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
						MarkdownDescription: "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dns_config": schema.SingleNestedAttribute{
						Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
						MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
						Attributes: map[string]schema.Attribute{
							"nameservers": schema.ListAttribute{
								Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
								MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"options": schema.ListNestedAttribute{
								Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
								MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
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
								Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
								MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
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

					"enforced_namespace_label": schema.StringAttribute{
						Description:         "EnforcedNamespaceLabel enforces adding a namespace label of origin for each alert and metric that is user created. The label value will always be the namespace of the object that is being created.",
						MarkdownDescription: "EnforcedNamespaceLabel enforces adding a namespace label of origin for each alert and metric that is user created. The label value will always be the namespace of the object that is being created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"evaluation_interval": schema.StringAttribute{
						Description:         "EvaluationInterval defines how often to evaluate rules by default",
						MarkdownDescription: "EvaluationInterval defines how often to evaluate rules by default",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
						},
					},

					"external_labels": schema.MapAttribute{
						Description:         "ExternalLabels in the form 'name: value' to add to all generated recording rules and alerts.",
						MarkdownDescription: "ExternalLabels in the form 'name: value' to add to all generated recording rules and alerts.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_args": schema.MapAttribute{
						Description:         "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
						MarkdownDescription: "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_envs": schema.ListAttribute{
						Description:         "ExtraEnvs that will be passed to the application container",
						MarkdownDescription: "ExtraEnvs that will be passed to the application container",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
						MarkdownDescription: "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
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

					"host_network": schema.BoolAttribute{
						Description:         "HostNetwork controls whether the pod may use the node network namespace",
						MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image - docker image settings if no specified operator uses default version from operator config",
						MarkdownDescription: "Image - docker image settings if no specified operator uses default version from operator config",
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
						Description:         "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
									MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

					"init_containers": schema.ListAttribute{
						Description:         "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
						MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"license": schema.SingleNestedAttribute{
						Description:         "License allows to configure license key to be used for enterprise features. Using license key is supported starting from VictoriaMetrics v1.94.0. See [here](https://docs.victoriametrics.com/enterprise)",
						MarkdownDescription: "License allows to configure license key to be used for enterprise features. Using license key is supported starting from VictoriaMetrics v1.94.0. See [here](https://docs.victoriametrics.com/enterprise)",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Enterprise license key. This flag is available only in [VictoriaMetrics enterprise](https://docs.victoriametrics.com/enterprise). To request a trial license, [go to](https://victoriametrics.com/products/enterprise/trial)",
								MarkdownDescription: "Enterprise license key. This flag is available only in [VictoriaMetrics enterprise](https://docs.victoriametrics.com/enterprise). To request a trial license, [go to](https://victoriametrics.com/products/enterprise/trial)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_ref": schema.SingleNestedAttribute{
								Description:         "KeyRef is reference to secret with license key for enterprise features.",
								MarkdownDescription: "KeyRef is reference to secret with license key for enterprise features.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from. Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
						Description:         "LogFormat for VMAlert to be configured with. default or json",
						MarkdownDescription: "LogFormat for VMAlert to be configured with. default or json",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("default", "json"),
						},
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel for VMAlert to be configured with.",
						MarkdownDescription: "LogLevel for VMAlert to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
						},
					},

					"min_ready_seconds": schema.Int64Attribute{
						Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
						MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
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

					"notifier": schema.SingleNestedAttribute{
						Description:         "Notifier prometheus alertmanager endpoint spec. Required at least one of notifier or notifiers when there are alerting rules. e.g. http://127.0.0.1:9093 If specified both notifier and notifiers, notifier will be added as last element to notifiers. only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						MarkdownDescription: "Notifier prometheus alertmanager endpoint spec. Required at least one of notifier or notifiers when there are alerting rules. e.g. http://127.0.0.1:9093 If specified both notifier and notifiers, notifier will be added as last element to notifiers. only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
								Attributes: map[string]schema.Attribute{
									"password": schema.SingleNestedAttribute{
										Description:         "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"password_file": schema.StringAttribute{
										Description:         "PasswordFile defines path to password file at disk must be pre-mounted",
										MarkdownDescription: "PasswordFile defines path to password file at disk must be pre-mounted",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"bearer_token_file": schema.StringAttribute{
								Description:         "Path to bearer token file",
								MarkdownDescription: "Path to bearer token file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bearer_token_secret": schema.SingleNestedAttribute{
								Description:         "Optional bearer auth token to use for -remoteWrite.url",
								MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from. Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"headers": schema.ListAttribute{
								Description:         "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								MarkdownDescription: "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2": schema.MapAttribute{
								Description:         "OAuth2 defines OAuth2 configuration",
								MarkdownDescription: "OAuth2 defines OAuth2 configuration",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector allows service discovery for alertmanager in this case all matched vmalertmanager replicas will be added into vmalert notifier.url as statefulset pod.fqdn",
								MarkdownDescription: "Selector allows service discovery for alertmanager in this case all matched vmalertmanager replicas will be added into vmalert notifier.url as statefulset pod.fqdn",
								Attributes: map[string]schema.Attribute{
									"label_selector": schema.SingleNestedAttribute{
										Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
										MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
										Description:         "NamespaceSelector is a selector for selecting either all namespaces or a list of namespaces.",
										MarkdownDescription: "NamespaceSelector is a selector for selecting either all namespaces or a list of namespaces.",
										Attributes: map[string]schema.Attribute{
											"any": schema.BoolAttribute{
												Description:         "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
												MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_names": schema.ListAttribute{
												Description:         "List of namespace names.",
												MarkdownDescription: "List of namespace names.",
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

							"tls_config": schema.MapAttribute{
								Description:         "TLSConfig specifies TLSConfig configuration parameters.",
								MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "AlertManager url. E.g. http://127.0.0.1:9093",
								MarkdownDescription: "AlertManager url. E.g. http://127.0.0.1:9093",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"notifier_config_ref": schema.SingleNestedAttribute{
						Description:         "NotifierConfigRef reference for secret with notifier configuration for vmalert only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						MarkdownDescription: "NotifierConfigRef reference for secret with notifier configuration for vmalert only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from. Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

					"notifiers": schema.ListNestedAttribute{
						Description:         "Notifiers prometheus alertmanager endpoints. Required at least one of notifier or notifiers when there are alerting rules. e.g. http://127.0.0.1:9093 If specified both notifier and notifiers, notifier will be added as last element to notifiers. only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						MarkdownDescription: "Notifiers prometheus alertmanager endpoints. Required at least one of notifier or notifiers when there are alerting rules. e.g. http://127.0.0.1:9093 If specified both notifier and notifiers, notifier will be added as last element to notifiers. only one of notifier options could be chosen: notifierConfigRef or notifiers + notifier",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"basic_auth": schema.SingleNestedAttribute{
									Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
									MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
											MarkdownDescription: "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from. Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

										"password_file": schema.StringAttribute{
											Description:         "PasswordFile defines path to password file at disk must be pre-mounted",
											MarkdownDescription: "PasswordFile defines path to password file at disk must be pre-mounted",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
											MarkdownDescription: "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from. Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

								"bearer_token_file": schema.StringAttribute{
									Description:         "Path to bearer token file",
									MarkdownDescription: "Path to bearer token file",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"bearer_token_secret": schema.SingleNestedAttribute{
									Description:         "Optional bearer auth token to use for -remoteWrite.url",
									MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from. Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

								"headers": schema.ListAttribute{
									Description:         "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
									MarkdownDescription: "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"oauth2": schema.MapAttribute{
									Description:         "OAuth2 defines OAuth2 configuration",
									MarkdownDescription: "OAuth2 defines OAuth2 configuration",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"selector": schema.SingleNestedAttribute{
									Description:         "Selector allows service discovery for alertmanager in this case all matched vmalertmanager replicas will be added into vmalert notifier.url as statefulset pod.fqdn",
									MarkdownDescription: "Selector allows service discovery for alertmanager in this case all matched vmalertmanager replicas will be added into vmalert notifier.url as statefulset pod.fqdn",
									Attributes: map[string]schema.Attribute{
										"label_selector": schema.SingleNestedAttribute{
											Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
											MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
											Description:         "NamespaceSelector is a selector for selecting either all namespaces or a list of namespaces.",
											MarkdownDescription: "NamespaceSelector is a selector for selecting either all namespaces or a list of namespaces.",
											Attributes: map[string]schema.Attribute{
												"any": schema.BoolAttribute{
													Description:         "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
													MarkdownDescription: "Boolean describing whether all namespaces are selected in contrast to a list restricting them.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"match_names": schema.ListAttribute{
													Description:         "List of namespace names.",
													MarkdownDescription: "List of namespace names.",
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

								"tls_config": schema.MapAttribute{
									Description:         "TLSConfig specifies TLSConfig configuration parameters.",
									MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "AlertManager url. E.g. http://127.0.0.1:9093",
									MarkdownDescription: "AlertManager url. E.g. http://127.0.0.1:9093",
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

					"paused": schema.BoolAttribute{
						Description:         "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
						MarkdownDescription: "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_disruption_budget": schema.SingleNestedAttribute{
						Description:         "PodDisruptionBudget created by operator",
						MarkdownDescription: "PodDisruptionBudget created by operator",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.StringAttribute{
								Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_available": schema.StringAttribute{
								Description:         "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
								MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector_labels": schema.MapAttribute{
								Description:         "replaces default labels selector generated by operator it's useful when you need to create custom budget",
								MarkdownDescription: "replaces default labels selector generated by operator it's useful when you need to create custom budget",
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
						Description:         "PodMetadata configures Labels and Annotations which are propagated to the VMAlert pods.",
						MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VMAlert pods.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": schema.StringAttribute{
						Description:         "Port listen address",
						MarkdownDescription: "Port listen address",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName class assigned to the Pods",
						MarkdownDescription: "PriorityClassName class assigned to the Pods",
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

					"remote_read": schema.SingleNestedAttribute{
						Description:         "RemoteRead Optional URL to read vmalert state (persisted via RemoteWrite) This configuration only makes sense if alerts state has been successfully persisted (via RemoteWrite) before. see -remoteRead.url docs in vmalerts for details. E.g. http://127.0.0.1:8428",
						MarkdownDescription: "RemoteRead Optional URL to read vmalert state (persisted via RemoteWrite) This configuration only makes sense if alerts state has been successfully persisted (via RemoteWrite) before. see -remoteRead.url docs in vmalerts for details. E.g. http://127.0.0.1:8428",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
								Attributes: map[string]schema.Attribute{
									"password": schema.SingleNestedAttribute{
										Description:         "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"password_file": schema.StringAttribute{
										Description:         "PasswordFile defines path to password file at disk must be pre-mounted",
										MarkdownDescription: "PasswordFile defines path to password file at disk must be pre-mounted",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"bearer_token_file": schema.StringAttribute{
								Description:         "Path to bearer token file",
								MarkdownDescription: "Path to bearer token file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bearer_token_secret": schema.SingleNestedAttribute{
								Description:         "Optional bearer auth token to use for -remoteWrite.url",
								MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from. Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"headers": schema.ListAttribute{
								Description:         "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								MarkdownDescription: "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"lookback": schema.StringAttribute{
								Description:         "Lookback defines how far to look into past for alerts timeseries. For example, if lookback=1h then range from now() to now()-1h will be scanned. (default 1h0m0s) Applied only to RemoteReadSpec",
								MarkdownDescription: "Lookback defines how far to look into past for alerts timeseries. For example, if lookback=1h then range from now() to now()-1h will be scanned. (default 1h0m0s) Applied only to RemoteReadSpec",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2": schema.MapAttribute{
								Description:         "OAuth2 defines OAuth2 configuration",
								MarkdownDescription: "OAuth2 defines OAuth2 configuration",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_config": schema.MapAttribute{
								Description:         "TLSConfig specifies TLSConfig configuration parameters.",
								MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "URL of the endpoint to send samples to.",
								MarkdownDescription: "URL of the endpoint to send samples to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_write": schema.SingleNestedAttribute{
						Description:         "RemoteWrite Optional URL to remote-write compatible storage to persist vmalert state and rule results to. Rule results will be persisted according to each rule. Alerts state will be persisted in the form of time series named ALERTS and ALERTS_FOR_STATE see -remoteWrite.url docs in vmalerts for details. E.g. http://127.0.0.1:8428",
						MarkdownDescription: "RemoteWrite Optional URL to remote-write compatible storage to persist vmalert state and rule results to. Rule results will be persisted according to each rule. Alerts state will be persisted in the form of time series named ALERTS and ALERTS_FOR_STATE see -remoteWrite.url docs in vmalerts for details. E.g. http://127.0.0.1:8428",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
								Attributes: map[string]schema.Attribute{
									"password": schema.SingleNestedAttribute{
										Description:         "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Password defines reference for secret with password value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

									"password_file": schema.StringAttribute{
										Description:         "PasswordFile defines path to password file at disk must be pre-mounted",
										MarkdownDescription: "PasswordFile defines path to password file at disk must be pre-mounted",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										MarkdownDescription: "Username defines reference for secret with username value The secret needs to be in the same namespace as scrape object",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
												MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"bearer_token_file": schema.StringAttribute{
								Description:         "Path to bearer token file",
								MarkdownDescription: "Path to bearer token file",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bearer_token_secret": schema.SingleNestedAttribute{
								Description:         "Optional bearer auth token to use for -remoteWrite.url",
								MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from. Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

							"concurrency": schema.Int64Attribute{
								Description:         "Defines number of readers that concurrently write into remote storage (default 1)",
								MarkdownDescription: "Defines number of readers that concurrently write into remote storage (default 1)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flush_interval": schema.StringAttribute{
								Description:         "Defines interval of flushes to remote write endpoint (default 5s)",
								MarkdownDescription: "Defines interval of flushes to remote write endpoint (default 5s)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
								},
							},

							"headers": schema.ListAttribute{
								Description:         "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								MarkdownDescription: "Headers allow configuring custom http headers Must be in form of semicolon separated header with value e.g. headerName:headerValue vmalert supports it since 1.79.0 version",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_batch_size": schema.Int64Attribute{
								Description:         "Defines defines max number of timeseries to be flushed at once (default 1000)",
								MarkdownDescription: "Defines defines max number of timeseries to be flushed at once (default 1000)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_queue_size": schema.Int64Attribute{
								Description:         "Defines the max number of pending datapoints to remote write endpoint (default 100000)",
								MarkdownDescription: "Defines the max number of pending datapoints to remote write endpoint (default 100000)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"oauth2": schema.MapAttribute{
								Description:         "OAuth2 defines OAuth2 configuration",
								MarkdownDescription: "OAuth2 defines OAuth2 configuration",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_config": schema.MapAttribute{
								Description:         "TLSConfig specifies TLSConfig configuration parameters.",
								MarkdownDescription: "TLSConfig specifies TLSConfig configuration parameters.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "URL of the endpoint to send samples to.",
								MarkdownDescription: "URL of the endpoint to send samples to.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replica_count": schema.Int64Attribute{
						Description:         "ReplicaCount is the expected size of the Application.",
						MarkdownDescription: "ReplicaCount is the expected size of the Application.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
						MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
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

					"revision_history_limit_count": schema.Int64Attribute{
						Description:         "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
						MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rolling_update": schema.SingleNestedAttribute{
						Description:         "RollingUpdate - overrides deployment update params.",
						MarkdownDescription: "RollingUpdate - overrides deployment update params.",
						Attributes: map[string]schema.Attribute{
							"max_surge": schema.StringAttribute{
								Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
								MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_unavailable": schema.StringAttribute{
								Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
								MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rule_namespace_selector": schema.SingleNestedAttribute{
						Description:         "RuleNamespaceSelector to be selected for VMRules discovery. Works in combination with Selector. If both nil - behaviour controlled by selectAllByDefault NamespaceSelector nil - only objects at VMAlert namespace.",
						MarkdownDescription: "RuleNamespaceSelector to be selected for VMRules discovery. Works in combination with Selector. If both nil - behaviour controlled by selectAllByDefault NamespaceSelector nil - only objects at VMAlert namespace.",
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

					"rule_path": schema.ListAttribute{
						Description:         "RulePath to the file with alert rules. Supports patterns. Flag can be specified multiple times. Examples: -rule /path/to/file. Path to a single file with alerting rules -rule dir/*.yaml -rule /*.yaml. Relative path to all .yaml files in folder, absolute path to all .yaml files in root. by default operator adds /etc/vmalert/configs/base/vmalert.yaml",
						MarkdownDescription: "RulePath to the file with alert rules. Supports patterns. Flag can be specified multiple times. Examples: -rule /path/to/file. Path to a single file with alerting rules -rule dir/*.yaml -rule /*.yaml. Relative path to all .yaml files in folder, absolute path to all .yaml files in root. by default operator adds /etc/vmalert/configs/base/vmalert.yaml",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rule_selector": schema.SingleNestedAttribute{
						Description:         "RuleSelector selector to select which VMRules to mount for loading alerting rules from. Works in combination with NamespaceSelector. If both nil - behaviour controlled by selectAllByDefault NamespaceSelector nil - only objects at VMAlert namespace.",
						MarkdownDescription: "RuleSelector selector to select which VMRules to mount for loading alerting rules from. Works in combination with NamespaceSelector. If both nil - behaviour controlled by selectAllByDefault NamespaceSelector nil - only objects at VMAlert namespace.",
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

					"runtime_class_name": schema.StringAttribute{
						Description:         "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
						MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
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
						Description:         "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
						MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_context": schema.MapAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"select_all_by_default": schema.BoolAttribute{
						Description:         "SelectAllByDefault changes default behavior for empty CRD selectors, such RuleSelector. with selectAllByDefault: true and empty serviceScrapeSelector and RuleNamespaceSelector Operator selects all exist serviceScrapes with selectAllByDefault: false - selects nothing",
						MarkdownDescription: "SelectAllByDefault changes default behavior for empty CRD selectors, such RuleSelector. with selectAllByDefault: true and empty serviceScrapeSelector and RuleNamespaceSelector Operator selects all exist serviceScrapes with selectAllByDefault: false - selects nothing",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run the pods",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run the pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_scrape_spec": schema.MapAttribute{
						Description:         "ServiceScrapeSpec that will be added to vmalert VMServiceScrape spec",
						MarkdownDescription: "ServiceScrapeSpec that will be added to vmalert VMServiceScrape spec",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_spec": schema.SingleNestedAttribute{
						Description:         "ServiceSpec that will be added to vmalert service spec",
						MarkdownDescription: "ServiceSpec that will be added to vmalert service spec",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
								MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
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
								Description:         "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
								MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"use_as_default": schema.BoolAttribute{
								Description:         "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
								MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
								Required:            false,
								Optional:            true,
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

					"topology_spread_constraints": schema.ListAttribute{
						Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"update_strategy": schema.StringAttribute{
						Description:         "UpdateStrategy - overrides default update strategy.",
						MarkdownDescription: "UpdateStrategy - overrides default update strategy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Recreate", "RollingUpdate"),
						},
					},

					"use_default_resources": schema.BoolAttribute{
						Description:         "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
						MarkdownDescription: "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_vm_config_reloader": schema.BoolAttribute{
						Description:         "UseVMConfigReloader replaces prometheus-like config-reloader with vm one. It uses secrets watch instead of file watch which greatly increases speed of config updates",
						MarkdownDescription: "UseVMConfigReloader replaces prometheus-like config-reloader with vm one. It uses secrets watch instead of file watch which greatly increases speed of config updates",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
						MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
									MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"mount_propagation": schema.StringAttribute{
									Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
									MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
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
									Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
									MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"recursive_read_only": schema.StringAttribute{
									Description:         "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
									MarkdownDescription: "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
									Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
									MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path_expr": schema.StringAttribute{
									Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
									MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						Description:         "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
						MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
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

func (r *OperatorVictoriametricsComVmalertV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_alert_v1beta1_manifest")

	var model OperatorVictoriametricsComVmalertV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMAlert")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
