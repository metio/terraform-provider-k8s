/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CamelApacheOrgIntegrationPlatformV1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgIntegrationPlatformV1Resource)(nil)
)

type CamelApacheOrgIntegrationPlatformV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgIntegrationPlatformV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Build *struct {
			PublishStrategyOptions *map[string]string `tfsdk:"publish_strategy_options" yaml:"PublishStrategyOptions,omitempty"`

			BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

			BuildStrategy *string `tfsdk:"build_strategy" yaml:"buildStrategy,omitempty"`

			KanikoBuildCache *bool `tfsdk:"kaniko_build_cache" yaml:"kanikoBuildCache,omitempty"`

			Maven *struct {
				CaSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"ca_secret" yaml:"caSecret,omitempty"`

				CaSecrets *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"ca_secrets" yaml:"caSecrets,omitempty"`

				CliOptions *[]string `tfsdk:"cli_options" yaml:"cliOptions,omitempty"`

				Extension *[]struct {
					ArtifactId *string `tfsdk:"artifact_id" yaml:"artifactId,omitempty"`

					GroupId *string `tfsdk:"group_id" yaml:"groupId,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"extension" yaml:"extension,omitempty"`

				LocalRepository *string `tfsdk:"local_repository" yaml:"localRepository,omitempty"`

				Properties *map[string]string `tfsdk:"properties" yaml:"properties,omitempty"`

				Settings *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"settings" yaml:"settings,omitempty"`

				SettingsSecurity *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"settings_security" yaml:"settingsSecurity,omitempty"`
			} `tfsdk:"maven" yaml:"maven,omitempty"`

			PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

			PublishStrategy *string `tfsdk:"publish_strategy" yaml:"publishStrategy,omitempty"`

			Registry *struct {
				Address *string `tfsdk:"address" yaml:"address,omitempty"`

				Ca *string `tfsdk:"ca" yaml:"ca,omitempty"`

				Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

				Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`

				Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"registry" yaml:"registry,omitempty"`

			RuntimeProvider *string `tfsdk:"runtime_provider" yaml:"runtimeProvider,omitempty"`

			RuntimeVersion *string `tfsdk:"runtime_version" yaml:"runtimeVersion,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"build" yaml:"build,omitempty"`

		Cluster *string `tfsdk:"cluster" yaml:"cluster,omitempty"`

		Configuration *[]struct {
			ResourceKey *string `tfsdk:"resource_key" yaml:"resourceKey,omitempty"`

			ResourceMountPoint *string `tfsdk:"resource_mount_point" yaml:"resourceMountPoint,omitempty"`

			ResourceType *string `tfsdk:"resource_type" yaml:"resourceType,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"configuration" yaml:"configuration,omitempty"`

		Kamelet *struct {
			Repositories *[]struct {
				Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
			} `tfsdk:"repositories" yaml:"repositories,omitempty"`
		} `tfsdk:"kamelet" yaml:"kamelet,omitempty"`

		Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`

		Resources *map[string]string `tfsdk:"resources" yaml:"resources,omitempty"`

		Traits *struct {
			Threescale *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"threescale" yaml:"3scale,omitempty"`

			Addons utilities.Dynamic `tfsdk:"addons" yaml:"addons,omitempty"`

			Affinity *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				NodeAffinityLabels *[]string `tfsdk:"node_affinity_labels" yaml:"nodeAffinityLabels,omitempty"`

				PodAffinity *bool `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

				PodAffinityLabels *[]string `tfsdk:"pod_affinity_labels" yaml:"podAffinityLabels,omitempty"`

				PodAntiAffinity *bool `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

				PodAntiAffinityLabels *[]string `tfsdk:"pod_anti_affinity_labels" yaml:"podAntiAffinityLabels,omitempty"`
			} `tfsdk:"affinity" yaml:"affinity,omitempty"`

			Builder *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Properties *[]string `tfsdk:"properties" yaml:"properties,omitempty"`

				Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`
			} `tfsdk:"builder" yaml:"builder,omitempty"`

			Camel *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Properties *[]string `tfsdk:"properties" yaml:"properties,omitempty"`

				RuntimeVersion *string `tfsdk:"runtime_version" yaml:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" yaml:"camel,omitempty"`

			Container *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Expose *bool `tfsdk:"expose" yaml:"expose,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				LimitCPU *string `tfsdk:"limit_cpu" yaml:"limitCPU,omitempty"`

				LimitMemory *string `tfsdk:"limit_memory" yaml:"limitMemory,omitempty"`

				LivenessFailureThreshold *int64 `tfsdk:"liveness_failure_threshold" yaml:"livenessFailureThreshold,omitempty"`

				LivenessInitialDelay *int64 `tfsdk:"liveness_initial_delay" yaml:"livenessInitialDelay,omitempty"`

				LivenessPeriod *int64 `tfsdk:"liveness_period" yaml:"livenessPeriod,omitempty"`

				LivenessScheme *string `tfsdk:"liveness_scheme" yaml:"livenessScheme,omitempty"`

				LivenessSuccessThreshold *int64 `tfsdk:"liveness_success_threshold" yaml:"livenessSuccessThreshold,omitempty"`

				LivenessTimeout *int64 `tfsdk:"liveness_timeout" yaml:"livenessTimeout,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

				ProbesEnabled *bool `tfsdk:"probes_enabled" yaml:"probesEnabled,omitempty"`

				ReadinessFailureThreshold *int64 `tfsdk:"readiness_failure_threshold" yaml:"readinessFailureThreshold,omitempty"`

				ReadinessInitialDelay *int64 `tfsdk:"readiness_initial_delay" yaml:"readinessInitialDelay,omitempty"`

				ReadinessPeriod *int64 `tfsdk:"readiness_period" yaml:"readinessPeriod,omitempty"`

				ReadinessScheme *string `tfsdk:"readiness_scheme" yaml:"readinessScheme,omitempty"`

				ReadinessSuccessThreshold *int64 `tfsdk:"readiness_success_threshold" yaml:"readinessSuccessThreshold,omitempty"`

				ReadinessTimeout *int64 `tfsdk:"readiness_timeout" yaml:"readinessTimeout,omitempty"`

				RequestCPU *string `tfsdk:"request_cpu" yaml:"requestCPU,omitempty"`

				RequestMemory *string `tfsdk:"request_memory" yaml:"requestMemory,omitempty"`

				ServicePort *int64 `tfsdk:"service_port" yaml:"servicePort,omitempty"`

				ServicePortName *string `tfsdk:"service_port_name" yaml:"servicePortName,omitempty"`
			} `tfsdk:"container" yaml:"container,omitempty"`

			Cron *struct {
				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				BackoffLimit *int64 `tfsdk:"backoff_limit" yaml:"backoffLimit,omitempty"`

				Components *string `tfsdk:"components" yaml:"components,omitempty"`

				ConcurrencyPolicy *string `tfsdk:"concurrency_policy" yaml:"concurrencyPolicy,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Fallback *bool `tfsdk:"fallback" yaml:"fallback,omitempty"`

				Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

				StartingDeadlineSeconds *int64 `tfsdk:"starting_deadline_seconds" yaml:"startingDeadlineSeconds,omitempty"`
			} `tfsdk:"cron" yaml:"cron,omitempty"`

			Dependencies *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

			Deployer *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				UseSSA *bool `tfsdk:"use_ssa" yaml:"useSSA,omitempty"`
			} `tfsdk:"deployer" yaml:"deployer,omitempty"`

			Deployment *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ProgressDeadlineSeconds *int64 `tfsdk:"progress_deadline_seconds" yaml:"progressDeadlineSeconds,omitempty"`

				RollingUpdateMaxSurge *int64 `tfsdk:"rolling_update_max_surge" yaml:"rollingUpdateMaxSurge,omitempty"`

				RollingUpdateMaxUnavailable *int64 `tfsdk:"rolling_update_max_unavailable" yaml:"rollingUpdateMaxUnavailable,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`
			} `tfsdk:"deployment" yaml:"deployment,omitempty"`

			Environment *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				ContainerMeta *bool `tfsdk:"container_meta" yaml:"containerMeta,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				HttpProxy *bool `tfsdk:"http_proxy" yaml:"httpProxy,omitempty"`

				Vars *[]string `tfsdk:"vars" yaml:"vars,omitempty"`
			} `tfsdk:"environment" yaml:"environment,omitempty"`

			Error_handler *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Ref *string `tfsdk:"ref" yaml:"ref,omitempty"`
			} `tfsdk:"error_handler" yaml:"error-handler,omitempty"`

			Gc *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				DiscoveryCache *string `tfsdk:"discovery_cache" yaml:"discoveryCache,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"gc" yaml:"gc,omitempty"`

			Health *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				LivenessFailureThreshold *int64 `tfsdk:"liveness_failure_threshold" yaml:"livenessFailureThreshold,omitempty"`

				LivenessInitialDelay *int64 `tfsdk:"liveness_initial_delay" yaml:"livenessInitialDelay,omitempty"`

				LivenessPeriod *int64 `tfsdk:"liveness_period" yaml:"livenessPeriod,omitempty"`

				LivenessProbeEnabled *bool `tfsdk:"liveness_probe_enabled" yaml:"livenessProbeEnabled,omitempty"`

				LivenessScheme *string `tfsdk:"liveness_scheme" yaml:"livenessScheme,omitempty"`

				LivenessSuccessThreshold *int64 `tfsdk:"liveness_success_threshold" yaml:"livenessSuccessThreshold,omitempty"`

				LivenessTimeout *int64 `tfsdk:"liveness_timeout" yaml:"livenessTimeout,omitempty"`

				ReadinessFailureThreshold *int64 `tfsdk:"readiness_failure_threshold" yaml:"readinessFailureThreshold,omitempty"`

				ReadinessInitialDelay *int64 `tfsdk:"readiness_initial_delay" yaml:"readinessInitialDelay,omitempty"`

				ReadinessPeriod *int64 `tfsdk:"readiness_period" yaml:"readinessPeriod,omitempty"`

				ReadinessProbeEnabled *bool `tfsdk:"readiness_probe_enabled" yaml:"readinessProbeEnabled,omitempty"`

				ReadinessScheme *string `tfsdk:"readiness_scheme" yaml:"readinessScheme,omitempty"`

				ReadinessSuccessThreshold *int64 `tfsdk:"readiness_success_threshold" yaml:"readinessSuccessThreshold,omitempty"`

				ReadinessTimeout *int64 `tfsdk:"readiness_timeout" yaml:"readinessTimeout,omitempty"`
			} `tfsdk:"health" yaml:"health,omitempty"`

			Ingress *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Istio *struct {
				Allow *string `tfsdk:"allow" yaml:"allow,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Inject *bool `tfsdk:"inject" yaml:"inject,omitempty"`
			} `tfsdk:"istio" yaml:"istio,omitempty"`

			Jolokia *struct {
				CACert *string `tfsdk:"ca_cert" yaml:"CACert,omitempty"`

				ClientPrincipal *[]string `tfsdk:"client_principal" yaml:"clientPrincipal,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				DiscoveryEnabled *bool `tfsdk:"discovery_enabled" yaml:"discoveryEnabled,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ExtendedClientCheck *bool `tfsdk:"extended_client_check" yaml:"extendedClientCheck,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

				Password *string `tfsdk:"password" yaml:"password,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				UseSSLClientAuthentication *bool `tfsdk:"use_ssl_client_authentication" yaml:"useSSLClientAuthentication,omitempty"`

				User *string `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"jolokia" yaml:"jolokia,omitempty"`

			Jvm *struct {
				Classpath *string `tfsdk:"classpath" yaml:"classpath,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Debug *bool `tfsdk:"debug" yaml:"debug,omitempty"`

				DebugAddress *string `tfsdk:"debug_address" yaml:"debugAddress,omitempty"`

				DebugSuspend *bool `tfsdk:"debug_suspend" yaml:"debugSuspend,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

				PrintCommand *bool `tfsdk:"print_command" yaml:"printCommand,omitempty"`
			} `tfsdk:"jvm" yaml:"jvm,omitempty"`

			Kamelets *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				List *string `tfsdk:"list" yaml:"list,omitempty"`
			} `tfsdk:"kamelets" yaml:"kamelets,omitempty"`

			Keda *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"keda" yaml:"keda,omitempty"`

			Knative *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				ChannelSinks *[]string `tfsdk:"channel_sinks" yaml:"channelSinks,omitempty"`

				ChannelSources *[]string `tfsdk:"channel_sources" yaml:"channelSources,omitempty"`

				Config *string `tfsdk:"config" yaml:"config,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				EndpointSinks *[]string `tfsdk:"endpoint_sinks" yaml:"endpointSinks,omitempty"`

				EndpointSources *[]string `tfsdk:"endpoint_sources" yaml:"endpointSources,omitempty"`

				EventSinks *[]string `tfsdk:"event_sinks" yaml:"eventSinks,omitempty"`

				EventSources *[]string `tfsdk:"event_sources" yaml:"eventSources,omitempty"`

				FilterSourceChannels *bool `tfsdk:"filter_source_channels" yaml:"filterSourceChannels,omitempty"`

				SinkBinding *bool `tfsdk:"sink_binding" yaml:"sinkBinding,omitempty"`
			} `tfsdk:"knative" yaml:"knative,omitempty"`

			Knative_service *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				AutoscalingMetric *string `tfsdk:"autoscaling_metric" yaml:"autoscalingMetric,omitempty"`

				AutoscalingTarget *int64 `tfsdk:"autoscaling_target" yaml:"autoscalingTarget,omitempty"`

				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				MaxScale *int64 `tfsdk:"max_scale" yaml:"maxScale,omitempty"`

				MinScale *int64 `tfsdk:"min_scale" yaml:"minScale,omitempty"`

				RolloutDuration *string `tfsdk:"rollout_duration" yaml:"rolloutDuration,omitempty"`

				Visibility *string `tfsdk:"visibility" yaml:"visibility,omitempty"`
			} `tfsdk:"knative_service" yaml:"knative-service,omitempty"`

			Logging *struct {
				Color *bool `tfsdk:"color" yaml:"color,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Format *string `tfsdk:"format" yaml:"format,omitempty"`

				Json *bool `tfsdk:"json" yaml:"json,omitempty"`

				JsonPrettyPrint *bool `tfsdk:"json_pretty_print" yaml:"jsonPrettyPrint,omitempty"`

				Level *string `tfsdk:"level" yaml:"level,omitempty"`
			} `tfsdk:"logging" yaml:"logging,omitempty"`

			Master *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"master" yaml:"master,omitempty"`

			Mount *struct {
				Configs *[]string `tfsdk:"configs" yaml:"configs,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

				Volumes *[]string `tfsdk:"volumes" yaml:"volumes,omitempty"`
			} `tfsdk:"mount" yaml:"mount,omitempty"`

			Openapi *struct {
				Configmaps *[]string `tfsdk:"configmaps" yaml:"configmaps,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"openapi" yaml:"openapi,omitempty"`

			Owner *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				TargetAnnotations *[]string `tfsdk:"target_annotations" yaml:"targetAnnotations,omitempty"`

				TargetLabels *[]string `tfsdk:"target_labels" yaml:"targetLabels,omitempty"`
			} `tfsdk:"owner" yaml:"owner,omitempty"`

			Pdb *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				MaxUnavailable *string `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

				MinAvailable *string `tfsdk:"min_available" yaml:"minAvailable,omitempty"`
			} `tfsdk:"pdb" yaml:"pdb,omitempty"`

			Platform *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				CreateDefault *bool `tfsdk:"create_default" yaml:"createDefault,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Global *bool `tfsdk:"global" yaml:"global,omitempty"`
			} `tfsdk:"platform" yaml:"platform,omitempty"`

			Pod *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"pod" yaml:"pod,omitempty"`

			Prometheus *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				PodMonitor *bool `tfsdk:"pod_monitor" yaml:"podMonitor,omitempty"`

				PodMonitorLabels *[]string `tfsdk:"pod_monitor_labels" yaml:"podMonitorLabels,omitempty"`
			} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`

			Pull_secret *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ImagePullerDelegation *bool `tfsdk:"image_puller_delegation" yaml:"imagePullerDelegation,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"pull_secret" yaml:"pull-secret,omitempty"`

			Quarkus *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				PackageTypes *[]string `tfsdk:"package_types" yaml:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" yaml:"quarkus,omitempty"`

			Registry *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"registry" yaml:"registry,omitempty"`

			Route *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				TlsCACertificate *string `tfsdk:"tls_ca_certificate" yaml:"tlsCACertificate,omitempty"`

				TlsCACertificateSecret *string `tfsdk:"tls_ca_certificate_secret" yaml:"tlsCACertificateSecret,omitempty"`

				TlsCertificate *string `tfsdk:"tls_certificate" yaml:"tlsCertificate,omitempty"`

				TlsCertificateSecret *string `tfsdk:"tls_certificate_secret" yaml:"tlsCertificateSecret,omitempty"`

				TlsDestinationCACertificate *string `tfsdk:"tls_destination_ca_certificate" yaml:"tlsDestinationCACertificate,omitempty"`

				TlsDestinationCACertificateSecret *string `tfsdk:"tls_destination_ca_certificate_secret" yaml:"tlsDestinationCACertificateSecret,omitempty"`

				TlsInsecureEdgeTerminationPolicy *string `tfsdk:"tls_insecure_edge_termination_policy" yaml:"tlsInsecureEdgeTerminationPolicy,omitempty"`

				TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`

				TlsKeySecret *string `tfsdk:"tls_key_secret" yaml:"tlsKeySecret,omitempty"`

				TlsTermination *string `tfsdk:"tls_termination" yaml:"tlsTermination,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Service *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				NodePort *bool `tfsdk:"node_port" yaml:"nodePort,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"service" yaml:"service,omitempty"`

			Service_binding *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Services *[]string `tfsdk:"services" yaml:"services,omitempty"`
			} `tfsdk:"service_binding" yaml:"service-binding,omitempty"`

			Strimzi *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"strimzi" yaml:"strimzi,omitempty"`

			Toleration *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Taints *[]string `tfsdk:"taints" yaml:"taints,omitempty"`
			} `tfsdk:"toleration" yaml:"toleration,omitempty"`

			Tracing *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"tracing" yaml:"tracing,omitempty"`
		} `tfsdk:"traits" yaml:"traits,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgIntegrationPlatformV1Resource() resource.Resource {
	return &CamelApacheOrgIntegrationPlatformV1Resource{}
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_integration_platform_v1"
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "IntegrationPlatform is the resource used to drive the Camel K operator behavior. It defines the behavior of all Custom Resources ('IntegrationKit', 'Integration', 'Kamelet') in the given namespace. When the Camel K operator is installed in 'global' mode, you will need to specify an 'IntegrationPlatform' in each namespace where you want the Camel K operator to be executed",
		MarkdownDescription: "IntegrationPlatform is the resource used to drive the Camel K operator behavior. It defines the behavior of all Custom Resources ('IntegrationKit', 'Integration', 'Kamelet') in the given namespace. When the Camel K operator is installed in 'global' mode, you will need to specify an 'IntegrationPlatform' in each namespace where you want the Camel K operator to be executed",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "IntegrationPlatformSpec defines the desired state of IntegrationPlatform",
				MarkdownDescription: "IntegrationPlatformSpec defines the desired state of IntegrationPlatform",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"build": {
						Description:         "specify how to build the Integration/IntegrationKits",
						MarkdownDescription: "specify how to build the Integration/IntegrationKits",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"publish_strategy_options": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"base_image": {
								Description:         "a base image that can be used as base layer for all images. It can be useful if you want to provide some custom base image with further utility softwares",
								MarkdownDescription: "a base image that can be used as base layer for all images. It can be useful if you want to provide some custom base image with further utility softwares",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"build_strategy": {
								Description:         "the strategy to adopt for building an Integration base image",
								MarkdownDescription: "the strategy to adopt for building an Integration base image",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("routine", "pod"),
								},
							},

							"kaniko_build_cache": {
								Description:         "Deprecated: Use PublishStrategyOptions instead enables Kaniko publish strategy cache",
								MarkdownDescription: "Deprecated: Use PublishStrategyOptions instead enables Kaniko publish strategy cache",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"maven": {
								Description:         "Maven configuration used to build the Camel/Camel-Quarkus applications",
								MarkdownDescription: "Maven configuration used to build the Camel/Camel-Quarkus applications",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_secret": {
										Description:         "Deprecated: use CASecrets The Secret name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
										MarkdownDescription: "Deprecated: use CASecrets The Secret name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ca_secrets": {
										Description:         "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
										MarkdownDescription: "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cli_options": {
										Description:         "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
										MarkdownDescription: "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extension": {
										Description:         "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
										MarkdownDescription: "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"artifact_id": {
												Description:         "Maven Artifact",
												MarkdownDescription: "Maven Artifact",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"group_id": {
												Description:         "Maven Group",
												MarkdownDescription: "Maven Group",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"version": {
												Description:         "Maven Version",
												MarkdownDescription: "Maven Version",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"local_repository": {
										Description:         "The path of the local Maven repository.",
										MarkdownDescription: "The path of the local Maven repository.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "The Maven properties.",
										MarkdownDescription: "The Maven properties.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"settings": {
										Description:         "A reference to the ConfigMap or Secret key that contains the Maven settings.",
										MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven settings.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"settings_security": {
										Description:         "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
										MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret.",
												MarkdownDescription: "Selects a key of a secret.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"persistent_volume_claim": {
								Description:         "Deprecated: Use PublishStrategyOptions instead the Persistent Volume Claim used by Kaniko publish strategy, if cache is enabled",
								MarkdownDescription: "Deprecated: Use PublishStrategyOptions instead the Persistent Volume Claim used by Kaniko publish strategy, if cache is enabled",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"publish_strategy": {
								Description:         "the strategy to adopt for publishing an Integration base image",
								MarkdownDescription: "the strategy to adopt for publishing an Integration base image",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registry": {
								Description:         "the image registry used to push/pull Integration images",
								MarkdownDescription: "the image registry used to push/pull Integration images",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"address": {
										Description:         "the URI to access",
										MarkdownDescription: "the URI to access",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ca": {
										Description:         "the configmap which stores the Certificate Authority",
										MarkdownDescription: "the configmap which stores the Certificate Authority",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure": {
										Description:         "if the container registry is insecure (ie, http only)",
										MarkdownDescription: "if the container registry is insecure (ie, http only)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"organization": {
										Description:         "the registry organization",
										MarkdownDescription: "the registry organization",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "the secret where credentials are stored",
										MarkdownDescription: "the secret where credentials are stored",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime_provider": {
								Description:         "the runtime used. Likely Camel Quarkus (we used to have main runtime which has been discontinued since version 1.5)",
								MarkdownDescription: "the runtime used. Likely Camel Quarkus (we used to have main runtime which has been discontinued since version 1.5)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime_version": {
								Description:         "the Camel K Runtime dependency version",
								MarkdownDescription: "the Camel K Runtime dependency version",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "how much time to wait before time out the build process",
								MarkdownDescription: "how much time to wait before time out the build process",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster": {
						Description:         "what kind of cluster you're running (ie, plain Kubernetes or OpenShift)",
						MarkdownDescription: "what kind of cluster you're running (ie, plain Kubernetes or OpenShift)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"configuration": {
						Description:         "list of configuration properties to be attached to all the Integration/IntegrationKits built from this IntegrationPlatform",
						MarkdownDescription: "list of configuration properties to be attached to all the Integration/IntegrationKits built from this IntegrationPlatform",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"resource_key": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_mount_point": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_type": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
								MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
								MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kamelet": {
						Description:         "configuration to be executed to all Kamelets controlled by this IntegrationPlatform",
						MarkdownDescription: "configuration to be executed to all Kamelets controlled by this IntegrationPlatform",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"repositories": {
								Description:         "remote repository used to retrieve Kamelet catalog",
								MarkdownDescription: "remote repository used to retrieve Kamelet catalog",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"uri": {
										Description:         "the remote repository in the format github:ORG/REPO/PATH_TO_KAMELETS_FOLDER",
										MarkdownDescription: "the remote repository in the format github:ORG/REPO/PATH_TO_KAMELETS_FOLDER",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"profile": {
						Description:         "the profile you wish to use. It will apply certain traits which are required by the specific profile chosen. It usually relates the Cluster with the optional definition of special profiles (ie, Knative)",
						MarkdownDescription: "the profile you wish to use. It will apply certain traits which are required by the specific profile chosen. It usually relates the Cluster with the optional definition of special profiles (ie, Knative)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Deprecated: not used",
						MarkdownDescription: "Deprecated: not used",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"traits": {
						Description:         "list of traits to be executed for all the Integration/IntegrationKits built from this IntegrationPlatform",
						MarkdownDescription: "list of traits to be executed for all the Integration/IntegrationKits built from this IntegrationPlatform",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"threescale": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"addons": {
								Description:         "The extension point with addon traits",
								MarkdownDescription: "The extension point with addon traits",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"affinity": {
								Description:         "The configuration of Affinity trait",
								MarkdownDescription: "The configuration of Affinity trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_affinity_labels": {
										Description:         "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										MarkdownDescription: "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_affinity": {
										Description:         "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Always co-locates multiple replicas of the integration in the same node (default *false*).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_affinity_labels": {
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_anti_affinity": {
										Description:         "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Never co-locates multiple replicas of the integration in the same node (default *false*).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_anti_affinity_labels": {
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"builder": {
								Description:         "The configuration of Builder trait",
								MarkdownDescription: "The configuration of Builder trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verbose": {
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"camel": {
								Description:         "The configuration of Camel trait",
								MarkdownDescription: "The configuration of Camel trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"runtime_version": {
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"container": {
								Description:         "The configuration of Container trait",
								MarkdownDescription: "The configuration of Container trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically enable the trait",
										MarkdownDescription: "To automatically enable the trait",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expose": {
										Description:         "Can be used to enable/disable exposure via kubernetes Service.",
										MarkdownDescription: "Can be used to enable/disable exposure via kubernetes Service.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "The main container image",
										MarkdownDescription: "The main container image",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_policy": {
										Description:         "The pull policy: Always|Never|IfNotPresent",
										MarkdownDescription: "The pull policy: Always|Never|IfNotPresent",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"limit_cpu": {
										Description:         "The maximum amount of CPU required.",
										MarkdownDescription: "The maximum amount of CPU required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"limit_memory": {
										Description:         "The maximum amount of memory required.",
										MarkdownDescription: "The maximum amount of memory required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_initial_delay": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_period": {
										Description:         "How often to perform the probe. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "How often to perform the probe. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_scheme": {
										Description:         "Scheme to use when connecting. Defaults to HTTP. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Scheme to use when connecting. Defaults to HTTP. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_timeout": {
										Description:         "Number of seconds after which the probe times out. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after which the probe times out. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "The main container name. It's named 'integration' by default.",
										MarkdownDescription: "The main container name. It's named 'integration' by default.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "To configure a different port exposed by the container (default '8080').",
										MarkdownDescription: "To configure a different port exposed by the container (default '8080').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port_name": {
										Description:         "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										MarkdownDescription: "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"probes_enabled": {
										Description:         "DeprecatedProbesEnabled enable/disable probes on the container (default 'false'). Deprecated: replaced by the health trait.",
										MarkdownDescription: "DeprecatedProbesEnabled enable/disable probes on the container (default 'false'). Deprecated: replaced by the health trait.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_initial_delay": {
										Description:         "Number of seconds after the container has started before readiness probes are initiated. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after the container has started before readiness probes are initiated. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_period": {
										Description:         "How often to perform the probe. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "How often to perform the probe. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_scheme": {
										Description:         "Scheme to use when connecting. Defaults to HTTP. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Scheme to use when connecting. Defaults to HTTP. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_timeout": {
										Description:         "Number of seconds after which the probe times out. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after which the probe times out. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_cpu": {
										Description:         "The minimum amount of CPU required.",
										MarkdownDescription: "The minimum amount of CPU required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_memory": {
										Description:         "The minimum amount of memory required.",
										MarkdownDescription: "The minimum amount of memory required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_port": {
										Description:         "To configure under which service port the container port is to be exposed (default '80').",
										MarkdownDescription: "To configure under which service port the container port is to be exposed (default '80').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_port_name": {
										Description:         "To configure under which service port name the container port is to be exposed (default 'http').",
										MarkdownDescription: "To configure under which service port name the container port is to be exposed (default 'http').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cron": {
								Description:         "The configuration of Cron trait",
								MarkdownDescription: "The configuration of Cron trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"active_deadline_seconds": {
										Description:         "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auto": {
										Description:         "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										MarkdownDescription: "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"backoff_limit": {
										Description:         "Specifies the number of retries before marking the job failed. It defaults to 2.",
										MarkdownDescription: "Specifies the number of retries before marking the job failed. It defaults to 2.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"components": {
										Description:         "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										MarkdownDescription: "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"concurrency_policy": {
										Description:         "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										MarkdownDescription: "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Allow", "Forbid", "Replace"),
										},
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fallback": {
										Description:         "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										MarkdownDescription: "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"schedule": {
										Description:         "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										MarkdownDescription: "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"starting_deadline_seconds": {
										Description:         "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										MarkdownDescription: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dependencies": {
								Description:         "The configuration of Dependencies trait",
								MarkdownDescription: "The configuration of Dependencies trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"deployer": {
								Description:         "The configuration of Deployer trait",
								MarkdownDescription: "The configuration of Deployer trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										MarkdownDescription: "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("deployment", "cron-job", "knative-service"),
										},
									},

									"use_ssa": {
										Description:         "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										MarkdownDescription: "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"deployment": {
								Description:         "The configuration of Deployment trait",
								MarkdownDescription: "The configuration of Deployment trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"progress_deadline_seconds": {
										Description:         "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rolling_update_max_surge": {
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rolling_update_max_unavailable": {
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strategy": {
										Description:         "The deployment strategy to use to replace existing pods with new ones.",
										MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Recreate", "RollingUpdate"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"environment": {
								Description:         "The configuration of Environment trait",
								MarkdownDescription: "The configuration of Environment trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"container_meta": {
										Description:         "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										MarkdownDescription: "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_proxy": {
										Description:         "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										MarkdownDescription: "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"vars": {
										Description:         "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
										MarkdownDescription: "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"error_handler": {
								Description:         "The configuration of Error Handler trait",
								MarkdownDescription: "The configuration of Error Handler trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ref": {
										Description:         "The error handler ref name provided or found in application properties",
										MarkdownDescription: "The error handler ref name provided or found in application properties",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gc": {
								Description:         "The configuration of GC trait",
								MarkdownDescription: "The configuration of GC trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"discovery_cache": {
										Description:         "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										MarkdownDescription: "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("disabled", "disk", "memory"),
										},
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"health": {
								Description:         "The configuration of Health trait",
								MarkdownDescription: "The configuration of Health trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_failure_threshold": {
										Description:         "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_initial_delay": {
										Description:         "Number of seconds after the container has started before the liveness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the liveness probe is initiated.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_period": {
										Description:         "How often to perform the liveness probe.",
										MarkdownDescription: "How often to perform the liveness probe.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe_enabled": {
										Description:         "Configures the liveness probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the liveness probe for the integration container (default 'false').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_scheme": {
										Description:         "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the liveness probe (default 'HTTP').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_success_threshold": {
										Description:         "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_timeout": {
										Description:         "Number of seconds after which the liveness probe times out.",
										MarkdownDescription: "Number of seconds after which the liveness probe times out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_failure_threshold": {
										Description:         "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_initial_delay": {
										Description:         "Number of seconds after the container has started before the readiness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the readiness probe is initiated.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_period": {
										Description:         "How often to perform the readiness probe.",
										MarkdownDescription: "How often to perform the readiness probe.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe_enabled": {
										Description:         "Configures the readiness probe for the integration container (default 'true').",
										MarkdownDescription: "Configures the readiness probe for the integration container (default 'true').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_scheme": {
										Description:         "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the readiness probe (default 'HTTP').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_success_threshold": {
										Description:         "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_timeout": {
										Description:         "Number of seconds after which the readiness probe times out.",
										MarkdownDescription: "Number of seconds after which the readiness probe times out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": {
								Description:         "The configuration of Ingress trait",
								MarkdownDescription: "The configuration of Ingress trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										MarkdownDescription: "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "To configure the host exposed by the ingress.",
										MarkdownDescription: "To configure the host exposed by the ingress.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"istio": {
								Description:         "The configuration of Istio trait",
								MarkdownDescription: "The configuration of Istio trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow": {
										Description:         "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										MarkdownDescription: "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"inject": {
										Description:         "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										MarkdownDescription: "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jolokia": {
								Description:         "The configuration of Jolokia trait",
								MarkdownDescription: "The configuration of Jolokia trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_cert": {
										Description:         "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										MarkdownDescription: "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_principal": {
										Description:         "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										MarkdownDescription: "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"discovery_enabled": {
										Description:         "Listen for multicast requests (default 'false')",
										MarkdownDescription: "Listen for multicast requests (default 'false')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extended_client_check": {
										Description:         "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										MarkdownDescription: "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										MarkdownDescription: "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"options": {
										Description:         "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										MarkdownDescription: "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "The password used for authentication, applicable when the 'user' option is set.",
										MarkdownDescription: "The password used for authentication, applicable when the 'user' option is set.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "The Jolokia endpoint port (default '8778').",
										MarkdownDescription: "The Jolokia endpoint port (default '8778').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										MarkdownDescription: "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_ssl_client_authentication": {
										Description:         "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										MarkdownDescription: "Whether client certificates should be used for authentication (default 'true' for OpenShift).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "The user to be used for authentication",
										MarkdownDescription: "The user to be used for authentication",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jvm": {
								Description:         "The configuration of JVM trait",
								MarkdownDescription: "The configuration of JVM trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"classpath": {
										Description:         "Additional JVM classpath (use 'Linux' classpath separator)",
										MarkdownDescription: "Additional JVM classpath (use 'Linux' classpath separator)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug": {
										Description:         "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										MarkdownDescription: "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug_address": {
										Description:         "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										MarkdownDescription: "Transport address at which to listen for the newly launched JVM (default '*:5005')",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug_suspend": {
										Description:         "Suspends the target JVM immediately before the main class is loaded",
										MarkdownDescription: "Suspends the target JVM immediately before the main class is loaded",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"options": {
										Description:         "A list of JVM options",
										MarkdownDescription: "A list of JVM options",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"print_command": {
										Description:         "Prints the command used the start the JVM in the container logs (default 'true')",
										MarkdownDescription: "Prints the command used the start the JVM in the container logs (default 'true')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kamelets": {
								Description:         "The configuration of Kamelets trait",
								MarkdownDescription: "The configuration of Kamelets trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										MarkdownDescription: "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"list": {
										Description:         "Comma separated list of Kamelet names to load into the current integration",
										MarkdownDescription: "Comma separated list of Kamelet names to load into the current integration",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keda": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"knative": {
								Description:         "The configuration of Knative trait",
								MarkdownDescription: "The configuration of Knative trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Enable automatic discovery of all trait properties.",
										MarkdownDescription: "Enable automatic discovery of all trait properties.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"channel_sinks": {
										Description:         "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"channel_sources": {
										Description:         "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config": {
										Description:         "Can be used to inject a Knative complete configuration in JSON format.",
										MarkdownDescription: "Can be used to inject a Knative complete configuration in JSON format.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint_sinks": {
										Description:         "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										MarkdownDescription: "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint_sources": {
										Description:         "List of channels used as source of integration routes.",
										MarkdownDescription: "List of channels used as source of integration routes.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"event_sinks": {
										Description:         "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										MarkdownDescription: "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"event_sources": {
										Description:         "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										MarkdownDescription: "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter_source_channels": {
										Description:         "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										MarkdownDescription: "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sink_binding": {
										Description:         "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										MarkdownDescription: "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"knative_service": {
								Description:         "The configuration of Knative Service trait",
								MarkdownDescription: "The configuration of Knative Service trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										MarkdownDescription: "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"autoscaling_metric": {
										Description:         "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"autoscaling_target": {
										Description:         "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"class": {
										Description:         "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("kpa.autoscaling.knative.dev", "hpa.autoscaling.knative.dev"),
										},
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_scale": {
										Description:         "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_scale": {
										Description:         "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rollout_duration": {
										Description:         "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										MarkdownDescription: "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"visibility": {
										Description:         "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("cluster-local"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "The configuration of Logging trait",
								MarkdownDescription: "The configuration of Logging trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"color": {
										Description:         "Colorize the log output",
										MarkdownDescription: "Colorize the log output",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"format": {
										Description:         "Logs message format",
										MarkdownDescription: "Logs message format",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json": {
										Description:         "Output the logs in JSON",
										MarkdownDescription: "Output the logs in JSON",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_pretty_print": {
										Description:         "Enable 'pretty printing' of the JSON logs",
										MarkdownDescription: "Enable 'pretty printing' of the JSON logs",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"level": {
										Description:         "Adjust the logging level (defaults to INFO)",
										MarkdownDescription: "Adjust the logging level (defaults to INFO)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FATAL", "WARN", "INFO", "DEBUG", "TRACE"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"master": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount": {
								Description:         "The configuration of Mount trait",
								MarkdownDescription: "The configuration of Mount trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configs": {
										Description:         "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										MarkdownDescription: "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										MarkdownDescription: "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volumes": {
										Description:         "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
										MarkdownDescription: "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"openapi": {
								Description:         "The configuration of OpenAPI trait",
								MarkdownDescription: "The configuration of OpenAPI trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configmaps": {
										Description:         "The configmaps holding the spec of the OpenAPI",
										MarkdownDescription: "The configmaps holding the spec of the OpenAPI",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"owner": {
								Description:         "The configuration of Owner trait",
								MarkdownDescription: "The configuration of Owner trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_annotations": {
										Description:         "The set of annotations to be transferred",
										MarkdownDescription: "The set of annotations to be transferred",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_labels": {
										Description:         "The set of labels to be transferred",
										MarkdownDescription: "The set of labels to be transferred",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pdb": {
								Description:         "The configuration of PDB trait",
								MarkdownDescription: "The configuration of PDB trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_unavailable": {
										Description:         "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_available": {
										Description:         "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"platform": {
								Description:         "The configuration of Platform trait",
								MarkdownDescription: "The configuration of Platform trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										MarkdownDescription: "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"create_default": {
										Description:         "To create a default (empty) platform when the platform is missing.",
										MarkdownDescription: "To create a default (empty) platform when the platform is missing.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"global": {
										Description:         "Indicates if the platform should be created globally in the case of global operator (default true).",
										MarkdownDescription: "Indicates if the platform should be created globally in the case of global operator (default true).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod": {
								Description:         "The configuration of Pod trait",
								MarkdownDescription: "The configuration of Pod trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prometheus": {
								Description:         "The configuration of Prometheus trait",
								MarkdownDescription: "The configuration of Prometheus trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_monitor": {
										Description:         "Whether a 'PodMonitor' resource is created (default 'true').",
										MarkdownDescription: "Whether a 'PodMonitor' resource is created (default 'true').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_monitor_labels": {
										Description:         "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
										MarkdownDescription: "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pull_secret": {
								Description:         "The configuration of Pull Secret trait",
								MarkdownDescription: "The configuration of Pull Secret trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										MarkdownDescription: "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_puller_delegation": {
										Description:         "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										MarkdownDescription: "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										MarkdownDescription: "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"quarkus": {
								Description:         "The configuration of Quarkus trait",
								MarkdownDescription: "The configuration of Quarkus trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"package_types": {
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registry": {
								Description:         "The configuration of Registry trait",
								MarkdownDescription: "The configuration of Registry trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"route": {
								Description:         "The configuration of Route trait",
								MarkdownDescription: "The configuration of Route trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "To configure the host exposed by the route.",
										MarkdownDescription: "To configure the host exposed by the route.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_ca_certificate": {
										Description:         "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_ca_certificate_secret": {
										Description:         "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_certificate": {
										Description:         "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_certificate_secret": {
										Description:         "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_destination_ca_certificate": {
										Description:         "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_destination_ca_certificate_secret": {
										Description:         "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_insecure_edge_termination_policy": {
										Description:         "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("None", "Allow", "Redirect"),
										},
									},

									"tls_key": {
										Description:         "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_key_secret": {
										Description:         "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_termination": {
										Description:         "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("edge", "reencrypt", "passthrough"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service": {
								Description:         "The configuration of Service trait",
								MarkdownDescription: "The configuration of Service trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically detect from the code if a Service needs to be created.",
										MarkdownDescription: "To automatically detect from the code if a Service needs to be created.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_port": {
										Description:         "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										MarkdownDescription: "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										MarkdownDescription: "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_binding": {
								Description:         "The configuration of Service Binding trait",
								MarkdownDescription: "The configuration of Service Binding trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"services": {
										Description:         "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
										MarkdownDescription: "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"strimzi": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"toleration": {
								Description:         "The configuration of Toleration trait",
								MarkdownDescription: "The configuration of Toleration trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"taints": {
										Description:         "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
										MarkdownDescription: "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_integration_platform_v1")

	var state CamelApacheOrgIntegrationPlatformV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationPlatformV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("IntegrationPlatform")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_integration_platform_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_integration_platform_v1")

	var state CamelApacheOrgIntegrationPlatformV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationPlatformV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("IntegrationPlatform")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CamelApacheOrgIntegrationPlatformV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_integration_platform_v1")
	// NO-OP: Terraform removes the state automatically for us
}
