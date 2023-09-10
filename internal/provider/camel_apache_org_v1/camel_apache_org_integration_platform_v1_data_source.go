/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CamelApacheOrgIntegrationPlatformV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgIntegrationPlatformV1DataSource{}
)

func NewCamelApacheOrgIntegrationPlatformV1DataSource() datasource.DataSource {
	return &CamelApacheOrgIntegrationPlatformV1DataSource{}
}

type CamelApacheOrgIntegrationPlatformV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgIntegrationPlatformV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Build *struct {
			PublishStrategyOptions  *map[string]string `tfsdk:"publish_strategy_options" json:"PublishStrategyOptions,omitempty"`
			BaseImage               *string            `tfsdk:"base_image" json:"baseImage,omitempty"`
			BuildCatalogToolTimeout *string            `tfsdk:"build_catalog_tool_timeout" json:"buildCatalogToolTimeout,omitempty"`
			BuildConfiguration      *struct {
				LimitCPU          *string `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory       *string `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				OperatorNamespace *string `tfsdk:"operator_namespace" json:"operatorNamespace,omitempty"`
				OrderStrategy     *string `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
				RequestCPU        *string `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory     *string `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				Strategy          *string `tfsdk:"strategy" json:"strategy,omitempty"`
				ToolImage         *string `tfsdk:"tool_image" json:"toolImage,omitempty"`
			} `tfsdk:"build_configuration" json:"buildConfiguration,omitempty"`
			Maven *struct {
				CaSecrets *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"ca_secrets" json:"caSecrets,omitempty"`
				CliOptions *[]string `tfsdk:"cli_options" json:"cliOptions,omitempty"`
				Extension  *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
					GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
					Version    *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"extension" json:"extension,omitempty"`
				LocalRepository *string `tfsdk:"local_repository" json:"localRepository,omitempty"`
				Profiles        *[]struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"profiles" json:"profiles,omitempty"`
				Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
				Settings   *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"settings" json:"settings,omitempty"`
				SettingsSecurity *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"settings_security" json:"settingsSecurity,omitempty"`
			} `tfsdk:"maven" json:"maven,omitempty"`
			MaxRunningBuilds *int64  `tfsdk:"max_running_builds" json:"maxRunningBuilds,omitempty"`
			PublishStrategy  *string `tfsdk:"publish_strategy" json:"publishStrategy,omitempty"`
			Registry         *struct {
				Address      *string `tfsdk:"address" json:"address,omitempty"`
				Ca           *string `tfsdk:"ca" json:"ca,omitempty"`
				Insecure     *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
				Organization *string `tfsdk:"organization" json:"organization,omitempty"`
				Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
			} `tfsdk:"registry" json:"registry,omitempty"`
			RuntimeProvider *string `tfsdk:"runtime_provider" json:"runtimeProvider,omitempty"`
			RuntimeVersion  *string `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
			Timeout         *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"build" json:"build,omitempty"`
		Cluster       *string `tfsdk:"cluster" json:"cluster,omitempty"`
		Configuration *[]struct {
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		Kamelet *struct {
			Repositories *[]struct {
				Uri *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"repositories" json:"repositories,omitempty"`
		} `tfsdk:"kamelet" json:"kamelet,omitempty"`
		Profile *string `tfsdk:"profile" json:"profile,omitempty"`
		Traits  *struct {
			Threescale *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"threescale" json:"3scale,omitempty"`
			Addons   *map[string]string `tfsdk:"addons" json:"addons,omitempty"`
			Affinity *struct {
				Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				NodeAffinityLabels    *[]string          `tfsdk:"node_affinity_labels" json:"nodeAffinityLabels,omitempty"`
				PodAffinity           *bool              `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
				PodAffinityLabels     *[]string          `tfsdk:"pod_affinity_labels" json:"podAffinityLabels,omitempty"`
				PodAntiAffinity       *bool              `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				PodAntiAffinityLabels *[]string          `tfsdk:"pod_anti_affinity_labels" json:"podAntiAffinityLabels,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
			Builder *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				LimitCPU      *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory   *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				MavenProfiles *[]string          `tfsdk:"maven_profiles" json:"mavenProfiles,omitempty"`
				OrderStrategy *string            `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
				Properties    *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RequestCPU    *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				Strategy      *string            `tfsdk:"strategy" json:"strategy,omitempty"`
				Tasks         *[]string          `tfsdk:"tasks" json:"tasks,omitempty"`
				Verbose       *bool              `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"builder" json:"builder,omitempty"`
			Camel *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Properties     *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RuntimeVersion *string            `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" json:"camel,omitempty"`
			Container *struct {
				Auto            *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration   *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled         *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Expose          *bool              `tfsdk:"expose" json:"expose,omitempty"`
				Image           *string            `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				LimitCPU        *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory     *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				Port            *int64             `tfsdk:"port" json:"port,omitempty"`
				PortName        *string            `tfsdk:"port_name" json:"portName,omitempty"`
				RequestCPU      *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory   *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				ServicePort     *int64             `tfsdk:"service_port" json:"servicePort,omitempty"`
				ServicePortName *string            `tfsdk:"service_port_name" json:"servicePortName,omitempty"`
			} `tfsdk:"container" json:"container,omitempty"`
			Cron *struct {
				ActiveDeadlineSeconds   *int64             `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
				Auto                    *bool              `tfsdk:"auto" json:"auto,omitempty"`
				BackoffLimit            *int64             `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
				Components              *string            `tfsdk:"components" json:"components,omitempty"`
				ConcurrencyPolicy       *string            `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
				Configuration           *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled                 *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Fallback                *bool              `tfsdk:"fallback" json:"fallback,omitempty"`
				Schedule                *string            `tfsdk:"schedule" json:"schedule,omitempty"`
				StartingDeadlineSeconds *int64             `tfsdk:"starting_deadline_seconds" json:"startingDeadlineSeconds,omitempty"`
			} `tfsdk:"cron" json:"cron,omitempty"`
			Dependencies *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"dependencies" json:"dependencies,omitempty"`
			Deployer *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Kind          *string            `tfsdk:"kind" json:"kind,omitempty"`
				UseSSA        *bool              `tfsdk:"use_ssa" json:"useSSA,omitempty"`
			} `tfsdk:"deployer" json:"deployer,omitempty"`
			Deployment *struct {
				Configuration               *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled                     *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				ProgressDeadlineSeconds     *int64             `tfsdk:"progress_deadline_seconds" json:"progressDeadlineSeconds,omitempty"`
				RollingUpdateMaxSurge       *int64             `tfsdk:"rolling_update_max_surge" json:"rollingUpdateMaxSurge,omitempty"`
				RollingUpdateMaxUnavailable *int64             `tfsdk:"rolling_update_max_unavailable" json:"rollingUpdateMaxUnavailable,omitempty"`
				Strategy                    *string            `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
			Environment *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				ContainerMeta *bool              `tfsdk:"container_meta" json:"containerMeta,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				HttpProxy     *bool              `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
				Vars          *[]string          `tfsdk:"vars" json:"vars,omitempty"`
			} `tfsdk:"environment" json:"environment,omitempty"`
			Error_handler *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Ref           *string            `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"error_handler" json:"error-handler,omitempty"`
			Gc *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				DiscoveryCache *string            `tfsdk:"discovery_cache" json:"discoveryCache,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"gc" json:"gc,omitempty"`
			Health *struct {
				Configuration             *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled                   *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				LivenessFailureThreshold  *int64             `tfsdk:"liveness_failure_threshold" json:"livenessFailureThreshold,omitempty"`
				LivenessInitialDelay      *int64             `tfsdk:"liveness_initial_delay" json:"livenessInitialDelay,omitempty"`
				LivenessPeriod            *int64             `tfsdk:"liveness_period" json:"livenessPeriod,omitempty"`
				LivenessProbeEnabled      *bool              `tfsdk:"liveness_probe_enabled" json:"livenessProbeEnabled,omitempty"`
				LivenessScheme            *string            `tfsdk:"liveness_scheme" json:"livenessScheme,omitempty"`
				LivenessSuccessThreshold  *int64             `tfsdk:"liveness_success_threshold" json:"livenessSuccessThreshold,omitempty"`
				LivenessTimeout           *int64             `tfsdk:"liveness_timeout" json:"livenessTimeout,omitempty"`
				ReadinessFailureThreshold *int64             `tfsdk:"readiness_failure_threshold" json:"readinessFailureThreshold,omitempty"`
				ReadinessInitialDelay     *int64             `tfsdk:"readiness_initial_delay" json:"readinessInitialDelay,omitempty"`
				ReadinessPeriod           *int64             `tfsdk:"readiness_period" json:"readinessPeriod,omitempty"`
				ReadinessProbeEnabled     *bool              `tfsdk:"readiness_probe_enabled" json:"readinessProbeEnabled,omitempty"`
				ReadinessScheme           *string            `tfsdk:"readiness_scheme" json:"readinessScheme,omitempty"`
				ReadinessSuccessThreshold *int64             `tfsdk:"readiness_success_threshold" json:"readinessSuccessThreshold,omitempty"`
				ReadinessTimeout          *int64             `tfsdk:"readiness_timeout" json:"readinessTimeout,omitempty"`
				StartupFailureThreshold   *int64             `tfsdk:"startup_failure_threshold" json:"startupFailureThreshold,omitempty"`
				StartupInitialDelay       *int64             `tfsdk:"startup_initial_delay" json:"startupInitialDelay,omitempty"`
				StartupPeriod             *int64             `tfsdk:"startup_period" json:"startupPeriod,omitempty"`
				StartupProbeEnabled       *bool              `tfsdk:"startup_probe_enabled" json:"startupProbeEnabled,omitempty"`
				StartupScheme             *string            `tfsdk:"startup_scheme" json:"startupScheme,omitempty"`
				StartupSuccessThreshold   *int64             `tfsdk:"startup_success_threshold" json:"startupSuccessThreshold,omitempty"`
				StartupTimeout            *int64             `tfsdk:"startup_timeout" json:"startupTimeout,omitempty"`
			} `tfsdk:"health" json:"health,omitempty"`
			Ingress *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Host          *string            `tfsdk:"host" json:"host,omitempty"`
				Path          *string            `tfsdk:"path" json:"path,omitempty"`
				PathType      *string            `tfsdk:"path_type" json:"pathType,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			Istio *struct {
				Allow         *string            `tfsdk:"allow" json:"allow,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Inject        *bool              `tfsdk:"inject" json:"inject,omitempty"`
			} `tfsdk:"istio" json:"istio,omitempty"`
			Jolokia *struct {
				CACert                     *string            `tfsdk:"ca_cert" json:"CACert,omitempty"`
				ClientPrincipal            *[]string          `tfsdk:"client_principal" json:"clientPrincipal,omitempty"`
				Configuration              *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				DiscoveryEnabled           *bool              `tfsdk:"discovery_enabled" json:"discoveryEnabled,omitempty"`
				Enabled                    *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				ExtendedClientCheck        *bool              `tfsdk:"extended_client_check" json:"extendedClientCheck,omitempty"`
				Host                       *string            `tfsdk:"host" json:"host,omitempty"`
				Options                    *[]string          `tfsdk:"options" json:"options,omitempty"`
				Password                   *string            `tfsdk:"password" json:"password,omitempty"`
				Port                       *int64             `tfsdk:"port" json:"port,omitempty"`
				Protocol                   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
				UseSSLClientAuthentication *bool              `tfsdk:"use_ssl_client_authentication" json:"useSSLClientAuthentication,omitempty"`
				User                       *string            `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"jolokia" json:"jolokia,omitempty"`
			Jvm *struct {
				Classpath     *string            `tfsdk:"classpath" json:"classpath,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Debug         *bool              `tfsdk:"debug" json:"debug,omitempty"`
				DebugAddress  *string            `tfsdk:"debug_address" json:"debugAddress,omitempty"`
				DebugSuspend  *bool              `tfsdk:"debug_suspend" json:"debugSuspend,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Options       *[]string          `tfsdk:"options" json:"options,omitempty"`
				PrintCommand  *bool              `tfsdk:"print_command" json:"printCommand,omitempty"`
			} `tfsdk:"jvm" json:"jvm,omitempty"`
			Kamelets *struct {
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				List          *string            `tfsdk:"list" json:"list,omitempty"`
			} `tfsdk:"kamelets" json:"kamelets,omitempty"`
			Keda *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"keda" json:"keda,omitempty"`
			Knative *struct {
				Auto                 *bool              `tfsdk:"auto" json:"auto,omitempty"`
				ChannelSinks         *[]string          `tfsdk:"channel_sinks" json:"channelSinks,omitempty"`
				ChannelSources       *[]string          `tfsdk:"channel_sources" json:"channelSources,omitempty"`
				Config               *string            `tfsdk:"config" json:"config,omitempty"`
				Configuration        *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled              *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				EndpointSinks        *[]string          `tfsdk:"endpoint_sinks" json:"endpointSinks,omitempty"`
				EndpointSources      *[]string          `tfsdk:"endpoint_sources" json:"endpointSources,omitempty"`
				EventSinks           *[]string          `tfsdk:"event_sinks" json:"eventSinks,omitempty"`
				EventSources         *[]string          `tfsdk:"event_sources" json:"eventSources,omitempty"`
				FilterSourceChannels *bool              `tfsdk:"filter_source_channels" json:"filterSourceChannels,omitempty"`
				NamespaceLabel       *bool              `tfsdk:"namespace_label" json:"namespaceLabel,omitempty"`
				SinkBinding          *bool              `tfsdk:"sink_binding" json:"sinkBinding,omitempty"`
			} `tfsdk:"knative" json:"knative,omitempty"`
			Knative_service *struct {
				Auto              *bool              `tfsdk:"auto" json:"auto,omitempty"`
				AutoscalingMetric *string            `tfsdk:"autoscaling_metric" json:"autoscalingMetric,omitempty"`
				AutoscalingTarget *int64             `tfsdk:"autoscaling_target" json:"autoscalingTarget,omitempty"`
				Class             *string            `tfsdk:"class" json:"class,omitempty"`
				Configuration     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				MaxScale          *int64             `tfsdk:"max_scale" json:"maxScale,omitempty"`
				MinScale          *int64             `tfsdk:"min_scale" json:"minScale,omitempty"`
				RolloutDuration   *string            `tfsdk:"rollout_duration" json:"rolloutDuration,omitempty"`
				Visibility        *string            `tfsdk:"visibility" json:"visibility,omitempty"`
			} `tfsdk:"knative_service" json:"knative-service,omitempty"`
			Logging *struct {
				Color           *bool              `tfsdk:"color" json:"color,omitempty"`
				Configuration   *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled         *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Format          *string            `tfsdk:"format" json:"format,omitempty"`
				Json            *bool              `tfsdk:"json" json:"json,omitempty"`
				JsonPrettyPrint *bool              `tfsdk:"json_pretty_print" json:"jsonPrettyPrint,omitempty"`
				Level           *string            `tfsdk:"level" json:"level,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			Master *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"master" json:"master,omitempty"`
			Mount *struct {
				Configs       *[]string          `tfsdk:"configs" json:"configs,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				HotReload     *bool              `tfsdk:"hot_reload" json:"hotReload,omitempty"`
				Resources     *[]string          `tfsdk:"resources" json:"resources,omitempty"`
				Volumes       *[]string          `tfsdk:"volumes" json:"volumes,omitempty"`
			} `tfsdk:"mount" json:"mount,omitempty"`
			Openapi *struct {
				Configmaps    *[]string          `tfsdk:"configmaps" json:"configmaps,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"openapi" json:"openapi,omitempty"`
			Owner *struct {
				Configuration     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				TargetAnnotations *[]string          `tfsdk:"target_annotations" json:"targetAnnotations,omitempty"`
				TargetLabels      *[]string          `tfsdk:"target_labels" json:"targetLabels,omitempty"`
			} `tfsdk:"owner" json:"owner,omitempty"`
			Pdb *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
			} `tfsdk:"pdb" json:"pdb,omitempty"`
			Platform *struct {
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				CreateDefault *bool              `tfsdk:"create_default" json:"createDefault,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Global        *bool              `tfsdk:"global" json:"global,omitempty"`
			} `tfsdk:"platform" json:"platform,omitempty"`
			Pod *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"pod" json:"pod,omitempty"`
			Prometheus *struct {
				Configuration    *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled          *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				PodMonitor       *bool              `tfsdk:"pod_monitor" json:"podMonitor,omitempty"`
				PodMonitorLabels *[]string          `tfsdk:"pod_monitor_labels" json:"podMonitorLabels,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
			Pull_secret *struct {
				Auto                  *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				ImagePullerDelegation *bool              `tfsdk:"image_puller_delegation" json:"imagePullerDelegation,omitempty"`
				SecretName            *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"pull_secret" json:"pull-secret,omitempty"`
			Quarkus *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				PackageTypes  *[]string          `tfsdk:"package_types" json:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" json:"quarkus,omitempty"`
			Registry *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"registry" json:"registry,omitempty"`
			Route *struct {
				Annotations                       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Configuration                     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled                           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Host                              *string            `tfsdk:"host" json:"host,omitempty"`
				TlsCACertificate                  *string            `tfsdk:"tls_ca_certificate" json:"tlsCACertificate,omitempty"`
				TlsCACertificateSecret            *string            `tfsdk:"tls_ca_certificate_secret" json:"tlsCACertificateSecret,omitempty"`
				TlsCertificate                    *string            `tfsdk:"tls_certificate" json:"tlsCertificate,omitempty"`
				TlsCertificateSecret              *string            `tfsdk:"tls_certificate_secret" json:"tlsCertificateSecret,omitempty"`
				TlsDestinationCACertificate       *string            `tfsdk:"tls_destination_ca_certificate" json:"tlsDestinationCACertificate,omitempty"`
				TlsDestinationCACertificateSecret *string            `tfsdk:"tls_destination_ca_certificate_secret" json:"tlsDestinationCACertificateSecret,omitempty"`
				TlsInsecureEdgeTerminationPolicy  *string            `tfsdk:"tls_insecure_edge_termination_policy" json:"tlsInsecureEdgeTerminationPolicy,omitempty"`
				TlsKey                            *string            `tfsdk:"tls_key" json:"tlsKey,omitempty"`
				TlsKeySecret                      *string            `tfsdk:"tls_key_secret" json:"tlsKeySecret,omitempty"`
				TlsTermination                    *string            `tfsdk:"tls_termination" json:"tlsTermination,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
			Service *struct {
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				NodePort      *bool              `tfsdk:"node_port" json:"nodePort,omitempty"`
				Type          *string            `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Service_binding *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Services      *[]string          `tfsdk:"services" json:"services,omitempty"`
			} `tfsdk:"service_binding" json:"service-binding,omitempty"`
			Strimzi *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"strimzi" json:"strimzi,omitempty"`
			Toleration *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Taints        *[]string          `tfsdk:"taints" json:"taints,omitempty"`
			} `tfsdk:"toleration" json:"toleration,omitempty"`
			Tracing *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
		} `tfsdk:"traits" json:"traits,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgIntegrationPlatformV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_integration_platform_v1"
}

func (r *CamelApacheOrgIntegrationPlatformV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IntegrationPlatform is the resource used to drive the Camel K operator behavior. It defines the behavior of all Custom Resources ('IntegrationKit', 'Integration', 'Kamelet') in the given namespace. When the Camel K operator is installed in 'global' mode, you will need to specify an 'IntegrationPlatform' in each namespace where you want the Camel K operator to be executed.",
		MarkdownDescription: "IntegrationPlatform is the resource used to drive the Camel K operator behavior. It defines the behavior of all Custom Resources ('IntegrationKit', 'Integration', 'Kamelet') in the given namespace. When the Camel K operator is installed in 'global' mode, you will need to specify an 'IntegrationPlatform' in each namespace where you want the Camel K operator to be executed.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "IntegrationPlatformSpec defines the desired state of IntegrationPlatform.",
				MarkdownDescription: "IntegrationPlatformSpec defines the desired state of IntegrationPlatform.",
				Attributes: map[string]schema.Attribute{
					"build": schema.SingleNestedAttribute{
						Description:         "specify how to build the Integration/IntegrationKits",
						MarkdownDescription: "specify how to build the Integration/IntegrationKits",
						Attributes: map[string]schema.Attribute{
							"publish_strategy_options": schema.MapAttribute{
								Description:         "Generic options that can used by any publish strategy",
								MarkdownDescription: "Generic options that can used by any publish strategy",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"base_image": schema.StringAttribute{
								Description:         "a base image that can be used as base layer for all images. It can be useful if you want to provide some custom base image with further utility softwares",
								MarkdownDescription: "a base image that can be used as base layer for all images. It can be useful if you want to provide some custom base image with further utility softwares",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"build_catalog_tool_timeout": schema.StringAttribute{
								Description:         "the timeout (in seconds) to use when creating the build tools container image",
								MarkdownDescription: "the timeout (in seconds) to use when creating the build tools container image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"build_configuration": schema.SingleNestedAttribute{
								Description:         "the configuration required to build an Integration container image",
								MarkdownDescription: "the configuration required to build an Integration container image",
								Attributes: map[string]schema.Attribute{
									"limit_cpu": schema.StringAttribute{
										Description:         "The maximum amount of CPU required. Only used for 'pod' strategy",
										MarkdownDescription: "The maximum amount of CPU required. Only used for 'pod' strategy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "The maximum amount of memory required. Only used for 'pod' strategy",
										MarkdownDescription: "The maximum amount of memory required. Only used for 'pod' strategy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"operator_namespace": schema.StringAttribute{
										Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
										MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"order_strategy": schema.StringAttribute{
										Description:         "the build order strategy to adopt",
										MarkdownDescription: "the build order strategy to adopt",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "The minimum amount of CPU required. Only used for 'pod' strategy",
										MarkdownDescription: "The minimum amount of CPU required. Only used for 'pod' strategy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "The minimum amount of memory required. Only used for 'pod' strategy",
										MarkdownDescription: "The minimum amount of memory required. Only used for 'pod' strategy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "the strategy to adopt",
										MarkdownDescription: "the strategy to adopt",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tool_image": schema.StringAttribute{
										Description:         "The container image to be used to run the build.",
										MarkdownDescription: "The container image to be used to run the build.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"maven": schema.SingleNestedAttribute{
								Description:         "Maven configuration used to build the Camel/Camel-Quarkus applications",
								MarkdownDescription: "Maven configuration used to build the Camel/Camel-Quarkus applications",
								Attributes: map[string]schema.Attribute{
									"ca_secrets": schema.ListNestedAttribute{
										Description:         "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
										MarkdownDescription: "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cli_options": schema.ListAttribute{
										Description:         "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
										MarkdownDescription: "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"extension": schema.ListNestedAttribute{
										Description:         "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
										MarkdownDescription: "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"artifact_id": schema.StringAttribute{
													Description:         "Maven Artifact",
													MarkdownDescription: "Maven Artifact",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"group_id": schema.StringAttribute{
													Description:         "Maven Group",
													MarkdownDescription: "Maven Group",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "Maven Version",
													MarkdownDescription: "Maven Version",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"local_repository": schema.StringAttribute{
										Description:         "The path of the local Maven repository.",
										MarkdownDescription: "The path of the local Maven repository.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"profiles": schema.ListNestedAttribute{
										Description:         "A reference to the ConfigMap or Secret key that contains the Maven profile.",
										MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven profile.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a ConfigMap.",
													MarkdownDescription: "Selects a key of a ConfigMap.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"secret_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a secret.",
													MarkdownDescription: "Selects a key of a secret.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its key must be defined",
															MarkdownDescription: "Specify whether the Secret or its key must be defined",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"properties": schema.MapAttribute{
										Description:         "The Maven properties.",
										MarkdownDescription: "The Maven properties.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"settings": schema.SingleNestedAttribute{
										Description:         "A reference to the ConfigMap or Secret key that contains the Maven settings.",
										MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven settings.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"settings_security": schema.SingleNestedAttribute{
										Description:         "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
										MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"max_running_builds": schema.Int64Attribute{
								Description:         "the maximum amount of parallel running pipelines started by this operator instance",
								MarkdownDescription: "the maximum amount of parallel running pipelines started by this operator instance",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"publish_strategy": schema.StringAttribute{
								Description:         "the strategy to adopt for publishing an Integration container image",
								MarkdownDescription: "the strategy to adopt for publishing an Integration container image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"registry": schema.SingleNestedAttribute{
								Description:         "the image registry used to push/pull Integration images",
								MarkdownDescription: "the image registry used to push/pull Integration images",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "the URI to access",
										MarkdownDescription: "the URI to access",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ca": schema.StringAttribute{
										Description:         "the configmap which stores the Certificate Authority",
										MarkdownDescription: "the configmap which stores the Certificate Authority",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"insecure": schema.BoolAttribute{
										Description:         "if the container registry is insecure (ie, http only)",
										MarkdownDescription: "if the container registry is insecure (ie, http only)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"organization": schema.StringAttribute{
										Description:         "the registry organization",
										MarkdownDescription: "the registry organization",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret": schema.StringAttribute{
										Description:         "the secret where credentials are stored",
										MarkdownDescription: "the secret where credentials are stored",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"runtime_provider": schema.StringAttribute{
								Description:         "the runtime used. Likely Camel Quarkus (we used to have main runtime which has been discontinued since version 1.5)",
								MarkdownDescription: "the runtime used. Likely Camel Quarkus (we used to have main runtime which has been discontinued since version 1.5)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"runtime_version": schema.StringAttribute{
								Description:         "the Camel K Runtime dependency version",
								MarkdownDescription: "the Camel K Runtime dependency version",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout": schema.StringAttribute{
								Description:         "how much time to wait before time out the pipeline process",
								MarkdownDescription: "how much time to wait before time out the pipeline process",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"cluster": schema.StringAttribute{
						Description:         "what kind of cluster you're running (ie, plain Kubernetes or OpenShift)",
						MarkdownDescription: "what kind of cluster you're running (ie, plain Kubernetes or OpenShift)",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"configuration": schema.ListNestedAttribute{
						Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes list of configuration properties to be attached to all the Integration/IntegrationKits built from this IntegrationPlatform",
						MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes list of configuration properties to be attached to all the Integration/IntegrationKits built from this IntegrationPlatform",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
									MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"kamelet": schema.SingleNestedAttribute{
						Description:         "configuration to be executed to all Kamelets controlled by this IntegrationPlatform",
						MarkdownDescription: "configuration to be executed to all Kamelets controlled by this IntegrationPlatform",
						Attributes: map[string]schema.Attribute{
							"repositories": schema.ListNestedAttribute{
								Description:         "remote repository used to retrieve Kamelet catalog",
								MarkdownDescription: "remote repository used to retrieve Kamelet catalog",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"uri": schema.StringAttribute{
											Description:         "the remote repository in the format github:ORG/REPO/PATH_TO_KAMELETS_FOLDER",
											MarkdownDescription: "the remote repository in the format github:ORG/REPO/PATH_TO_KAMELETS_FOLDER",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"profile": schema.StringAttribute{
						Description:         "the profile you wish to use. It will apply certain traits which are required by the specific profile chosen. It usually relates the Cluster with the optional definition of special profiles (ie, Knative)",
						MarkdownDescription: "the profile you wish to use. It will apply certain traits which are required by the specific profile chosen. It usually relates the Cluster with the optional definition of special profiles (ie, Knative)",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"traits": schema.SingleNestedAttribute{
						Description:         "list of traits to be executed for all the Integration/IntegrationKits built from this IntegrationPlatform",
						MarkdownDescription: "list of traits to be executed for all the Integration/IntegrationKits built from this IntegrationPlatform",
						Attributes: map[string]schema.Attribute{
							"threescale": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"addons": schema.MapAttribute{
								Description:         "The extension point with addon traits",
								MarkdownDescription: "The extension point with addon traits",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"affinity": schema.SingleNestedAttribute{
								Description:         "The configuration of Affinity trait",
								MarkdownDescription: "The configuration of Affinity trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										MarkdownDescription: "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_affinity": schema.BoolAttribute{
										Description:         "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_anti_affinity": schema.BoolAttribute{
										Description:         "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_anti_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"builder": schema.SingleNestedAttribute{
								Description:         "The configuration of Builder trait",
								MarkdownDescription: "The configuration of Builder trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"maven_profiles": schema.ListAttribute{
										Description:         "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										MarkdownDescription: "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"order_strategy": schema.StringAttribute{
										Description:         "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										MarkdownDescription: "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "The strategy to use, either 'pod' or 'routine' (default routine)",
										MarkdownDescription: "The strategy to use, either 'pod' or 'routine' (default routine)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tasks": schema.ListAttribute{
										Description:         "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										MarkdownDescription: "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"verbose": schema.BoolAttribute{
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"camel": schema.SingleNestedAttribute{
								Description:         "The configuration of Camel trait",
								MarkdownDescription: "The configuration of Camel trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"runtime_version": schema.StringAttribute{
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"container": schema.SingleNestedAttribute{
								Description:         "The configuration of Container trait",
								MarkdownDescription: "The configuration of Container trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically enable the trait",
										MarkdownDescription: "To automatically enable the trait",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"expose": schema.BoolAttribute{
										Description:         "Can be used to enable/disable exposure via kubernetes Service.",
										MarkdownDescription: "Can be used to enable/disable exposure via kubernetes Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image": schema.StringAttribute{
										Description:         "The main container image",
										MarkdownDescription: "The main container image",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "The pull policy: Always|Never|IfNotPresent",
										MarkdownDescription: "The pull policy: Always|Never|IfNotPresent",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "The maximum amount of CPU required.",
										MarkdownDescription: "The maximum amount of CPU required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "The maximum amount of memory required.",
										MarkdownDescription: "The maximum amount of memory required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "The main container name. It's named 'integration' by default.",
										MarkdownDescription: "The main container name. It's named 'integration' by default.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "To configure a different port exposed by the container (default '8080').",
										MarkdownDescription: "To configure a different port exposed by the container (default '8080').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port_name": schema.StringAttribute{
										Description:         "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										MarkdownDescription: "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "The minimum amount of CPU required.",
										MarkdownDescription: "The minimum amount of CPU required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "The minimum amount of memory required.",
										MarkdownDescription: "The minimum amount of memory required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_port": schema.Int64Attribute{
										Description:         "To configure under which service port the container port is to be exposed (default '80').",
										MarkdownDescription: "To configure under which service port the container port is to be exposed (default '80').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_port_name": schema.StringAttribute{
										Description:         "To configure under which service port name the container port is to be exposed (default 'http').",
										MarkdownDescription: "To configure under which service port name the container port is to be exposed (default 'http').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"cron": schema.SingleNestedAttribute{
								Description:         "The configuration of Cron trait",
								MarkdownDescription: "The configuration of Cron trait",
								Attributes: map[string]schema.Attribute{
									"active_deadline_seconds": schema.Int64Attribute{
										Description:         "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"auto": schema.BoolAttribute{
										Description:         "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										MarkdownDescription: "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"backoff_limit": schema.Int64Attribute{
										Description:         "Specifies the number of retries before marking the job failed. It defaults to 2.",
										MarkdownDescription: "Specifies the number of retries before marking the job failed. It defaults to 2.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"components": schema.StringAttribute{
										Description:         "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										MarkdownDescription: "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"concurrency_policy": schema.StringAttribute{
										Description:         "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										MarkdownDescription: "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"fallback": schema.BoolAttribute{
										Description:         "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										MarkdownDescription: "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"schedule": schema.StringAttribute{
										Description:         "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										MarkdownDescription: "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"starting_deadline_seconds": schema.Int64Attribute{
										Description:         "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										MarkdownDescription: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"dependencies": schema.SingleNestedAttribute{
								Description:         "The configuration of Dependencies trait",
								MarkdownDescription: "The configuration of Dependencies trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"deployer": schema.SingleNestedAttribute{
								Description:         "The configuration of Deployer trait",
								MarkdownDescription: "The configuration of Deployer trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kind": schema.StringAttribute{
										Description:         "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										MarkdownDescription: "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"use_ssa": schema.BoolAttribute{
										Description:         "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										MarkdownDescription: "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"deployment": schema.SingleNestedAttribute{
								Description:         "The configuration of Deployment trait",
								MarkdownDescription: "The configuration of Deployment trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"progress_deadline_seconds": schema.Int64Attribute{
										Description:         "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rolling_update_max_surge": schema.Int64Attribute{
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rolling_update_max_unavailable": schema.Int64Attribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "The deployment strategy to use to replace existing pods with new ones.",
										MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"environment": schema.SingleNestedAttribute{
								Description:         "The configuration of Environment trait",
								MarkdownDescription: "The configuration of Environment trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"container_meta": schema.BoolAttribute{
										Description:         "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										MarkdownDescription: "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"http_proxy": schema.BoolAttribute{
										Description:         "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										MarkdownDescription: "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vars": schema.ListAttribute{
										Description:         "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
										MarkdownDescription: "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"error_handler": schema.SingleNestedAttribute{
								Description:         "The configuration of Error Handler trait",
								MarkdownDescription: "The configuration of Error Handler trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ref": schema.StringAttribute{
										Description:         "The error handler ref name provided or found in application properties",
										MarkdownDescription: "The error handler ref name provided or found in application properties",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"gc": schema.SingleNestedAttribute{
								Description:         "The configuration of GC trait",
								MarkdownDescription: "The configuration of GC trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"discovery_cache": schema.StringAttribute{
										Description:         "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										MarkdownDescription: "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"health": schema.SingleNestedAttribute{
								Description:         "The configuration of Health trait",
								MarkdownDescription: "The configuration of Health trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the liveness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the liveness probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_period": schema.Int64Attribute{
										Description:         "How often to perform the liveness probe.",
										MarkdownDescription: "How often to perform the liveness probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the liveness probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the liveness probe for the integration container (default 'false').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the liveness probe times out.",
										MarkdownDescription: "Number of seconds after which the liveness probe times out.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the readiness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the readiness probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_period": schema.Int64Attribute{
										Description:         "How often to perform the readiness probe.",
										MarkdownDescription: "How often to perform the readiness probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the readiness probe for the integration container (default 'true').",
										MarkdownDescription: "Configures the readiness probe for the integration container (default 'true').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the readiness probe times out.",
										MarkdownDescription: "Number of seconds after which the readiness probe times out.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the startup probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the startup probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_period": schema.Int64Attribute{
										Description:         "How often to perform the startup probe.",
										MarkdownDescription: "How often to perform the startup probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the startup probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the startup probe for the integration container (default 'false').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the startup probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the startup probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the startup probe times out.",
										MarkdownDescription: "Number of seconds after which the startup probe times out.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ingress": schema.SingleNestedAttribute{
								Description:         "The configuration of Ingress trait",
								MarkdownDescription: "The configuration of Ingress trait",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
										MarkdownDescription: "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"auto": schema.BoolAttribute{
										Description:         "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										MarkdownDescription: "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "To configure the host exposed by the ingress.",
										MarkdownDescription: "To configure the host exposed by the ingress.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path": schema.StringAttribute{
										Description:         "To configure the path exposed by the ingress (default '/').",
										MarkdownDescription: "To configure the path exposed by the ingress (default '/').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path_type": schema.StringAttribute{
										Description:         "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
										MarkdownDescription: "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"istio": schema.SingleNestedAttribute{
								Description:         "The configuration of Istio trait",
								MarkdownDescription: "The configuration of Istio trait",
								Attributes: map[string]schema.Attribute{
									"allow": schema.StringAttribute{
										Description:         "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										MarkdownDescription: "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"inject": schema.BoolAttribute{
										Description:         "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										MarkdownDescription: "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"jolokia": schema.SingleNestedAttribute{
								Description:         "The configuration of Jolokia trait",
								MarkdownDescription: "The configuration of Jolokia trait",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.StringAttribute{
										Description:         "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										MarkdownDescription: "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"client_principal": schema.ListAttribute{
										Description:         "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										MarkdownDescription: "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"discovery_enabled": schema.BoolAttribute{
										Description:         "Listen for multicast requests (default 'false')",
										MarkdownDescription: "Listen for multicast requests (default 'false')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"extended_client_check": schema.BoolAttribute{
										Description:         "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										MarkdownDescription: "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										MarkdownDescription: "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"options": schema.ListAttribute{
										Description:         "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										MarkdownDescription: "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"password": schema.StringAttribute{
										Description:         "The password used for authentication, applicable when the 'user' option is set.",
										MarkdownDescription: "The password used for authentication, applicable when the 'user' option is set.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "The Jolokia endpoint port (default '8778').",
										MarkdownDescription: "The Jolokia endpoint port (default '8778').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"protocol": schema.StringAttribute{
										Description:         "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										MarkdownDescription: "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"use_ssl_client_authentication": schema.BoolAttribute{
										Description:         "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										MarkdownDescription: "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"user": schema.StringAttribute{
										Description:         "The user to be used for authentication",
										MarkdownDescription: "The user to be used for authentication",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"jvm": schema.SingleNestedAttribute{
								Description:         "The configuration of JVM trait",
								MarkdownDescription: "The configuration of JVM trait",
								Attributes: map[string]schema.Attribute{
									"classpath": schema.StringAttribute{
										Description:         "Additional JVM classpath (use 'Linux' classpath separator)",
										MarkdownDescription: "Additional JVM classpath (use 'Linux' classpath separator)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug": schema.BoolAttribute{
										Description:         "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										MarkdownDescription: "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug_address": schema.StringAttribute{
										Description:         "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										MarkdownDescription: "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug_suspend": schema.BoolAttribute{
										Description:         "Suspends the target JVM immediately before the main class is loaded",
										MarkdownDescription: "Suspends the target JVM immediately before the main class is loaded",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"options": schema.ListAttribute{
										Description:         "A list of JVM options",
										MarkdownDescription: "A list of JVM options",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"print_command": schema.BoolAttribute{
										Description:         "Prints the command used the start the JVM in the container logs (default 'true')",
										MarkdownDescription: "Prints the command used the start the JVM in the container logs (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"kamelets": schema.SingleNestedAttribute{
								Description:         "The configuration of Kamelets trait",
								MarkdownDescription: "The configuration of Kamelets trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										MarkdownDescription: "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"list": schema.StringAttribute{
										Description:         "Comma separated list of Kamelet names to load into the current integration",
										MarkdownDescription: "Comma separated list of Kamelet names to load into the current integration",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"keda": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"knative": schema.SingleNestedAttribute{
								Description:         "The configuration of Knative trait",
								MarkdownDescription: "The configuration of Knative trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Enable automatic discovery of all trait properties.",
										MarkdownDescription: "Enable automatic discovery of all trait properties.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"channel_sinks": schema.ListAttribute{
										Description:         "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"channel_sources": schema.ListAttribute{
										Description:         "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"config": schema.StringAttribute{
										Description:         "Can be used to inject a Knative complete configuration in JSON format.",
										MarkdownDescription: "Can be used to inject a Knative complete configuration in JSON format.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"endpoint_sinks": schema.ListAttribute{
										Description:         "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										MarkdownDescription: "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"endpoint_sources": schema.ListAttribute{
										Description:         "List of channels used as source of integration routes.",
										MarkdownDescription: "List of channels used as source of integration routes.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"event_sinks": schema.ListAttribute{
										Description:         "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										MarkdownDescription: "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"event_sources": schema.ListAttribute{
										Description:         "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										MarkdownDescription: "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"filter_source_channels": schema.BoolAttribute{
										Description:         "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										MarkdownDescription: "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace_label": schema.BoolAttribute{
										Description:         "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
										MarkdownDescription: "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"sink_binding": schema.BoolAttribute{
										Description:         "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										MarkdownDescription: "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"knative_service": schema.SingleNestedAttribute{
								Description:         "The configuration of Knative Service trait",
								MarkdownDescription: "The configuration of Knative Service trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										MarkdownDescription: "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"autoscaling_metric": schema.StringAttribute{
										Description:         "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"autoscaling_target": schema.Int64Attribute{
										Description:         "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"class": schema.StringAttribute{
										Description:         "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_scale": schema.Int64Attribute{
										Description:         "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_scale": schema.Int64Attribute{
										Description:         "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rollout_duration": schema.StringAttribute{
										Description:         "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										MarkdownDescription: "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"visibility": schema.StringAttribute{
										Description:         "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "The configuration of Logging trait",
								MarkdownDescription: "The configuration of Logging trait",
								Attributes: map[string]schema.Attribute{
									"color": schema.BoolAttribute{
										Description:         "Colorize the log output",
										MarkdownDescription: "Colorize the log output",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"format": schema.StringAttribute{
										Description:         "Logs message format",
										MarkdownDescription: "Logs message format",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"json": schema.BoolAttribute{
										Description:         "Output the logs in JSON",
										MarkdownDescription: "Output the logs in JSON",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"json_pretty_print": schema.BoolAttribute{
										Description:         "Enable 'pretty printing' of the JSON logs",
										MarkdownDescription: "Enable 'pretty printing' of the JSON logs",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"level": schema.StringAttribute{
										Description:         "Adjust the logging level (defaults to INFO)",
										MarkdownDescription: "Adjust the logging level (defaults to INFO)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"master": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mount": schema.SingleNestedAttribute{
								Description:         "The configuration of Mount trait",
								MarkdownDescription: "The configuration of Mount trait",
								Attributes: map[string]schema.Attribute{
									"configs": schema.ListAttribute{
										Description:         "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										MarkdownDescription: "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"hot_reload": schema.BoolAttribute{
										Description:         "Enable 'hot reload' when a secret/configmap mounted is edited (default false)",
										MarkdownDescription: "Enable 'hot reload' when a secret/configmap mounted is edited (default false)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resources": schema.ListAttribute{
										Description:         "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										MarkdownDescription: "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volumes": schema.ListAttribute{
										Description:         "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
										MarkdownDescription: "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"openapi": schema.SingleNestedAttribute{
								Description:         "The configuration of OpenAPI trait",
								MarkdownDescription: "The configuration of OpenAPI trait",
								Attributes: map[string]schema.Attribute{
									"configmaps": schema.ListAttribute{
										Description:         "The configmaps holding the spec of the OpenAPI",
										MarkdownDescription: "The configmaps holding the spec of the OpenAPI",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"owner": schema.SingleNestedAttribute{
								Description:         "The configuration of Owner trait",
								MarkdownDescription: "The configuration of Owner trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_annotations": schema.ListAttribute{
										Description:         "The set of annotations to be transferred",
										MarkdownDescription: "The set of annotations to be transferred",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_labels": schema.ListAttribute{
										Description:         "The set of labels to be transferred",
										MarkdownDescription: "The set of labels to be transferred",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pdb": schema.SingleNestedAttribute{
								Description:         "The configuration of PDB trait",
								MarkdownDescription: "The configuration of PDB trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_unavailable": schema.StringAttribute{
										Description:         "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_available": schema.StringAttribute{
										Description:         "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"platform": schema.SingleNestedAttribute{
								Description:         "The configuration of Platform trait",
								MarkdownDescription: "The configuration of Platform trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										MarkdownDescription: "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"create_default": schema.BoolAttribute{
										Description:         "To create a default (empty) platform when the platform is missing.",
										MarkdownDescription: "To create a default (empty) platform when the platform is missing.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"global": schema.BoolAttribute{
										Description:         "Indicates if the platform should be created globally in the case of global operator (default true).",
										MarkdownDescription: "Indicates if the platform should be created globally in the case of global operator (default true).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pod": schema.SingleNestedAttribute{
								Description:         "The configuration of Pod trait",
								MarkdownDescription: "The configuration of Pod trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"prometheus": schema.SingleNestedAttribute{
								Description:         "The configuration of Prometheus trait",
								MarkdownDescription: "The configuration of Prometheus trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_monitor": schema.BoolAttribute{
										Description:         "Whether a 'PodMonitor' resource is created (default 'true').",
										MarkdownDescription: "Whether a 'PodMonitor' resource is created (default 'true').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_monitor_labels": schema.ListAttribute{
										Description:         "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
										MarkdownDescription: "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pull_secret": schema.SingleNestedAttribute{
								Description:         "The configuration of Pull Secret trait",
								MarkdownDescription: "The configuration of Pull Secret trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										MarkdownDescription: "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_puller_delegation": schema.BoolAttribute{
										Description:         "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										MarkdownDescription: "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										MarkdownDescription: "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"quarkus": schema.SingleNestedAttribute{
								Description:         "The configuration of Quarkus trait",
								MarkdownDescription: "The configuration of Quarkus trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"package_types": schema.ListAttribute{
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"registry": schema.SingleNestedAttribute{
								Description:         "The configuration of Registry trait",
								MarkdownDescription: "The configuration of Registry trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"route": schema.SingleNestedAttribute{
								Description:         "The configuration of Route trait",
								MarkdownDescription: "The configuration of Route trait",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "The annotations added to route. This can be used to set route specific annotations For annotations options see https://docs.openshift.com/container-platform/3.11/architecture/networking/routes.html#route-specific-annotations CLI usage example: -t 'route.annotations.'haproxy.router.openshift.io/balance'=true'",
										MarkdownDescription: "The annotations added to route. This can be used to set route specific annotations For annotations options see https://docs.openshift.com/container-platform/3.11/architecture/networking/routes.html#route-specific-annotations CLI usage example: -t 'route.annotations.'haproxy.router.openshift.io/balance'=true'",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "To configure the host exposed by the route.",
										MarkdownDescription: "To configure the host exposed by the route.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_ca_certificate": schema.StringAttribute{
										Description:         "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_ca_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_certificate": schema.StringAttribute{
										Description:         "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_destination_ca_certificate": schema.StringAttribute{
										Description:         "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_destination_ca_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_insecure_edge_termination_policy": schema.StringAttribute{
										Description:         "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_key": schema.StringAttribute{
										Description:         "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_key_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_termination": schema.StringAttribute{
										Description:         "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"service": schema.SingleNestedAttribute{
								Description:         "The configuration of Service trait",
								MarkdownDescription: "The configuration of Service trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically detect from the code if a Service needs to be created.",
										MarkdownDescription: "To automatically detect from the code if a Service needs to be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_port": schema.BoolAttribute{
										Description:         "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										MarkdownDescription: "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										MarkdownDescription: "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"service_binding": schema.SingleNestedAttribute{
								Description:         "The configuration of Service Binding trait",
								MarkdownDescription: "The configuration of Service Binding trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"services": schema.ListAttribute{
										Description:         "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
										MarkdownDescription: "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"strimzi": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"toleration": schema.SingleNestedAttribute{
								Description:         "The configuration of Toleration trait",
								MarkdownDescription: "The configuration of Toleration trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"taints": schema.ListAttribute{
										Description:         "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
										MarkdownDescription: "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"tracing": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CamelApacheOrgIntegrationPlatformV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CamelApacheOrgIntegrationPlatformV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_integration_platform_v1")

	var data CamelApacheOrgIntegrationPlatformV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "IntegrationPlatform"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CamelApacheOrgIntegrationPlatformV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("IntegrationPlatform")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
