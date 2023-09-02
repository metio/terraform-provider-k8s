/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_grafana_com_v1

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
	_ datasource.DataSource              = &ConfigGrafanaComProjectConfigV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigGrafanaComProjectConfigV1DataSource{}
)

func NewConfigGrafanaComProjectConfigV1DataSource() datasource.DataSource {
	return &ConfigGrafanaComProjectConfigV1DataSource{}
}

type ConfigGrafanaComProjectConfigV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ConfigGrafanaComProjectConfigV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	CacheNamespace *string `tfsdk:"cache_namespace" json:"cacheNamespace,omitempty"`
	Controller     *struct {
		CacheSyncTimeout     *int64             `tfsdk:"cache_sync_timeout" json:"cacheSyncTimeout,omitempty"`
		GroupKindConcurrency *map[string]string `tfsdk:"group_kind_concurrency" json:"groupKindConcurrency,omitempty"`
	} `tfsdk:"controller" json:"controller,omitempty"`
	FeatureFlags *struct {
		EnableAlertingRuleWebhook     *bool `tfsdk:"enable_alerting_rule_webhook" json:"enableAlertingRuleWebhook,omitempty"`
		EnableCertSigningService      *bool `tfsdk:"enable_cert_signing_service" json:"enableCertSigningService,omitempty"`
		EnableGrafanaLabsStats        *bool `tfsdk:"enable_grafana_labs_stats" json:"enableGrafanaLabsStats,omitempty"`
		EnableLokiStackAlerts         *bool `tfsdk:"enable_loki_stack_alerts" json:"enableLokiStackAlerts,omitempty"`
		EnableLokiStackGateway        *bool `tfsdk:"enable_loki_stack_gateway" json:"enableLokiStackGateway,omitempty"`
		EnableLokiStackGatewayRoute   *bool `tfsdk:"enable_loki_stack_gateway_route" json:"enableLokiStackGatewayRoute,omitempty"`
		EnableRecordingRuleWebhook    *bool `tfsdk:"enable_recording_rule_webhook" json:"enableRecordingRuleWebhook,omitempty"`
		EnableRulerConfigWebhook      *bool `tfsdk:"enable_ruler_config_webhook" json:"enableRulerConfigWebhook,omitempty"`
		EnableServiceMonitors         *bool `tfsdk:"enable_service_monitors" json:"enableServiceMonitors,omitempty"`
		EnableTlsGrpcServices         *bool `tfsdk:"enable_tls_grpc_services" json:"enableTlsGrpcServices,omitempty"`
		EnableTlsHttpServices         *bool `tfsdk:"enable_tls_http_services" json:"enableTlsHttpServices,omitempty"`
		EnableTlsServiceMonitorConfig *bool `tfsdk:"enable_tls_service_monitor_config" json:"enableTlsServiceMonitorConfig,omitempty"`
	} `tfsdk:"feature_flags" json:"featureFlags,omitempty"`
	GracefulShutDown *string `tfsdk:"graceful_shut_down" json:"gracefulShutDown,omitempty"`
	Health           *struct {
		HealthProbeBindAddress *string `tfsdk:"health_probe_bind_address" json:"healthProbeBindAddress,omitempty"`
		LivenessEndpointName   *string `tfsdk:"liveness_endpoint_name" json:"livenessEndpointName,omitempty"`
		ReadinessEndpointName  *string `tfsdk:"readiness_endpoint_name" json:"readinessEndpointName,omitempty"`
	} `tfsdk:"health" json:"health,omitempty"`
	LeaderElection *struct {
		LeaderElect       *bool   `tfsdk:"leader_elect" json:"leaderElect,omitempty"`
		LeaseDuration     *string `tfsdk:"lease_duration" json:"leaseDuration,omitempty"`
		RenewDeadline     *string `tfsdk:"renew_deadline" json:"renewDeadline,omitempty"`
		ResourceLock      *string `tfsdk:"resource_lock" json:"resourceLock,omitempty"`
		ResourceName      *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
		ResourceNamespace *string `tfsdk:"resource_namespace" json:"resourceNamespace,omitempty"`
		RetryPeriod       *string `tfsdk:"retry_period" json:"retryPeriod,omitempty"`
	} `tfsdk:"leader_election" json:"leaderElection,omitempty"`
	Metrics *struct {
		BindAddress *string `tfsdk:"bind_address" json:"bindAddress,omitempty"`
	} `tfsdk:"metrics" json:"metrics,omitempty"`
	SyncPeriod *string `tfsdk:"sync_period" json:"syncPeriod,omitempty"`
	Webhook    *struct {
		CertDir *string `tfsdk:"cert_dir" json:"certDir,omitempty"`
		Host    *string `tfsdk:"host" json:"host,omitempty"`
		Port    *int64  `tfsdk:"port" json:"port,omitempty"`
	} `tfsdk:"webhook" json:"webhook,omitempty"`
}

func (r *ConfigGrafanaComProjectConfigV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_grafana_com_project_config_v1"
}

func (r *ConfigGrafanaComProjectConfigV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProjectConfig is the Schema for the projectconfigs API",
		MarkdownDescription: "ProjectConfig is the Schema for the projectconfigs API",
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

			"cache_namespace": schema.StringAttribute{
				Description:         "CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces  Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.",
				MarkdownDescription: "CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces  Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node).  For namespaced resources the cache will only hold objects from the desired namespace.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"controller": schema.SingleNestedAttribute{
				Description:         "Controller contains global configuration options for controllers registered within this manager.",
				MarkdownDescription: "Controller contains global configuration options for controllers registered within this manager.",
				Attributes: map[string]schema.Attribute{
					"cache_sync_timeout": schema.Int64Attribute{
						Description:         "CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set.",
						MarkdownDescription: "CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"group_kind_concurrency": schema.MapAttribute{
						Description:         "GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller.  When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified.  The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be 'ReplicaSet.apps'.",
						MarkdownDescription: "GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller.  When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified.  The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be 'ReplicaSet.apps'.",
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

			"feature_flags": schema.SingleNestedAttribute{
				Description:         "FeatureFlags is a set of operator feature flags.",
				MarkdownDescription: "FeatureFlags is a set of operator feature flags.",
				Attributes: map[string]schema.Attribute{
					"enable_alerting_rule_webhook": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_cert_signing_service": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_grafana_labs_stats": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_loki_stack_alerts": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_loki_stack_gateway": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_loki_stack_gateway_route": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_recording_rule_webhook": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_ruler_config_webhook": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_service_monitors": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_tls_grpc_services": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_tls_http_services": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_tls_service_monitor_config": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"graceful_shut_down": schema.StringAttribute{
				Description:         "GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost.",
				MarkdownDescription: "GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"health": schema.SingleNestedAttribute{
				Description:         "Health contains the controller health configuration",
				MarkdownDescription: "Health contains the controller health configuration",
				Attributes: map[string]schema.Attribute{
					"health_probe_bind_address": schema.StringAttribute{
						Description:         "HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes",
						MarkdownDescription: "HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"liveness_endpoint_name": schema.StringAttribute{
						Description:         "LivenessEndpointName, defaults to 'healthz'",
						MarkdownDescription: "LivenessEndpointName, defaults to 'healthz'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"readiness_endpoint_name": schema.StringAttribute{
						Description:         "ReadinessEndpointName, defaults to 'readyz'",
						MarkdownDescription: "ReadinessEndpointName, defaults to 'readyz'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"leader_election": schema.SingleNestedAttribute{
				Description:         "LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election",
				MarkdownDescription: "LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election",
				Attributes: map[string]schema.Attribute{
					"leader_elect": schema.BoolAttribute{
						Description:         "leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability.",
						MarkdownDescription: "leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"lease_duration": schema.StringAttribute{
						Description:         "leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.",
						MarkdownDescription: "leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"renew_deadline": schema.StringAttribute{
						Description:         "renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled.",
						MarkdownDescription: "renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_lock": schema.StringAttribute{
						Description:         "resourceLock indicates the resource object type that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceLock indicates the resource object type that will be used to lock during leader election cycles.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_name": schema.StringAttribute{
						Description:         "resourceName indicates the name of resource object that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceName indicates the name of resource object that will be used to lock during leader election cycles.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_namespace": schema.StringAttribute{
						Description:         "resourceName indicates the namespace of resource object that will be used to lock during leader election cycles.",
						MarkdownDescription: "resourceName indicates the namespace of resource object that will be used to lock during leader election cycles.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"retry_period": schema.StringAttribute{
						Description:         "retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled.",
						MarkdownDescription: "retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"metrics": schema.SingleNestedAttribute{
				Description:         "Metrics contains thw controller metrics configuration",
				MarkdownDescription: "Metrics contains thw controller metrics configuration",
				Attributes: map[string]schema.Attribute{
					"bind_address": schema.StringAttribute{
						Description:         "BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to '0' to disable the metrics serving.",
						MarkdownDescription: "BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to '0' to disable the metrics serving.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"sync_period": schema.StringAttribute{
				Description:         "SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously.",
				MarkdownDescription: "SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"webhook": schema.SingleNestedAttribute{
				Description:         "Webhook contains the controllers webhook configuration",
				MarkdownDescription: "Webhook contains the controllers webhook configuration",
				Attributes: map[string]schema.Attribute{
					"cert_dir": schema.StringAttribute{
						Description:         "CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively.",
						MarkdownDescription: "CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"host": schema.StringAttribute{
						Description:         "Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host.",
						MarkdownDescription: "Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"port": schema.Int64Attribute{
						Description:         "Port is the port that the webhook server serves at. It is used to set webhook.Server.Port.",
						MarkdownDescription: "Port is the port that the webhook server serves at. It is used to set webhook.Server.Port.",
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
	}
}

func (r *ConfigGrafanaComProjectConfigV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ConfigGrafanaComProjectConfigV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_config_grafana_com_project_config_v1")

	var data ConfigGrafanaComProjectConfigV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "config.grafana.com", Version: "v1", Resource: "ProjectConfig"}).
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

	var readResponse ConfigGrafanaComProjectConfigV1DataSourceData
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
	data.ApiVersion = pointer.String("config.grafana.com/v1")
	data.Kind = pointer.String("ProjectConfig")
	data.Metadata = readResponse.Metadata
	data.CacheNamespace = readResponse.CacheNamespace
	data.Controller = readResponse.Controller
	data.FeatureFlags = readResponse.FeatureFlags
	data.GracefulShutDown = readResponse.GracefulShutDown
	data.Health = readResponse.Health
	data.LeaderElection = readResponse.LeaderElection
	data.Metrics = readResponse.Metrics
	data.SyncPeriod = readResponse.SyncPeriod
	data.Webhook = readResponse.Webhook

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
