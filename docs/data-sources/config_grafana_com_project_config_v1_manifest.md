---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_config_grafana_com_project_config_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "config.grafana.com"
description: |-
  ProjectConfig is the Schema for the projectconfigs API
---

# k8s_config_grafana_com_project_config_v1_manifest (Data Source)

ProjectConfig is the Schema for the projectconfigs API

## Example Usage

```terraform
data "k8s_config_grafana_com_project_config_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `cache_namespace` (String) CacheNamespace if specified restricts the manager's cache to watch objects in the desired namespace Defaults to all namespaces Note: If a namespace is specified, controllers can still Watch for a cluster-scoped resource (e.g Node). For namespaced resources the cache will only hold objects from the desired namespace.
- `controller` (Attributes) Controller contains global configuration options for controllers registered within this manager. (see [below for nested schema](#nestedatt--controller))
- `feature_flags` (Attributes) FeatureFlags is a set of operator feature flags. (see [below for nested schema](#nestedatt--feature_flags))
- `graceful_shut_down` (String) GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop. To disable graceful shutdown, set to time.Duration(0) To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1) The graceful shutdown is skipped for safety reasons in case the leader election lease is lost.
- `health` (Attributes) Health contains the controller health configuration (see [below for nested schema](#nestedatt--health))
- `leader_election` (Attributes) LeaderElection is the LeaderElection config to be used when configuring the manager.Manager leader election (see [below for nested schema](#nestedatt--leader_election))
- `metrics` (Attributes) Metrics contains thw controller metrics configuration (see [below for nested schema](#nestedatt--metrics))
- `sync_period` (String) SyncPeriod determines the minimum frequency at which watched resources are reconciled. A lower period will correct entropy more quickly, but reduce responsiveness to change if there are many watched resources. Change this value only if you know what you are doing. Defaults to 10 hours if unset. there will a 10 percent jitter between the SyncPeriod of all controllers so that all controllers will not send list requests simultaneously.
- `webhook` (Attributes) Webhook contains the controllers webhook configuration (see [below for nested schema](#nestedatt--webhook))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--controller"></a>
### Nested Schema for `controller`

Optional:

- `cache_sync_timeout` (Number) CacheSyncTimeout refers to the time limit set to wait for syncing caches. Defaults to 2 minutes if not set.
- `group_kind_concurrency` (Map of String) GroupKindConcurrency is a map from a Kind to the number of concurrent reconciliation allowed for that controller. When a controller is registered within this manager using the builder utilities, users have to specify the type the controller reconciles in the For(...) call. If the object's kind passed matches one of the keys in this map, the concurrency for that controller is set to the number specified. The key is expected to be consistent in form with GroupKind.String(), e.g. ReplicaSet in apps group (regardless of version) would be 'ReplicaSet.apps'.


<a id="nestedatt--feature_flags"></a>
### Nested Schema for `feature_flags`

Optional:

- `enable_alerting_rule_webhook` (Boolean)
- `enable_cert_signing_service` (Boolean)
- `enable_grafana_labs_stats` (Boolean)
- `enable_loki_stack_alerts` (Boolean)
- `enable_loki_stack_gateway` (Boolean)
- `enable_loki_stack_gateway_route` (Boolean)
- `enable_recording_rule_webhook` (Boolean)
- `enable_ruler_config_webhook` (Boolean)
- `enable_service_monitors` (Boolean)
- `enable_tls_grpc_services` (Boolean)
- `enable_tls_http_services` (Boolean)
- `enable_tls_service_monitor_config` (Boolean)


<a id="nestedatt--health"></a>
### Nested Schema for `health`

Optional:

- `health_probe_bind_address` (String) HealthProbeBindAddress is the TCP address that the controller should bind to for serving health probes
- `liveness_endpoint_name` (String) LivenessEndpointName, defaults to 'healthz'
- `readiness_endpoint_name` (String) ReadinessEndpointName, defaults to 'readyz'


<a id="nestedatt--leader_election"></a>
### Nested Schema for `leader_election`

Required:

- `leader_elect` (Boolean) leaderElect enables a leader election client to gain leadership before executing the main loop. Enable this when running replicated components for high availability.
- `lease_duration` (String) leaseDuration is the duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.
- `renew_deadline` (String) renewDeadline is the interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled.
- `resource_lock` (String) resourceLock indicates the resource object type that will be used to lock during leader election cycles.
- `resource_name` (String) resourceName indicates the name of resource object that will be used to lock during leader election cycles.
- `resource_namespace` (String) resourceName indicates the namespace of resource object that will be used to lock during leader election cycles.
- `retry_period` (String) retryPeriod is the duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled.


<a id="nestedatt--metrics"></a>
### Nested Schema for `metrics`

Optional:

- `bind_address` (String) BindAddress is the TCP address that the controller should bind to for serving prometheus metrics. It can be set to '0' to disable the metrics serving.


<a id="nestedatt--webhook"></a>
### Nested Schema for `webhook`

Optional:

- `cert_dir` (String) CertDir is the directory that contains the server key and certificate. if not set, webhook server would look up the server key and certificate in {TempDir}/k8s-webhook-server/serving-certs. The server key and certificate must be named tls.key and tls.crt, respectively.
- `host` (String) Host is the hostname that the webhook server binds to. It is used to set webhook.Server.Host.
- `port` (Number) Port is the port that the webhook server serves at. It is used to set webhook.Server.Port.
