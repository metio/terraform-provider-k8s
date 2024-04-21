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
	_ datasource.DataSource = &OperatorVictoriametricsComVmagentV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmagentV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmagentV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmagentV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmagentV1Beta1ManifestData struct {
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
		APIServerConfig *struct {
			Authorization *struct {
				Credentials *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
				Type            *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
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
			BearerToken     *string `tfsdk:"bearer_token" json:"bearerToken,omitempty"`
			BearerTokenFile *string `tfsdk:"bearer_token_file" json:"bearerTokenFile,omitempty"`
			Host            *string `tfsdk:"host" json:"host,omitempty"`
			TlsConfig       *struct {
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
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
		} `tfsdk:"a_pi_server_config" json:"aPIServerConfig,omitempty"`
		AdditionalScrapeConfigs *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"additional_scrape_configs" json:"additionalScrapeConfigs,omitempty"`
		Affinity                    *map[string]string `tfsdk:"affinity" json:"affinity,omitempty"`
		ArbitraryFSAccessThroughSMs *struct {
			Deny *bool `tfsdk:"deny" json:"deny,omitempty"`
		} `tfsdk:"arbitrary_fs_access_through_s_ms" json:"arbitraryFSAccessThroughSMs,omitempty"`
		ClaimTemplates *[]struct {
			ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec       *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSource  *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source" json:"dataSource,omitempty"`
				DataSourceRef *struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			Status *struct {
				AccessModes        *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
				AllocatedResources *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
				Capacity           *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
				Conditions         *[]struct {
					LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
					LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
					Message            *string `tfsdk:"message" json:"message,omitempty"`
					Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
					Status             *string `tfsdk:"status" json:"status,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
				ResizeStatus *string `tfsdk:"resize_status" json:"resizeStatus,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"claim_templates" json:"claimTemplates,omitempty"`
		ConfigMaps              *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
		ConfigReloaderExtraArgs *map[string]string   `tfsdk:"config_reloader_extra_args" json:"configReloaderExtraArgs,omitempty"`
		Containers              *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
		DnsConfig               *struct {
			Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
			Options     *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
		} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
		DnsPolicy              *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
		EnforcedNamespaceLabel *string              `tfsdk:"enforced_namespace_label" json:"enforcedNamespaceLabel,omitempty"`
		ExternalLabels         *map[string]string   `tfsdk:"external_labels" json:"externalLabels,omitempty"`
		ExtraArgs              *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
		ExtraEnvs              *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
		HostNetwork            *bool                `tfsdk:"host_network" json:"hostNetwork,omitempty"`
		Host_aliases           *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
		IgnoreNamespaceSelectors *bool `tfsdk:"ignore_namespace_selectors" json:"ignoreNamespaceSelectors,omitempty"`
		Image                    *struct {
			PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		IngestOnlyMode      *bool                `tfsdk:"ingest_only_mode" json:"ingestOnlyMode,omitempty"`
		InitContainers      *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
		InlineRelabelConfig *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"inline_relabel_config" json:"inlineRelabelConfig,omitempty"`
		InlineScrapeConfig *string `tfsdk:"inline_scrape_config" json:"inlineScrapeConfig,omitempty"`
		InsertPorts        *struct {
			GraphitePort     *string `tfsdk:"graphite_port" json:"graphitePort,omitempty"`
			InfluxPort       *string `tfsdk:"influx_port" json:"influxPort,omitempty"`
			OpenTSDBHTTPPort *string `tfsdk:"open_tsdbhttp_port" json:"openTSDBHTTPPort,omitempty"`
			OpenTSDBPort     *string `tfsdk:"open_tsdb_port" json:"openTSDBPort,omitempty"`
		} `tfsdk:"insert_ports" json:"insertPorts,omitempty"`
		License *struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			KeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_ref" json:"keyRef,omitempty"`
		} `tfsdk:"license" json:"license,omitempty"`
		LivenessProbe               *map[string]string `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		LogFormat                   *string            `tfsdk:"log_format" json:"logFormat,omitempty"`
		LogLevel                    *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		MaxScrapeInterval           *string            `tfsdk:"max_scrape_interval" json:"maxScrapeInterval,omitempty"`
		MinReadySeconds             *int64             `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
		MinScrapeInterval           *string            `tfsdk:"min_scrape_interval" json:"minScrapeInterval,omitempty"`
		NodeScrapeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_scrape_namespace_selector" json:"nodeScrapeNamespaceSelector,omitempty"`
		NodeScrapeRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"node_scrape_relabel_template" json:"nodeScrapeRelabelTemplate,omitempty"`
		NodeScrapeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_scrape_selector" json:"nodeScrapeSelector,omitempty"`
		NodeSelector            *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		OverrideHonorLabels     *bool              `tfsdk:"override_honor_labels" json:"overrideHonorLabels,omitempty"`
		OverrideHonorTimestamps *bool              `tfsdk:"override_honor_timestamps" json:"overrideHonorTimestamps,omitempty"`
		PodDisruptionBudget     *struct {
			MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
			SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
		} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		PodScrapeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"pod_scrape_namespace_selector" json:"podScrapeNamespaceSelector,omitempty"`
		PodScrapeRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"pod_scrape_relabel_template" json:"podScrapeRelabelTemplate,omitempty"`
		PodScrapeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"pod_scrape_selector" json:"podScrapeSelector,omitempty"`
		Port                   *string `tfsdk:"port" json:"port,omitempty"`
		PriorityClassName      *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ProbeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"probe_namespace_selector" json:"probeNamespaceSelector,omitempty"`
		ProbeScrapeRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"probe_scrape_relabel_template" json:"probeScrapeRelabelTemplate,omitempty"`
		ProbeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"probe_selector" json:"probeSelector,omitempty"`
		ReadinessGates *[]struct {
			ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
		} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
		ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		RelabelConfig  *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"relabel_config" json:"relabelConfig,omitempty"`
		RemoteWrite *[]struct {
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
			BearerTokenSecret *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
			Headers                *[]string `tfsdk:"headers" json:"headers,omitempty"`
			InlineUrlRelabelConfig *[]struct {
				Action       *string            `tfsdk:"action" json:"action,omitempty"`
				If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
				Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Match        *string            `tfsdk:"match" json:"match,omitempty"`
				Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"inline_url_relabel_config" json:"inlineUrlRelabelConfig,omitempty"`
			Oauth2 *struct {
				Client_id *struct {
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
				} `tfsdk:"client_id" json:"client_id,omitempty"`
				Client_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"client_secret" json:"client_secret,omitempty"`
				Client_secret_file *string            `tfsdk:"client_secret_file" json:"client_secret_file,omitempty"`
				Endpoint_params    *map[string]string `tfsdk:"endpoint_params" json:"endpoint_params,omitempty"`
				Scopes             *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
				Token_url          *string            `tfsdk:"token_url" json:"token_url,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			SendTimeout      *string `tfsdk:"send_timeout" json:"sendTimeout,omitempty"`
			StreamAggrConfig *struct {
				DedupInterval *string `tfsdk:"dedup_interval" json:"dedupInterval,omitempty"`
				DropInput     *bool   `tfsdk:"drop_input" json:"dropInput,omitempty"`
				KeepInput     *bool   `tfsdk:"keep_input" json:"keepInput,omitempty"`
				Rules         *[]struct {
					By                    *[]string `tfsdk:"by" json:"by,omitempty"`
					Flush_on_shutdown     *bool     `tfsdk:"flush_on_shutdown" json:"flush_on_shutdown,omitempty"`
					Input_relabel_configs *[]struct {
						Action       *string            `tfsdk:"action" json:"action,omitempty"`
						If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
						Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Match        *string            `tfsdk:"match" json:"match,omitempty"`
						Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"input_relabel_configs" json:"input_relabel_configs,omitempty"`
					Interval               *string            `tfsdk:"interval" json:"interval,omitempty"`
					Match                  *map[string]string `tfsdk:"match" json:"match,omitempty"`
					Output_relabel_configs *[]struct {
						Action       *string            `tfsdk:"action" json:"action,omitempty"`
						If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
						Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Match        *string            `tfsdk:"match" json:"match,omitempty"`
						Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"output_relabel_configs" json:"output_relabel_configs,omitempty"`
					Outputs            *[]string `tfsdk:"outputs" json:"outputs,omitempty"`
					Staleness_interval *string   `tfsdk:"staleness_interval" json:"staleness_interval,omitempty"`
					Without            *[]string `tfsdk:"without" json:"without,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
			} `tfsdk:"stream_aggr_config" json:"streamAggrConfig,omitempty"`
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
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url              *string `tfsdk:"url" json:"url,omitempty"`
			UrlRelabelConfig *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"url_relabel_config" json:"urlRelabelConfig,omitempty"`
		} `tfsdk:"remote_write" json:"remoteWrite,omitempty"`
		RemoteWriteSettings *struct {
			FlushInterval      *string            `tfsdk:"flush_interval" json:"flushInterval,omitempty"`
			Label              *map[string]string `tfsdk:"label" json:"label,omitempty"`
			MaxBlockSize       *int64             `tfsdk:"max_block_size" json:"maxBlockSize,omitempty"`
			MaxDiskUsagePerURL *int64             `tfsdk:"max_disk_usage_per_url" json:"maxDiskUsagePerURL,omitempty"`
			Queues             *int64             `tfsdk:"queues" json:"queues,omitempty"`
			ShowURL            *bool              `tfsdk:"show_url" json:"showURL,omitempty"`
			TmpDataPath        *string            `tfsdk:"tmp_data_path" json:"tmpDataPath,omitempty"`
			UseMultiTenantMode *bool              `tfsdk:"use_multi_tenant_mode" json:"useMultiTenantMode,omitempty"`
		} `tfsdk:"remote_write_settings" json:"remoteWriteSettings,omitempty"`
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
		RuntimeClassName              *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
		SchedulerName                 *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		ScrapeConfigNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"scrape_config_namespace_selector" json:"scrapeConfigNamespaceSelector,omitempty"`
		ScrapeConfigRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"scrape_config_relabel_template" json:"scrapeConfigRelabelTemplate,omitempty"`
		ScrapeConfigSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"scrape_config_selector" json:"scrapeConfigSelector,omitempty"`
		ScrapeInterval                 *string            `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
		ScrapeTimeout                  *string            `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
		Secrets                        *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
		SecurityContext                *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
		SelectAllByDefault             *bool              `tfsdk:"select_all_by_default" json:"selectAllByDefault,omitempty"`
		ServiceAccountName             *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		ServiceScrapeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"service_scrape_namespace_selector" json:"serviceScrapeNamespaceSelector,omitempty"`
		ServiceScrapeRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"service_scrape_relabel_template" json:"serviceScrapeRelabelTemplate,omitempty"`
		ServiceScrapeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"service_scrape_selector" json:"serviceScrapeSelector,omitempty"`
		ServiceScrapeSpec *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
		ServiceSpec       *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
		} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
		ShardCount                    *int64             `tfsdk:"shard_count" json:"shardCount,omitempty"`
		StartupProbe                  *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
		StatefulMode                  *bool              `tfsdk:"stateful_mode" json:"statefulMode,omitempty"`
		StatefulRollingUpdateStrategy *string            `tfsdk:"stateful_rolling_update_strategy" json:"statefulRollingUpdateStrategy,omitempty"`
		StatefulStorage               *struct {
			DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" json:"disableMountSubPath,omitempty"`
			EmptyDir            *struct {
				Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
				SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
			} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
			VolumeClaimTemplate *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
				Metadata   *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					DataSource  *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"data_source" json:"dataSource,omitempty"`
					DataSourceRef *struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
					Resources *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
				Status *struct {
					AccessModes        *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
					AllocatedResources *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
					Capacity           *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
					Conditions         *[]struct {
						LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
						LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
						Message            *string `tfsdk:"message" json:"message,omitempty"`
						Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
						Status             *string `tfsdk:"status" json:"status,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"conditions" json:"conditions,omitempty"`
					Phase        *string `tfsdk:"phase" json:"phase,omitempty"`
					ResizeStatus *string `tfsdk:"resize_status" json:"resizeStatus,omitempty"`
				} `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
		} `tfsdk:"stateful_storage" json:"statefulStorage,omitempty"`
		StaticScrapeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"static_scrape_namespace_selector" json:"staticScrapeNamespaceSelector,omitempty"`
		StaticScrapeRelabelTemplate *[]struct {
			Action       *string            `tfsdk:"action" json:"action,omitempty"`
			If           *map[string]string `tfsdk:"if" json:"if,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Match        *string            `tfsdk:"match" json:"match,omitempty"`
			Modulus      *int64             `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *map[string]string `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string            `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string            `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string          `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string            `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"static_scrape_relabel_template" json:"staticScrapeRelabelTemplate,omitempty"`
		StaticScrapeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"static_scrape_selector" json:"staticScrapeSelector,omitempty"`
		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
		Tolerations                   *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		UpdateStrategy            *string              `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		VmAgentExternalLabelName  *string              `tfsdk:"vm_agent_external_label_name" json:"vmAgentExternalLabelName,omitempty"`
		VolumeMounts              *[]struct {
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

func (r *OperatorVictoriametricsComVmagentV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_agent_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmagentV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMAgent - is a tiny but brave agent, which helps you collect metrics from various sources and stores them in VictoriaMetricsor any other Prometheus-compatible storage system that supports the remote_write protocol.",
		MarkdownDescription: "VMAgent - is a tiny but brave agent, which helps you collect metrics from various sources and stores them in VictoriaMetricsor any other Prometheus-compatible storage system that supports the remote_write protocol.",
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
				Description:         "VMAgentSpec defines the desired state of VMAgent",
				MarkdownDescription: "VMAgentSpec defines the desired state of VMAgent",
				Attributes: map[string]schema.Attribute{
					"a_pi_server_config": schema.SingleNestedAttribute{
						Description:         "APIServerConfig allows specifying a host and auth methods to access apiserver.If left empty, VMAgent is assumed to run inside of the clusterand will discover API servers automatically and use the pod's CA certificateand bearer token file at /var/run/secrets/kubernetes.io/serviceaccount/.",
						MarkdownDescription: "APIServerConfig allows specifying a host and auth methods to access apiserver.If left empty, VMAgent is assumed to run inside of the clusterand will discover API servers automatically and use the pod's CA certificateand bearer token file at /var/run/secrets/kubernetes.io/serviceaccount/.",
						Attributes: map[string]schema.Attribute{
							"authorization": schema.SingleNestedAttribute{
								Description:         "Authorization configures generic authorization params",
								MarkdownDescription: "Authorization configures generic authorization params",
								Attributes: map[string]schema.Attribute{
									"credentials": schema.SingleNestedAttribute{
										Description:         "Reference to the secret with value for authorization",
										MarkdownDescription: "Reference to the secret with value for authorization",
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

									"credentials_file": schema.StringAttribute{
										Description:         "File with value for authorization",
										MarkdownDescription: "File with value for authorization",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type of authorization, default to bearer",
										MarkdownDescription: "Type of authorization, default to bearer",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
								Attributes: map[string]schema.Attribute{
									"password": schema.SingleNestedAttribute{
										Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
										MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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

									"password_file": schema.StringAttribute{
										Description:         "PasswordFile defines path to password file at disk",
										MarkdownDescription: "PasswordFile defines path to password file at disk",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
										MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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

							"bearer_token": schema.StringAttribute{
								Description:         "Bearer token for accessing apiserver.",
								MarkdownDescription: "Bearer token for accessing apiserver.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bearer_token_file": schema.StringAttribute{
								Description:         "File to read bearer token for accessing apiserver.",
								MarkdownDescription: "File to read bearer token for accessing apiserver.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "Host of apiserver.A valid string consisting of a hostname or IP followed by an optional port number",
								MarkdownDescription: "Host of apiserver.A valid string consisting of a hostname or IP followed by an optional port number",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tls_config": schema.SingleNestedAttribute{
								Description:         "TLSConfig Config to use for accessing apiserver.",
								MarkdownDescription: "TLSConfig Config to use for accessing apiserver.",
								Attributes: map[string]schema.Attribute{
									"ca": schema.SingleNestedAttribute{
										Description:         "Stuct containing the CA cert to use for the targets.",
										MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

									"ca_file": schema.StringAttribute{
										Description:         "Path to the CA cert in the container to use for the targets.",
										MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cert": schema.SingleNestedAttribute{
										Description:         "Struct containing the client cert file for the targets.",
										MarkdownDescription: "Struct containing the client cert file for the targets.",
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
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

									"cert_file": schema.StringAttribute{
										Description:         "Path to the client cert file in the container for the targets.",
										MarkdownDescription: "Path to the client cert file in the container for the targets.",
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
										Description:         "Path to the client key file in the container for the targets.",
										MarkdownDescription: "Path to the client key file in the container for the targets.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key_secret": schema.SingleNestedAttribute{
										Description:         "Secret containing the client key file for the targets.",
										MarkdownDescription: "Secret containing the client key file for the targets.",
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

					"additional_scrape_configs": schema.SingleNestedAttribute{
						Description:         "AdditionalScrapeConfigs As scrape configs are appended, the user is responsible to make sure itis valid. Note that using this feature may expose the possibility tobreak upgrades of VMAgent. It is advised to review VMAgent releasenotes to ensure that no incompatible scrape configs are going to breakVMAgent after the upgrade.",
						MarkdownDescription: "AdditionalScrapeConfigs As scrape configs are appended, the user is responsible to make sure itis valid. Note that using this feature may expose the possibility tobreak upgrades of VMAgent. It is advised to review VMAgent releasenotes to ensure that no incompatible scrape configs are going to breakVMAgent after the upgrade.",
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

					"affinity": schema.MapAttribute{
						Description:         "Affinity If specified, the pod's scheduling constraints.",
						MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"arbitrary_fs_access_through_s_ms": schema.SingleNestedAttribute{
						Description:         "ArbitraryFSAccessThroughSMs configures whether configurationbased on a service scrape can access arbitrary files on the file systemof the VMAgent container e.g. bearer token files.",
						MarkdownDescription: "ArbitraryFSAccessThroughSMs configures whether configurationbased on a service scrape can access arbitrary files on the file systemof the VMAgent container e.g. bearer token files.",
						Attributes: map[string]schema.Attribute{
							"deny": schema.BoolAttribute{
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

					"claim_templates": schema.ListNestedAttribute{
						Description:         "ClaimTemplates allows adding additional VolumeClaimTemplates for VMAgent in StatefulMode",
						MarkdownDescription: "ClaimTemplates allows adding additional VolumeClaimTemplates for VMAgent in StatefulMode",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
									MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metadata": schema.MapAttribute{
									Description:         "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									MarkdownDescription: "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"data_source": schema.SingleNestedAttribute{
											Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
											MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
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

										"data_source_ref": schema.SingleNestedAttribute{
											Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
											MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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

												"namespace": schema.StringAttribute{
													Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
											MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
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

										"selector": schema.SingleNestedAttribute{
											Description:         "selector is a label query over volumes to consider for binding.",
											MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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

										"storage_class_name": schema.StringAttribute{
											Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
											MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_mode": schema.StringAttribute{
											Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
											MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_name": schema.StringAttribute{
											Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
											MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"status": schema.SingleNestedAttribute{
									Description:         "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allocated_resources": schema.MapAttribute{
											Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"capacity": schema.MapAttribute{
											Description:         "capacity represents the actual resources of the underlying volume.",
											MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"conditions": schema.ListNestedAttribute{
											Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
											MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"last_probe_time": schema.StringAttribute{
														Description:         "lastProbeTime is the time we probed the condition.",
														MarkdownDescription: "lastProbeTime is the time we probed the condition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"last_transition_time": schema.StringAttribute{
														Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
														MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"message": schema.StringAttribute{
														Description:         "message is the human-readable message indicating details about last transition.",
														MarkdownDescription: "message is the human-readable message indicating details about last transition.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"reason": schema.StringAttribute{
														Description:         "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
														MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"status": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
														MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

										"phase": schema.StringAttribute{
											Description:         "phase represents the current phase of PersistentVolumeClaim.",
											MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resize_status": schema.StringAttribute{
											Description:         "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
											MarkdownDescription: "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
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

					"config_maps": schema.ListAttribute{
						Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the vmagentobject, which shall be mounted into the vmagent Pods.will be mounted at path  /etc/vm/configs",
						MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the vmagentobject, which shall be mounted into the vmagent Pods.will be mounted at path  /etc/vm/configs",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config_reloader_extra_args": schema.MapAttribute{
						Description:         "ConfigReloaderExtraArgs that will be passed to  VMAuths config-reloader containerfor example resyncInterval: '30s'",
						MarkdownDescription: "ConfigReloaderExtraArgs that will be passed to  VMAuths config-reloader containerfor example resyncInterval: '30s'",
						ElementType:         types.StringType,
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
						Description:         "DNSPolicy set DNS policy for the pod",
						MarkdownDescription: "DNSPolicy set DNS policy for the pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enforced_namespace_label": schema.StringAttribute{
						Description:         "EnforcedNamespaceLabel enforces adding a namespace label of origin for each alertand metric that is user created. The label value will always be the namespace of the object that isbeing created.",
						MarkdownDescription: "EnforcedNamespaceLabel enforces adding a namespace label of origin for each alertand metric that is user created. The label value will always be the namespace of the object that isbeing created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_labels": schema.MapAttribute{
						Description:         "ExternalLabels The labels to add to any time series scraped by vmagent.it doesn't affect metrics ingested directly by push API's",
						MarkdownDescription: "ExternalLabels The labels to add to any time series scraped by vmagent.it doesn't affect metrics ingested directly by push API's",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_args": schema.MapAttribute{
						Description:         "ExtraArgs that will be passed to  VMAgent podfor example remoteWrite.tmpDataPath: /tmpit would be converted to flag --remoteWrite.tmpDataPath=/tmp",
						MarkdownDescription: "ExtraArgs that will be passed to  VMAgent podfor example remoteWrite.tmpDataPath: /tmpit would be converted to flag --remoteWrite.tmpDataPath=/tmp",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_envs": schema.ListAttribute{
						Description:         "ExtraEnvs that will be added to VMAgent pod",
						MarkdownDescription: "ExtraEnvs that will be added to VMAgent pod",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_network": schema.BoolAttribute{
						Description:         "HostNetwork controls whether the pod may use the node network namespace",
						MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "HostAliases provides mapping between ip and hostnames,that would be propagated to pod,cannot be used with HostNetwork.",
						MarkdownDescription: "HostAliases provides mapping between ip and hostnames,that would be propagated to pod,cannot be used with HostNetwork.",
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

					"ignore_namespace_selectors": schema.BoolAttribute{
						Description:         "IgnoreNamespaceSelectors if set to true will ignore NamespaceSelector settings fromthe podscrape and vmservicescrape configs, and they will only discover endpointswithin their current namespace.  Defaults to false.",
						MarkdownDescription: "IgnoreNamespaceSelectors if set to true will ignore NamespaceSelector settings fromthe podscrape and vmservicescrape configs, and they will only discover endpointswithin their current namespace.  Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image - docker image settings for VMAgentif no specified operator uses default config version",
						MarkdownDescription: "Image - docker image settings for VMAgentif no specified operator uses default config version",
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

					"ingest_only_mode": schema.BoolAttribute{
						Description:         "IngestOnlyMode switches vmagent into unmanaged modeit disables any config generation for scrapingCurrently it prevents vmagent from managing tls and auth options for remote write",
						MarkdownDescription: "IngestOnlyMode switches vmagent into unmanaged modeit disables any config generation for scrapingCurrently it prevents vmagent from managing tls and auth options for remote write",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_containers": schema.ListAttribute{
						Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmagent configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g.fetch secrets for injection into the vmagent configuration from external sources. Anyerrors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/Using initContainers for any use case other then secret fetching is entirely outside the scopeof what the maintainers will support and by doing so, you accept that this behaviour may breakat any time without notice.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"inline_relabel_config": schema.ListNestedAttribute{
						Description:         "InlineRelabelConfig - defines GlobalRelabelConfig for vmagent, can be defined directly at CRD.",
						MarkdownDescription: "InlineRelabelConfig - defines GlobalRelabelConfig for vmagent, can be defined directly at CRD.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"inline_scrape_config": schema.StringAttribute{
						Description:         "InlineScrapeConfig As scrape configs are appended, the user is responsible to make sure itis valid. Note that using this feature may expose the possibility tobreak upgrades of VMAgent. It is advised to review VMAgent releasenotes to ensure that no incompatible scrape configs are going to breakVMAgent after the upgrade.it should be defined as single yaml file.inlineScrapeConfig: |    - job_name: 'prometheus'      static_configs:      - targets: ['localhost:9090']",
						MarkdownDescription: "InlineScrapeConfig As scrape configs are appended, the user is responsible to make sure itis valid. Note that using this feature may expose the possibility tobreak upgrades of VMAgent. It is advised to review VMAgent releasenotes to ensure that no incompatible scrape configs are going to breakVMAgent after the upgrade.it should be defined as single yaml file.inlineScrapeConfig: |    - job_name: 'prometheus'      static_configs:      - targets: ['localhost:9090']",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"insert_ports": schema.SingleNestedAttribute{
						Description:         "InsertPorts - additional listen ports for data ingestion.",
						MarkdownDescription: "InsertPorts - additional listen ports for data ingestion.",
						Attributes: map[string]schema.Attribute{
							"graphite_port": schema.StringAttribute{
								Description:         "GraphitePort listen port",
								MarkdownDescription: "GraphitePort listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"influx_port": schema.StringAttribute{
								Description:         "InfluxPort listen port",
								MarkdownDescription: "InfluxPort listen port",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_tsdbhttp_port": schema.StringAttribute{
								Description:         "OpenTSDBHTTPPort for http connections.",
								MarkdownDescription: "OpenTSDBHTTPPort for http connections.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_tsdb_port": schema.StringAttribute{
								Description:         "OpenTSDBPort for tcp and udp listen",
								MarkdownDescription: "OpenTSDBPort for tcp and udp listen",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
						Description:         "LogFormat for VMAgent to be configured with.",
						MarkdownDescription: "LogFormat for VMAgent to be configured with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("default", "json"),
						},
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel for VMAgent to be configured with.INFO, WARN, ERROR, FATAL, PANIC",
						MarkdownDescription: "LogLevel for VMAgent to be configured with.INFO, WARN, ERROR, FATAL, PANIC",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
						},
					},

					"max_scrape_interval": schema.StringAttribute{
						Description:         "MaxScrapeInterval allows limiting maximum scrape interval for VMServiceScrape, VMPodScrape and other scrapesIf interval is higher than defined limit, 'maxScrapeInterval' will be used.",
						MarkdownDescription: "MaxScrapeInterval allows limiting maximum scrape interval for VMServiceScrape, VMPodScrape and other scrapesIf interval is higher than defined limit, 'maxScrapeInterval' will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_ready_seconds": schema.Int64Attribute{
						Description:         "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
						MarkdownDescription: "MinReadySeconds defines a minim number os seconds to wait before starting update next podif previous in healthy state",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_scrape_interval": schema.StringAttribute{
						Description:         "MinScrapeInterval allows limiting minimal scrape interval for VMServiceScrape, VMPodScrape and other scrapesIf interval is lower than defined limit, 'minScrapeInterval' will be used.",
						MarkdownDescription: "MinScrapeInterval allows limiting minimal scrape interval for VMServiceScrape, VMPodScrape and other scrapesIf interval is lower than defined limit, 'minScrapeInterval' will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_scrape_namespace_selector": schema.SingleNestedAttribute{
						Description:         "NodeScrapeNamespaceSelector defines Namespaces to be selected for VMNodeScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "NodeScrapeNamespaceSelector defines Namespaces to be selected for VMNodeScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"node_scrape_relabel_template": schema.ListNestedAttribute{
						Description:         "NodeScrapeRelabelTemplate defines relabel config, that will be added to each VMNodeScrape.it's useful for adding specific labels to all targets",
						MarkdownDescription: "NodeScrapeRelabelTemplate defines relabel config, that will be added to each VMNodeScrape.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"node_scrape_selector": schema.SingleNestedAttribute{
						Description:         "NodeScrapeSelector defines VMNodeScrape to be selected for scraping.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "NodeScrapeSelector defines VMNodeScrape to be selected for scraping.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
						MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"override_honor_labels": schema.BoolAttribute{
						Description:         "OverrideHonorLabels if set to true overrides all user configured honor_labels.If HonorLabels is set in ServiceScrape or PodScrape to true, this overrides honor_labels to false.",
						MarkdownDescription: "OverrideHonorLabels if set to true overrides all user configured honor_labels.If HonorLabels is set in ServiceScrape or PodScrape to true, this overrides honor_labels to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"override_honor_timestamps": schema.BoolAttribute{
						Description:         "OverrideHonorTimestamps allows to globally enforce honoring timestamps in all scrape configs.",
						MarkdownDescription: "OverrideHonorTimestamps allows to globally enforce honoring timestamps in all scrape configs.",
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
						Description:         "PodMetadata configures Labels and Annotations which are propagated to the vmagent pods.",
						MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the vmagent pods.",
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

					"pod_scrape_namespace_selector": schema.SingleNestedAttribute{
						Description:         "PodScrapeNamespaceSelector defines Namespaces to be selected for VMPodScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "PodScrapeNamespaceSelector defines Namespaces to be selected for VMPodScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"pod_scrape_relabel_template": schema.ListNestedAttribute{
						Description:         "PodScrapeRelabelTemplate defines relabel config, that will be added to each VMPodScrape.it's useful for adding specific labels to all targets",
						MarkdownDescription: "PodScrapeRelabelTemplate defines relabel config, that will be added to each VMPodScrape.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"pod_scrape_selector": schema.SingleNestedAttribute{
						Description:         "PodScrapeSelector defines PodScrapes to be selected for target discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "PodScrapeSelector defines PodScrapes to be selected for target discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"port": schema.StringAttribute{
						Description:         "Port listen address",
						MarkdownDescription: "Port listen address",
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

					"probe_namespace_selector": schema.SingleNestedAttribute{
						Description:         "ProbeNamespaceSelector defines Namespaces to be selected for VMProbe discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "ProbeNamespaceSelector defines Namespaces to be selected for VMProbe discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"probe_scrape_relabel_template": schema.ListNestedAttribute{
						Description:         "ProbeScrapeRelabelTemplate defines relabel config, that will be added to each VMProbeScrape.it's useful for adding specific labels to all targets",
						MarkdownDescription: "ProbeScrapeRelabelTemplate defines relabel config, that will be added to each VMProbeScrape.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"probe_selector": schema.SingleNestedAttribute{
						Description:         "ProbeSelector defines VMProbe to be selected for target probing.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "ProbeSelector defines VMProbe to be selected for target probing.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"relabel_config": schema.SingleNestedAttribute{
						Description:         "RelabelConfig ConfigMap with global relabel config -remoteWrite.relabelConfigThis relabeling is applied to all the collected metrics before sending them to remote storage.",
						MarkdownDescription: "RelabelConfig ConfigMap with global relabel config -remoteWrite.relabelConfigThis relabeling is applied to all the collected metrics before sending them to remote storage.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key to select.",
								MarkdownDescription: "The key to select.",
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

					"remote_write": schema.ListNestedAttribute{
						Description:         "RemoteWrite list of victoria metrics /some other remote write systemfor vm it must looks like: http://victoria-metrics-single:8429/api/v1/writeor for cluster different urlhttps://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/app/vmagent#splitting-data-streams-among-multiple-systems",
						MarkdownDescription: "RemoteWrite list of victoria metrics /some other remote write systemfor vm it must looks like: http://victoria-metrics-single:8429/api/v1/writeor for cluster different urlhttps://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/app/vmagent#splitting-data-streams-among-multiple-systems",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"basic_auth": schema.SingleNestedAttribute{
									Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
									MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
											MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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

										"password_file": schema.StringAttribute{
											Description:         "PasswordFile defines path to password file at disk",
											MarkdownDescription: "PasswordFile defines path to password file at disk",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
											MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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

								"bearer_token_secret": schema.SingleNestedAttribute{
									Description:         "Optional bearer auth token to use for -remoteWrite.url",
									MarkdownDescription: "Optional bearer auth token to use for -remoteWrite.url",
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

								"headers": schema.ListAttribute{
									Description:         "Headers allow configuring custom http headersMust be in form of semicolon separated header with valuee.g.headerName: headerValuevmagent supports since 1.79.0 version",
									MarkdownDescription: "Headers allow configuring custom http headersMust be in form of semicolon separated header with valuee.g.headerName: headerValuevmagent supports since 1.79.0 version",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"inline_url_relabel_config": schema.ListNestedAttribute{
									Description:         "InlineUrlRelabelConfig defines relabeling config for remoteWriteURL, it can be defined at crd spec.",
									MarkdownDescription: "InlineUrlRelabelConfig defines relabeling config for remoteWriteURL, it can be defined at crd spec.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "Action to perform based on regex matching. Default is 'replace'",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"if": schema.MapAttribute{
												Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels is used together with Match for 'action: graphite'",
												MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match": schema.StringAttribute{
												Description:         "Match is used together with Labels for 'action: graphite'",
												MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"modulus": schema.Int64Attribute{
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"regex": schema.MapAttribute{
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replacement": schema.StringAttribute{
												Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"separator": schema.StringAttribute{
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source_labels": schema.ListAttribute{
												Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_label": schema.StringAttribute{
												Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

								"oauth2": schema.SingleNestedAttribute{
									Description:         "OAuth2 defines auth configuration",
									MarkdownDescription: "OAuth2 defines auth configuration",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.SingleNestedAttribute{
											Description:         "The secret or configmap containing the OAuth2 client id",
											MarkdownDescription: "The secret or configmap containing the OAuth2 client id",
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
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "The secret containing the OAuth2 client secret",
											MarkdownDescription: "The secret containing the OAuth2 client secret",
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

										"client_secret_file": schema.StringAttribute{
											Description:         "ClientSecretFile defines path for client secret file.",
											MarkdownDescription: "ClientSecretFile defines path for client secret file.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"endpoint_params": schema.MapAttribute{
											Description:         "Parameters to append to the token URL",
											MarkdownDescription: "Parameters to append to the token URL",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scopes": schema.ListAttribute{
											Description:         "OAuth2 scopes used for the token request",
											MarkdownDescription: "OAuth2 scopes used for the token request",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"token_url": schema.StringAttribute{
											Description:         "The URL to fetch the token from",
											MarkdownDescription: "The URL to fetch the token from",
											Required:            true,
											Optional:            false,
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

								"send_timeout": schema.StringAttribute{
									Description:         "Timeout for sending a single block of data to -remoteWrite.url (default 1m0s)",
									MarkdownDescription: "Timeout for sending a single block of data to -remoteWrite.url (default 1m0s)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
									},
								},

								"stream_aggr_config": schema.SingleNestedAttribute{
									Description:         "StreamAggrConfig defines stream aggregation configuration for VMAgent for -remoteWrite.url",
									MarkdownDescription: "StreamAggrConfig defines stream aggregation configuration for VMAgent for -remoteWrite.url",
									Attributes: map[string]schema.Attribute{
										"dedup_interval": schema.StringAttribute{
											Description:         "Allows setting different de-duplication intervals per each configured remote storage",
											MarkdownDescription: "Allows setting different de-duplication intervals per each configured remote storage",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"drop_input": schema.BoolAttribute{
											Description:         "Allow drop all the input samples after the aggregation",
											MarkdownDescription: "Allow drop all the input samples after the aggregation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keep_input": schema.BoolAttribute{
											Description:         "Allows writing both raw and aggregate data",
											MarkdownDescription: "Allows writing both raw and aggregate data",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rules": schema.ListNestedAttribute{
											Description:         "Stream aggregation rules",
											MarkdownDescription: "Stream aggregation rules",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"by": schema.ListAttribute{
														Description:         "By is an optional list of labels for grouping input series.See also Without.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
														MarkdownDescription: "By is an optional list of labels for grouping input series.See also Without.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"flush_on_shutdown": schema.BoolAttribute{
														Description:         "FlushOnShutdown defines whether to flush the aggregation state on process terminationor config reload. Is 'false' by default.It is not recommended changing this setting, unless unfinished aggregations statesare preferred to missing data points.",
														MarkdownDescription: "FlushOnShutdown defines whether to flush the aggregation state on process terminationor config reload. Is 'false' by default.It is not recommended changing this setting, unless unfinished aggregations statesare preferred to missing data points.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"input_relabel_configs": schema.ListNestedAttribute{
														Description:         "InputRelabelConfigs is an optional relabeling rules, which are applied on the inputbefore aggregation.",
														MarkdownDescription: "InputRelabelConfigs is an optional relabeling rules, which are applied on the inputbefore aggregation.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Action to perform based on regex matching. Default is 'replace'",
																	MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"if": schema.MapAttribute{
																	Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
																	MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"labels": schema.MapAttribute{
																	Description:         "Labels is used together with Match for 'action: graphite'",
																	MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"match": schema.StringAttribute{
																	Description:         "Match is used together with Labels for 'action: graphite'",
																	MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"modulus": schema.Int64Attribute{
																	Description:         "Modulus to take of the hash of the source label values.",
																	MarkdownDescription: "Modulus to take of the hash of the source label values.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"regex": schema.MapAttribute{
																	Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
																	MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"replacement": schema.StringAttribute{
																	Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
																	MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"separator": schema.StringAttribute{
																	Description:         "Separator placed between concatenated source label values. default is ';'.",
																	MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"source_labels": schema.ListAttribute{
																	Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
																	MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_label": schema.StringAttribute{
																	Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
																	MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

													"interval": schema.StringAttribute{
														Description:         "Interval is the interval between aggregations.",
														MarkdownDescription: "Interval is the interval between aggregations.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"match": schema.MapAttribute{
														Description:         "Match is a label selector (or list of label selectors) for filtering time series for the given selector.If the match isn't set, then all the input time series are processed.",
														MarkdownDescription: "Match is a label selector (or list of label selectors) for filtering time series for the given selector.If the match isn't set, then all the input time series are processed.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"output_relabel_configs": schema.ListNestedAttribute{
														Description:         "OutputRelabelConfigs is an optional relabeling rules, which are appliedon the aggregated output before being sent to remote storage.",
														MarkdownDescription: "OutputRelabelConfigs is an optional relabeling rules, which are appliedon the aggregated output before being sent to remote storage.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "Action to perform based on regex matching. Default is 'replace'",
																	MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"if": schema.MapAttribute{
																	Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
																	MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"labels": schema.MapAttribute{
																	Description:         "Labels is used together with Match for 'action: graphite'",
																	MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"match": schema.StringAttribute{
																	Description:         "Match is used together with Labels for 'action: graphite'",
																	MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"modulus": schema.Int64Attribute{
																	Description:         "Modulus to take of the hash of the source label values.",
																	MarkdownDescription: "Modulus to take of the hash of the source label values.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"regex": schema.MapAttribute{
																	Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
																	MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"replacement": schema.StringAttribute{
																	Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
																	MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"separator": schema.StringAttribute{
																	Description:         "Separator placed between concatenated source label values. default is ';'.",
																	MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"source_labels": schema.ListAttribute{
																	Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
																	MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_label": schema.StringAttribute{
																	Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
																	MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

													"outputs": schema.ListAttribute{
														Description:         "Outputs is a list of output aggregate functions to produce.The following names are allowed:- total - aggregates input counters- increase - counts the increase over input counters- count_series - counts the input series- count_samples - counts the input samples- sum_samples - sums the input samples- last - the last biggest sample value- min - the minimum sample value- max - the maximum sample value- avg - the average value across all the samples- stddev - standard deviation across all the samples- stdvar - standard variance across all the samples- histogram_bucket - creates VictoriaMetrics histogram for input samples- quantiles(phi1, ..., phiN) - quantiles' estimation for phi in the range [0..1]The output time series will have the following names:  input_name:aggr_<interval>_<output>",
														MarkdownDescription: "Outputs is a list of output aggregate functions to produce.The following names are allowed:- total - aggregates input counters- increase - counts the increase over input counters- count_series - counts the input series- count_samples - counts the input samples- sum_samples - sums the input samples- last - the last biggest sample value- min - the minimum sample value- max - the maximum sample value- avg - the average value across all the samples- stddev - standard deviation across all the samples- stdvar - standard variance across all the samples- histogram_bucket - creates VictoriaMetrics histogram for input samples- quantiles(phi1, ..., phiN) - quantiles' estimation for phi in the range [0..1]The output time series will have the following names:  input_name:aggr_<interval>_<output>",
														ElementType:         types.StringType,
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"staleness_interval": schema.StringAttribute{
														Description:         "StalenessInterval defines an interval after which the series state will be reset if no samples have been sent during it.",
														MarkdownDescription: "StalenessInterval defines an interval after which the series state will be reset if no samples have been sent during it.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"without": schema.ListAttribute{
														Description:         "Without is an optional list of labels, which must be excluded when grouping input series.See also By.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
														MarkdownDescription: "Without is an optional list of labels, which must be excluded when grouping input series.See also By.If neither By nor Without are set, then the Outputs are calculatedindividually per each input time series.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
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

								"tls_config": schema.SingleNestedAttribute{
									Description:         "TLSConfig describes tls configuration for remote write target",
									MarkdownDescription: "TLSConfig describes tls configuration for remote write target",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "Stuct containing the CA cert to use for the targets.",
											MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

										"ca_file": schema.StringAttribute{
											Description:         "Path to the CA cert in the container to use for the targets.",
											MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cert": schema.SingleNestedAttribute{
											Description:         "Struct containing the client cert file for the targets.",
											MarkdownDescription: "Struct containing the client cert file for the targets.",
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
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

										"cert_file": schema.StringAttribute{
											Description:         "Path to the client cert file in the container for the targets.",
											MarkdownDescription: "Path to the client cert file in the container for the targets.",
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
											Description:         "Path to the client key file in the container for the targets.",
											MarkdownDescription: "Path to the client key file in the container for the targets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_secret": schema.SingleNestedAttribute{
											Description:         "Secret containing the client key file for the targets.",
											MarkdownDescription: "Secret containing the client key file for the targets.",
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

								"url": schema.StringAttribute{
									Description:         "URL of the endpoint to send samples to.",
									MarkdownDescription: "URL of the endpoint to send samples to.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"url_relabel_config": schema.SingleNestedAttribute{
									Description:         "ConfigMap with relabeling config which is applied to metrics before sending them to the corresponding -remoteWrite.url",
									MarkdownDescription: "ConfigMap with relabeling config which is applied to metrics before sending them to the corresponding -remoteWrite.url",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key to select.",
											MarkdownDescription: "The key to select.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_write_settings": schema.SingleNestedAttribute{
						Description:         "RemoteWriteSettings defines global settings for all remoteWrite urls.",
						MarkdownDescription: "RemoteWriteSettings defines global settings for all remoteWrite urls.",
						Attributes: map[string]schema.Attribute{
							"flush_interval": schema.StringAttribute{
								Description:         "Interval for flushing the data to remote storage. (default 1s)",
								MarkdownDescription: "Interval for flushing the data to remote storage. (default 1s)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
								},
							},

							"label": schema.MapAttribute{
								Description:         "Labels in the form 'name=value' to add to all the metrics before sending them. This overrides the label if it already exists.",
								MarkdownDescription: "Labels in the form 'name=value' to add to all the metrics before sending them. This overrides the label if it already exists.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_block_size": schema.Int64Attribute{
								Description:         "The maximum size in bytes of unpacked request to send to remote storage",
								MarkdownDescription: "The maximum size in bytes of unpacked request to send to remote storage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_disk_usage_per_url": schema.Int64Attribute{
								Description:         "The maximum file-based buffer size in bytes at -remoteWrite.tmpDataPath",
								MarkdownDescription: "The maximum file-based buffer size in bytes at -remoteWrite.tmpDataPath",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"queues": schema.Int64Attribute{
								Description:         "The number of concurrent queues",
								MarkdownDescription: "The number of concurrent queues",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"show_url": schema.BoolAttribute{
								Description:         "Whether to show -remoteWrite.url in the exported metrics. It is hidden by default, since it can contain sensitive auth info",
								MarkdownDescription: "Whether to show -remoteWrite.url in the exported metrics. It is hidden by default, since it can contain sensitive auth info",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tmp_data_path": schema.StringAttribute{
								Description:         "Path to directory where temporary data for remote write component is stored (default vmagent-remotewrite-data)",
								MarkdownDescription: "Path to directory where temporary data for remote write component is stored (default vmagent-remotewrite-data)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_multi_tenant_mode": schema.BoolAttribute{
								Description:         "Configures vmagent in multi-tenant mode with direct cluster supportdocs https://docs.victoriametrics.com/vmagent.html#multitenancyit's global setting and affects all remote storage configurations",
								MarkdownDescription: "Configures vmagent in multi-tenant mode with direct cluster supportdocs https://docs.victoriametrics.com/vmagent.html#multitenancyit's global setting and affects all remote storage configurations",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replica_count": schema.Int64Attribute{
						Description:         "ReplicaCount is the expected size of the VMAgent cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.NOTE enable VMSingle deduplication for replica usage",
						MarkdownDescription: "ReplicaCount is the expected size of the VMAgent cluster. The controller willeventually make the size of the running cluster equal to the expectedsize.NOTE enable VMSingle deduplication for replica usage",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not specified - default setting will be used",
						MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/if not specified - default setting will be used",
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

					"rolling_update": schema.SingleNestedAttribute{
						Description:         "RollingUpdate - overrides deployment update params.",
						MarkdownDescription: "RollingUpdate - overrides deployment update params.",
						Attributes: map[string]schema.Attribute{
							"max_surge": schema.StringAttribute{
								Description:         "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
								MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_unavailable": schema.StringAttribute{
								Description:         "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
								MarkdownDescription: "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
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

					"scrape_config_namespace_selector": schema.SingleNestedAttribute{
						Description:         "ScrapeConfigNamespaceSelector defines Namespaces to be selected for VMScrapeConfig discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "ScrapeConfigNamespaceSelector defines Namespaces to be selected for VMScrapeConfig discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"scrape_config_relabel_template": schema.ListNestedAttribute{
						Description:         "ScrapeConfigRelabelTemplate defines relabel config, that will be added to each VMScrapeConfig.it's useful for adding specific labels to all targets",
						MarkdownDescription: "ScrapeConfigRelabelTemplate defines relabel config, that will be added to each VMScrapeConfig.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"scrape_config_selector": schema.SingleNestedAttribute{
						Description:         "ScrapeConfigSelector defines VMScrapeConfig to be selected for target discovery.Works in combination with NamespaceSelector.",
						MarkdownDescription: "ScrapeConfigSelector defines VMScrapeConfig to be selected for target discovery.Works in combination with NamespaceSelector.",
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

					"scrape_interval": schema.StringAttribute{
						Description:         "ScrapeInterval defines how often scrape targets by default",
						MarkdownDescription: "ScrapeInterval defines how often scrape targets by default",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
						},
					},

					"scrape_timeout": schema.StringAttribute{
						Description:         "ScrapeTimeout defines global timeout for targets scrape",
						MarkdownDescription: "ScrapeTimeout defines global timeout for targets scrape",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
						},
					},

					"secrets": schema.ListAttribute{
						Description:         "Secrets is a list of Secrets in the same namespace as the vmagentobject, which shall be mounted into the vmagent Pods.will be mounted at path /etc/vm/secrets",
						MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the vmagentobject, which shall be mounted into the vmagent Pods.will be mounted at path /etc/vm/secrets",
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
						Description:         "SelectAllByDefault changes default behavior for empty CRD selectors, such ServiceScrapeSelector.with selectAllByDefault: true and empty serviceScrapeSelector and ServiceScrapeNamespaceSelectorOperator selects all exist serviceScrapeswith selectAllByDefault: false - selects nothing",
						MarkdownDescription: "SelectAllByDefault changes default behavior for empty CRD selectors, such ServiceScrapeSelector.with selectAllByDefault: true and empty serviceScrapeSelector and ServiceScrapeNamespaceSelectorOperator selects all exist serviceScrapeswith selectAllByDefault: false - selects nothing",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run theVMAgent Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run theVMAgent Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_scrape_namespace_selector": schema.SingleNestedAttribute{
						Description:         "ServiceScrapeNamespaceSelector Namespaces to be selected for VMServiceScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "ServiceScrapeNamespaceSelector Namespaces to be selected for VMServiceScrape discovery.Works in combination with Selector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"service_scrape_relabel_template": schema.ListNestedAttribute{
						Description:         "ServiceScrapeRelabelTemplate defines relabel config, that will be added to each VMServiceScrape.it's useful for adding specific labels to all targets",
						MarkdownDescription: "ServiceScrapeRelabelTemplate defines relabel config, that will be added to each VMServiceScrape.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"service_scrape_selector": schema.SingleNestedAttribute{
						Description:         "ServiceScrapeSelector defines ServiceScrapes to be selected for target discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "ServiceScrapeSelector defines ServiceScrapes to be selected for target discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"service_scrape_spec": schema.MapAttribute{
						Description:         "ServiceScrapeSpec that will be added to vmagent VMServiceScrape spec",
						MarkdownDescription: "ServiceScrapeSpec that will be added to vmagent VMServiceScrape spec",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_spec": schema.SingleNestedAttribute{
						Description:         "ServiceSpec that will be added to vmagent service spec",
						MarkdownDescription: "ServiceSpec that will be added to vmagent service spec",
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

							"use_as_default": schema.BoolAttribute{
								Description:         "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
								MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object ServiceChaning from headless service to clusterIP or loadbalancer may break cross-component communication",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"shard_count": schema.Int64Attribute{
						Description:         "ShardCount - numbers of shards of VMAgentin this case operator will use 1 deployment/sts per shard withreplicas count according to spec.replicas,see https://docs.victoriametrics.com/vmagent.html#scraping-big-number-of-targets",
						MarkdownDescription: "ShardCount - numbers of shards of VMAgentin this case operator will use 1 deployment/sts per shard withreplicas count according to spec.replicas,see https://docs.victoriametrics.com/vmagent.html#scraping-big-number-of-targets",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"startup_probe": schema.MapAttribute{
						Description:         "StartupProbe that will be added to CRD pod",
						MarkdownDescription: "StartupProbe that will be added to CRD pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stateful_mode": schema.BoolAttribute{
						Description:         "StatefulMode enables StatefulSet for 'VMAgent' instead of Deploymentit allows using persistent storage for vmagent's persistentQueue",
						MarkdownDescription: "StatefulMode enables StatefulSet for 'VMAgent' instead of Deploymentit allows using persistent storage for vmagent's persistentQueue",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stateful_rolling_update_strategy": schema.StringAttribute{
						Description:         "StatefulRollingUpdateStrategy allows configuration for strategyTypeset it to RollingUpdate for disabling operator statefulSet rollingUpdate",
						MarkdownDescription: "StatefulRollingUpdateStrategy allows configuration for strategyTypeset it to RollingUpdate for disabling operator statefulSet rollingUpdate",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stateful_storage": schema.SingleNestedAttribute{
						Description:         "StatefulStorage configures storage for StatefulSet",
						MarkdownDescription: "StatefulStorage configures storage for StatefulSet",
						Attributes: map[string]schema.Attribute{
							"disable_mount_sub_path": schema.BoolAttribute{
								Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
								MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary.DisableMountSubPath allows to remove any subPath usage in volume mounts.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"empty_dir": schema.SingleNestedAttribute{
								Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
								MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. Moreinfo: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
								Attributes: map[string]schema.Attribute{
									"medium": schema.StringAttribute{
										Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"size_limit": schema.StringAttribute{
										Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_claim_template": schema.SingleNestedAttribute{
								Description:         "A PVC spec to be used by the VMAlertManager StatefulSets.",
								MarkdownDescription: "A PVC spec to be used by the VMAlertManager StatefulSets.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",
										MarkdownDescription: "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",
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

									"spec": schema.SingleNestedAttribute{
										Description:         "Spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "Spec defines the desired characteristics of a volume requested by a pod author.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										Attributes: map[string]schema.Attribute{
											"access_modes": schema.ListAttribute{
												Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"data_source": schema.SingleNestedAttribute{
												Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
												MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
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

											"data_source_ref": schema.SingleNestedAttribute{
												Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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

													"namespace": schema.StringAttribute{
														Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
														MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
												MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
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

											"selector": schema.SingleNestedAttribute{
												Description:         "selector is a label query over volumes to consider for binding.",
												MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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

											"storage_class_name": schema.StringAttribute{
												Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
												MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_mode": schema.StringAttribute{
												Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
												MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_name": schema.StringAttribute{
												Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
												MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": schema.SingleNestedAttribute{
										Description:         "Status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "Status represents the current information/status of a persistent volume claim.Read-only.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										Attributes: map[string]schema.Attribute{
											"access_modes": schema.ListAttribute{
												Description:         "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allocated_resources": schema.MapAttribute{
												Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
												MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It maybe larger than the actual capacity when a volume expansion operation is requested.For storage quota, the larger value from allocatedResources and PVC.spec.resources is used.If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation.If a volume expansion capacity request is lowered, allocatedResources is onlylowered if there are no expansion operations in progress and if the actual volume capacityis equal or lower than the requested capacity.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capacity": schema.MapAttribute{
												Description:         "capacity represents the actual resources of the underlying volume.",
												MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"conditions": schema.ListNestedAttribute{
												Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
												MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is beingresized then the Condition will be set to 'ResizeStarted'.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"last_probe_time": schema.StringAttribute{
															Description:         "lastProbeTime is the time we probed the condition.",
															MarkdownDescription: "lastProbeTime is the time we probed the condition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																validators.DateTime64Validator(),
															},
														},

														"last_transition_time": schema.StringAttribute{
															Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
															MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																validators.DateTime64Validator(),
															},
														},

														"message": schema.StringAttribute{
															Description:         "message is the human-readable message indicating details about last transition.",
															MarkdownDescription: "message is the human-readable message indicating details about last transition.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"reason": schema.StringAttribute{
															Description:         "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
															MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reasonfor condition's last transition. If it reports 'ResizeStarted' that means the underlyingpersistent volume is being resized.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
															MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

											"phase": schema.StringAttribute{
												Description:         "phase represents the current phase of PersistentVolumeClaim.",
												MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resize_status": schema.StringAttribute{
												Description:         "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
												MarkdownDescription: "resizeStatus stores status of resize operation.ResizeStatus is not set by default but when expansion is complete resizeStatus is set to emptystring by resize controller or kubelet.This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
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

					"static_scrape_namespace_selector": schema.SingleNestedAttribute{
						Description:         "StaticScrapeNamespaceSelector defines Namespaces to be selected for VMStaticScrape discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
						MarkdownDescription: "StaticScrapeNamespaceSelector defines Namespaces to be selected for VMStaticScrape discovery.Works in combination with NamespaceSelector.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.If both nil - behaviour controlled by selectAllByDefault",
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

					"static_scrape_relabel_template": schema.ListNestedAttribute{
						Description:         "StaticScrapeRelabelTemplate defines relabel config, that will be added to each VMStaticScrape.it's useful for adding specific labels to all targets",
						MarkdownDescription: "StaticScrapeRelabelTemplate defines relabel config, that will be added to each VMStaticScrape.it's useful for adding specific labels to all targets",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on regex matching. Default is 'replace'",
									MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"if": schema.MapAttribute{
									Description:         "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									MarkdownDescription: "If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels is used together with Match for 'action: graphite'",
									MarkdownDescription: "Labels is used together with Match for 'action: graphite'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.StringAttribute{
									Description:         "Match is used together with Labels for 'action: graphite'",
									MarkdownDescription: "Match is used together with Labels for 'action: graphite'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"regex": schema.MapAttribute{
									Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'victoriaMetrics supports multiline regex joined with |https://docs.victoriametrics.com/vmagent/#relabeling-enhancements",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									MarkdownDescription: "Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator placed between concatenated source label values. default is ';'.",
									MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.",
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

					"static_scrape_selector": schema.SingleNestedAttribute{
						Description:         "StaticScrapeSelector defines PodScrapes to be selected for target discovery.Works in combination with NamespaceSelector.If both nil - match everything.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.",
						MarkdownDescription: "StaticScrapeSelector defines PodScrapes to be selected for target discovery.Works in combination with NamespaceSelector.If both nil - match everything.NamespaceSelector nil - only objects at VMAgent namespace.Selector nil - only objects at NamespaceSelector namespaces.",
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

					"update_strategy": schema.StringAttribute{
						Description:         "UpdateStrategy - overrides default update strategy.works only for deployments, statefulset always use OnDelete.",
						MarkdownDescription: "UpdateStrategy - overrides default update strategy.works only for deployments, statefulset always use OnDelete.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Recreate", "RollingUpdate"),
						},
					},

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for componentit restricts disk writes accessuses non-root user out of the boxdrops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vm_agent_external_label_name": schema.StringAttribute{
						Description:         "VMAgentExternalLabelName Name of vmAgent external label used to denote vmAgent instancename. Defaults to the value of 'prometheus'. External label will_not_ be added when value is set to empty string ('''').",
						MarkdownDescription: "VMAgentExternalLabelName Name of vmAgent external label used to denote vmAgent instancename. Defaults to the value of 'prometheus'. External label will_not_ be added when value is set to empty string ('''').",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output deploy definition.VolumeMounts specified will be appended to other VolumeMounts in the vmagent container,that are generated as a result of StorageSpec objects.",
						MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output deploy definition.VolumeMounts specified will be appended to other VolumeMounts in the vmagent container,that are generated as a result of StorageSpec objects.",
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

func (r *OperatorVictoriametricsComVmagentV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_agent_v1beta1_manifest")

	var model OperatorVictoriametricsComVmagentV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMAgent")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
