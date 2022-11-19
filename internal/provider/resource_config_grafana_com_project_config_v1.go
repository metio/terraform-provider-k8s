/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type ConfigGrafanaComProjectConfigV1Resource struct{}

var (
	_ resource.Resource = (*ConfigGrafanaComProjectConfigV1Resource)(nil)
)

type ConfigGrafanaComProjectConfigV1TerraformModel struct {
	Id               types.Int64  `tfsdk:"id"`
	YAML             types.String `tfsdk:"yaml"`
	ApiVersion       types.String `tfsdk:"api_version"`
	Kind             types.String `tfsdk:"kind"`
	Metadata         types.Object `tfsdk:"metadata"`
	CacheNamespace   types.String `tfsdk:"cache_namespace"`
	Controller       types.Object `tfsdk:"controller"`
	FeatureFlags     types.Object `tfsdk:"feature_flags"`
	GracefulShutDown types.String `tfsdk:"graceful_shut_down"`
	Health           types.Object `tfsdk:"health"`
	LeaderElection   types.Object `tfsdk:"leader_election"`
	Metrics          types.Object `tfsdk:"metrics"`
	SyncPeriod       types.String `tfsdk:"sync_period"`
	Webhook          types.Object `tfsdk:"webhook"`
}

type ConfigGrafanaComProjectConfigV1GoModel struct {
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

	CacheNamespace *string `tfsdk:"cache_namespace" yaml:"cacheNamespace,omitempty"`

	Controller *struct {
		CacheSyncTimeout *int64 `tfsdk:"cache_sync_timeout" yaml:"cacheSyncTimeout,omitempty"`

		GroupKindConcurrency *map[string]string `tfsdk:"group_kind_concurrency" yaml:"groupKindConcurrency,omitempty"`
	} `tfsdk:"controller" yaml:"controller,omitempty"`

	FeatureFlags *struct {
		EnableAlertingRuleWebhook *bool `tfsdk:"enable_alerting_rule_webhook" yaml:"enableAlertingRuleWebhook,omitempty"`

		EnableCertSigningService *bool `tfsdk:"enable_cert_signing_service" yaml:"enableCertSigningService,omitempty"`

		EnableGrafanaLabsStats *bool `tfsdk:"enable_grafana_labs_stats" yaml:"enableGrafanaLabsStats,omitempty"`

		EnableLokiStackAlerts *bool `tfsdk:"enable_loki_stack_alerts" yaml:"enableLokiStackAlerts,omitempty"`

		EnableLokiStackGateway *bool `tfsdk:"enable_loki_stack_gateway" yaml:"enableLokiStackGateway,omitempty"`

		EnableLokiStackGatewayRoute *bool `tfsdk:"enable_loki_stack_gateway_route" yaml:"enableLokiStackGatewayRoute,omitempty"`

		EnableRecordingRuleWebhook *bool `tfsdk:"enable_recording_rule_webhook" yaml:"enableRecordingRuleWebhook,omitempty"`

		EnableServiceMonitors *bool `tfsdk:"enable_service_monitors" yaml:"enableServiceMonitors,omitempty"`

		EnableTlsGrpcServices *bool `tfsdk:"enable_tls_grpc_services" yaml:"enableTlsGrpcServices,omitempty"`

		EnableTlsHttpServices *bool `tfsdk:"enable_tls_http_services" yaml:"enableTlsHttpServices,omitempty"`

		EnableTlsServiceMonitorConfig *bool `tfsdk:"enable_tls_service_monitor_config" yaml:"enableTlsServiceMonitorConfig,omitempty"`
	} `tfsdk:"feature_flags" yaml:"featureFlags,omitempty"`

	GracefulShutDown *string `tfsdk:"graceful_shut_down" yaml:"gracefulShutDown,omitempty"`

	Health *struct {
		HealthProbeBindAddress *string `tfsdk:"health_probe_bind_address" yaml:"healthProbeBindAddress,omitempty"`

		LivenessEndpointName *string `tfsdk:"liveness_endpoint_name" yaml:"livenessEndpointName,omitempty"`

		ReadinessEndpointName *string `tfsdk:"readiness_endpoint_name" yaml:"readinessEndpointName,omitempty"`
	} `tfsdk:"health" yaml:"health,omitempty"`

	LeaderElection *struct {
		LeaderElect *bool `tfsdk:"leader_elect" yaml:"leaderElect,omitempty"`

		LeaseDuration *string `tfsdk:"lease_duration" yaml:"leaseDuration,omitempty"`

		RenewDeadline *string `tfsdk:"renew_deadline" yaml:"renewDeadline,omitempty"`

		ResourceLock *string `tfsdk:"resource_lock" yaml:"resourceLock,omitempty"`

		ResourceName *string `tfsdk:"resource_name" yaml:"resourceName,omitempty"`

		ResourceNamespace *string `tfsdk:"resource_namespace" yaml:"resourceNamespace,omitempty"`

		RetryPeriod *string `tfsdk:"retry_period" yaml:"retryPeriod,omitempty"`
	} `tfsdk:"leader_election" yaml:"leaderElection,omitempty"`

	Metrics *struct {
		BindAddress *string `tfsdk:"bind_address" yaml:"bindAddress,omitempty"`
	} `tfsdk:"metrics" yaml:"metrics,omitempty"`

	SyncPeriod *string `tfsdk:"sync_period" yaml:"syncPeriod,omitempty"`

	Webhook *struct {
		CertDir *string `tfsdk:"cert_dir" yaml:"certDir,omitempty"`

		Host *string `tfsdk:"host" yaml:"host,omitempty"`

		Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
	} `tfsdk:"webhook" yaml:"webhook,omitempty"`
}

func NewConfigGrafanaComProjectConfigV1Resource() resource.Resource {
	return &ConfigGrafanaComProjectConfigV1Resource{}
}

func (r *ConfigGrafanaComProjectConfigV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_grafana_com_project_config_v1"
}

func (r *ConfigGrafanaComProjectConfigV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ProjectConfig is the Schema for the projectconfigs API",
		MarkdownDescription: "ProjectConfig is the Schema for the projectconfigs API",
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

			"cache_namespace": {
				Description:         "CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces  Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.",
				MarkdownDescription: "CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces  Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"controller": {
				Description:         "Controller contains global configuration options for controllers registered within this manager.",
				MarkdownDescription: "Controller contains global configuration options for controllers registered within this manager.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cache_sync_timeout": {
						Description:         "CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set.",
						MarkdownDescription: "CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"group_kind_concurrency": {
						Description:         "GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller.  When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified.  The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be 'ReplicaSet.apps'.",
						MarkdownDescription: "GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller.  When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified.  The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be 'ReplicaSet.apps'.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},

			"feature_flags": {
				Description:         "FeatureFlags is a set of operator feature flags.",
				MarkdownDescription: "FeatureFlags is a set of operator feature flags.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"enable_alerting_rule_webhook": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_cert_signing_service": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_grafana_labs_stats": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_loki_stack_alerts": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_loki_stack_gateway": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_loki_stack_gateway_route": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_recording_rule_webhook": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_service_monitors": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_tls_grpc_services": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_tls_http_services": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_tls_service_monitor_config": {
						Description:         "",
						MarkdownDescription: "",

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

			"graceful_shut_down": {
				Description:         "GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost.",
				MarkdownDescription: "GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"health": {
				Description:         "Health contains the controller health configuration",
				MarkdownDescription: "Health contains the controller health configuration",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"health_probe_bind_address": {
						Description:         "HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes",
						MarkdownDescription: "HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"liveness_endpoint_name": {
						Description:         "LivenessEndpointName, defaults to 'healthz'",
						MarkdownDescription: "LivenessEndpointName, defaults to 'healthz'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"readiness_endpoint_name": {
						Description:         "ReadinessEndpointName, defaults to 'readyz'",
						MarkdownDescription: "ReadinessEndpointName, defaults to 'readyz'",

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

			"leader_election": {
				Description:         "LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election",
				MarkdownDescription: "LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"leader_elect": {
						Description:         "leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability.",
						MarkdownDescription: "leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability.",

						Type: types.BoolType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"lease_duration": {
						Description:         "leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.",
						MarkdownDescription: "leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"renew_deadline": {
						Description:         "renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled.",
						MarkdownDescription: "renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource_lock": {
						Description:         "resourceLock indicates the resource object type that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceLock indicates the resource object type that will be used to lock during leader election cycles.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource_name": {
						Description:         "resourceName indicates the name of resource object that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceName indicates the name of resource object that will be used to lock during leader election cycles.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource_namespace": {
						Description:         "resourceName indicates the namespace of resource object that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceName indicates the namespace of resource object that will be used to lock during leader election cycles.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"retry_period": {
						Description:         "retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled.",
						MarkdownDescription: "retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled.",

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

			"metrics": {
				Description:         "Metrics contains thw controller metrics configuration",
				MarkdownDescription: "Metrics contains thw controller metrics configuration",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"bind_address": {
						Description:         "BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to '0' to disable the metrics serving.",
						MarkdownDescription: "BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to '0' to disable the metrics serving.",

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

			"sync_period": {
				Description:         "SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously.",
				MarkdownDescription: "SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"webhook": {
				Description:         "Webhook contains the controllers webhook configuration",
				MarkdownDescription: "Webhook contains the controllers webhook configuration",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cert_dir": {
						Description:         "CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively.",
						MarkdownDescription: "CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": {
						Description:         "Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host.",
						MarkdownDescription: "Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": {
						Description:         "Port is the port that the webhook server serves at. It is used to set webhook.Server.Port.",
						MarkdownDescription: "Port is the port that the webhook server serves at. It is used to set webhook.Server.Port.",

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
		},
	}, nil
}

func (r *ConfigGrafanaComProjectConfigV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_config_grafana_com_project_config_v1")

	var state ConfigGrafanaComProjectConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ConfigGrafanaComProjectConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("config.grafana.com/v1")
	goModel.Kind = utilities.Ptr("ProjectConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ConfigGrafanaComProjectConfigV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_grafana_com_project_config_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *ConfigGrafanaComProjectConfigV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_config_grafana_com_project_config_v1")

	var state ConfigGrafanaComProjectConfigV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ConfigGrafanaComProjectConfigV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("config.grafana.com/v1")
	goModel.Kind = utilities.Ptr("ProjectConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ConfigGrafanaComProjectConfigV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_config_grafana_com_project_config_v1")
	// NO-OP: Terraform removes the state automatically for us
}
