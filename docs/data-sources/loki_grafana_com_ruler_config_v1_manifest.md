---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_loki_grafana_com_ruler_config_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "loki.grafana.com"
description: |-
  RulerConfig is the Schema for the rulerconfigs API
---

# k8s_loki_grafana_com_ruler_config_v1_manifest (Data Source)

RulerConfig is the Schema for the rulerconfigs API

## Example Usage

```terraform
data "k8s_loki_grafana_com_ruler_config_v1_manifest" "example" {
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

- `spec` (Attributes) RulerConfigSpec defines the desired state of Ruler (see [below for nested schema](#nestedatt--spec))

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


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `alertmanager` (Attributes) Defines alert manager configuration to notify on firing alerts. (see [below for nested schema](#nestedatt--spec--alertmanager))
- `evaluation_interval` (String) Interval on how frequently to evaluate rules.
- `overrides` (Attributes) Overrides defines the config overrides to be applied per-tenant. (see [below for nested schema](#nestedatt--spec--overrides))
- `poll_interval` (String) Interval on how frequently to poll for new rule definitions.
- `remote_write` (Attributes) Defines a remote write endpoint to write recording rule metrics. (see [below for nested schema](#nestedatt--spec--remote_write))

<a id="nestedatt--spec--alertmanager"></a>
### Nested Schema for `spec.alertmanager`

Required:

- `endpoints` (List of String) List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated asa separate group in the configuration. Multiple Alertmanagers in HA per group can besupported by using DNS resolution (See EnableDNSDiscovery).

Optional:

- `client` (Attributes) Client configuration for reaching the alertmanager endpoint. (see [below for nested schema](#nestedatt--spec--alertmanager--client))
- `discovery` (Attributes) Defines the configuration for DNS-based discovery of AlertManager hosts. (see [below for nested schema](#nestedatt--spec--alertmanager--discovery))
- `enable_v2` (Boolean) If enabled, then requests to Alertmanager use the v2 API.
- `external_labels` (Map of String) Additional labels to add to all alerts.
- `external_url` (String) URL for alerts return path.
- `notification_queue` (Attributes) Defines the configuration for the notification queue to AlertManager hosts. (see [below for nested schema](#nestedatt--spec--alertmanager--notification_queue))
- `relabel_configs` (Attributes List) List of alert relabel configurations. (see [below for nested schema](#nestedatt--spec--alertmanager--relabel_configs))

<a id="nestedatt--spec--alertmanager--client"></a>
### Nested Schema for `spec.alertmanager.client`

Optional:

- `basic_auth` (Attributes) Basic authentication configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--alertmanager--client--basic_auth))
- `header_auth` (Attributes) Header authentication configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--alertmanager--client--header_auth))
- `tls` (Attributes) TLS configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--alertmanager--client--tls))

<a id="nestedatt--spec--alertmanager--client--basic_auth"></a>
### Nested Schema for `spec.alertmanager.client.basic_auth`

Optional:

- `password` (String) The subject's password for the basic authentication configuration.
- `username` (String) The subject's username for the basic authentication configuration.


<a id="nestedatt--spec--alertmanager--client--header_auth"></a>
### Nested Schema for `spec.alertmanager.client.header_auth`

Optional:

- `credentials` (String) The credentials for the header authentication configuration.
- `credentials_file` (String) The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.
- `type` (String) The authentication type for the header authentication configuration.


<a id="nestedatt--spec--alertmanager--client--tls"></a>
### Nested Schema for `spec.alertmanager.client.tls`

Optional:

- `ca_path` (String) The CA certificate file path for the TLS configuration.
- `cert_path` (String) The client-side certificate file path for the TLS configuration.
- `key_path` (String) The client-side key file path for the TLS configuration.
- `server_name` (String) The server name to validate in the alertmanager server certificates.



<a id="nestedatt--spec--alertmanager--discovery"></a>
### Nested Schema for `spec.alertmanager.discovery`

Optional:

- `enable_srv` (Boolean) Use DNS SRV records to discover Alertmanager hosts.
- `refresh_interval` (String) How long to wait between refreshing DNS resolutions of Alertmanager hosts.


<a id="nestedatt--spec--alertmanager--notification_queue"></a>
### Nested Schema for `spec.alertmanager.notification_queue`

Optional:

- `capacity` (Number) Capacity of the queue for notifications to be sent to the Alertmanager.
- `for_grace_period` (String) Minimum duration between alert and restored 'for' state. This is maintainedonly for alerts with configured 'for' time greater than the grace period.
- `for_outage_tolerance` (String) Max time to tolerate outage for restoring 'for' state of alert.
- `resend_delay` (String) Minimum amount of time to wait before resending an alert to Alertmanager.
- `timeout` (String) HTTP timeout duration when sending notifications to the Alertmanager.


<a id="nestedatt--spec--alertmanager--relabel_configs"></a>
### Nested Schema for `spec.alertmanager.relabel_configs`

Required:

- `source_labels` (List of String) The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.

Optional:

- `action` (String) Action to perform based on regex matching. Default is 'replace'
- `modulus` (Number) Modulus to take of the hash of the source label values.
- `regex` (String) Regular expression against which the extracted value is matched. Default is '(.*)'
- `replacement` (String) Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'
- `separator` (String) Separator placed between concatenated source label values. default is ';'.
- `target_label` (String) Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.



<a id="nestedatt--spec--overrides"></a>
### Nested Schema for `spec.overrides`

Optional:

- `alertmanager` (Attributes) AlertManagerOverrides defines the overrides to apply to the alertmanager config. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager))

<a id="nestedatt--spec--overrides--alertmanager"></a>
### Nested Schema for `spec.overrides.alertmanager`

Required:

- `endpoints` (List of String) List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated asa separate group in the configuration. Multiple Alertmanagers in HA per group can besupported by using DNS resolution (See EnableDNSDiscovery).

Optional:

- `client` (Attributes) Client configuration for reaching the alertmanager endpoint. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--client))
- `discovery` (Attributes) Defines the configuration for DNS-based discovery of AlertManager hosts. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--discovery))
- `enable_v2` (Boolean) If enabled, then requests to Alertmanager use the v2 API.
- `external_labels` (Map of String) Additional labels to add to all alerts.
- `external_url` (String) URL for alerts return path.
- `notification_queue` (Attributes) Defines the configuration for the notification queue to AlertManager hosts. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--notification_queue))
- `relabel_configs` (Attributes List) List of alert relabel configurations. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--relabel_configs))

<a id="nestedatt--spec--overrides--alertmanager--client"></a>
### Nested Schema for `spec.overrides.alertmanager.client`

Optional:

- `basic_auth` (Attributes) Basic authentication configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--relabel_configs--basic_auth))
- `header_auth` (Attributes) Header authentication configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--relabel_configs--header_auth))
- `tls` (Attributes) TLS configuration for reaching the alertmanager endpoints. (see [below for nested schema](#nestedatt--spec--overrides--alertmanager--relabel_configs--tls))

<a id="nestedatt--spec--overrides--alertmanager--relabel_configs--basic_auth"></a>
### Nested Schema for `spec.overrides.alertmanager.relabel_configs.basic_auth`

Optional:

- `password` (String) The subject's password for the basic authentication configuration.
- `username` (String) The subject's username for the basic authentication configuration.


<a id="nestedatt--spec--overrides--alertmanager--relabel_configs--header_auth"></a>
### Nested Schema for `spec.overrides.alertmanager.relabel_configs.header_auth`

Optional:

- `credentials` (String) The credentials for the header authentication configuration.
- `credentials_file` (String) The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.
- `type` (String) The authentication type for the header authentication configuration.


<a id="nestedatt--spec--overrides--alertmanager--relabel_configs--tls"></a>
### Nested Schema for `spec.overrides.alertmanager.relabel_configs.tls`

Optional:

- `ca_path` (String) The CA certificate file path for the TLS configuration.
- `cert_path` (String) The client-side certificate file path for the TLS configuration.
- `key_path` (String) The client-side key file path for the TLS configuration.
- `server_name` (String) The server name to validate in the alertmanager server certificates.



<a id="nestedatt--spec--overrides--alertmanager--discovery"></a>
### Nested Schema for `spec.overrides.alertmanager.discovery`

Optional:

- `enable_srv` (Boolean) Use DNS SRV records to discover Alertmanager hosts.
- `refresh_interval` (String) How long to wait between refreshing DNS resolutions of Alertmanager hosts.


<a id="nestedatt--spec--overrides--alertmanager--notification_queue"></a>
### Nested Schema for `spec.overrides.alertmanager.notification_queue`

Optional:

- `capacity` (Number) Capacity of the queue for notifications to be sent to the Alertmanager.
- `for_grace_period` (String) Minimum duration between alert and restored 'for' state. This is maintainedonly for alerts with configured 'for' time greater than the grace period.
- `for_outage_tolerance` (String) Max time to tolerate outage for restoring 'for' state of alert.
- `resend_delay` (String) Minimum amount of time to wait before resending an alert to Alertmanager.
- `timeout` (String) HTTP timeout duration when sending notifications to the Alertmanager.


<a id="nestedatt--spec--overrides--alertmanager--relabel_configs"></a>
### Nested Schema for `spec.overrides.alertmanager.relabel_configs`

Required:

- `source_labels` (List of String) The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.

Optional:

- `action` (String) Action to perform based on regex matching. Default is 'replace'
- `modulus` (Number) Modulus to take of the hash of the source label values.
- `regex` (String) Regular expression against which the extracted value is matched. Default is '(.*)'
- `replacement` (String) Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'
- `separator` (String) Separator placed between concatenated source label values. default is ';'.
- `target_label` (String) Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.




<a id="nestedatt--spec--remote_write"></a>
### Nested Schema for `spec.remote_write`

Optional:

- `client` (Attributes) Defines the configuration for remote write client. (see [below for nested schema](#nestedatt--spec--remote_write--client))
- `enabled` (Boolean) Enable remote-write functionality.
- `queue` (Attributes) Defines the configuration for remote write client queue. (see [below for nested schema](#nestedatt--spec--remote_write--queue))
- `refresh_period` (String) Minimum period to wait between refreshing remote-write reconfigurations.

<a id="nestedatt--spec--remote_write--client"></a>
### Nested Schema for `spec.remote_write.client`

Required:

- `authorization` (String) Type of authorzation to use to access the remote write endpoint
- `authorization_secret_name` (String) Name of a secret in the namespace configured for authorization secrets.
- `name` (String) Name of the remote write config, which if specified must be unique among remote write configs.
- `url` (String) The URL of the endpoint to send samples to.

Optional:

- `additional_headers` (Map of String) Additional HTTP headers to be sent along with each remote write request.
- `follow_redirects` (Boolean) Configure whether HTTP requests follow HTTP 3xx redirects.
- `proxy_url` (String) Optional proxy URL.
- `relabel_configs` (Attributes List) List of remote write relabel configurations. (see [below for nested schema](#nestedatt--spec--remote_write--client--relabel_configs))
- `timeout` (String) Timeout for requests to the remote write endpoint.

<a id="nestedatt--spec--remote_write--client--relabel_configs"></a>
### Nested Schema for `spec.remote_write.client.relabel_configs`

Required:

- `source_labels` (List of String) The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.

Optional:

- `action` (String) Action to perform based on regex matching. Default is 'replace'
- `modulus` (Number) Modulus to take of the hash of the source label values.
- `regex` (String) Regular expression against which the extracted value is matched. Default is '(.*)'
- `replacement` (String) Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'
- `separator` (String) Separator placed between concatenated source label values. default is ';'.
- `target_label` (String) Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.



<a id="nestedatt--spec--remote_write--queue"></a>
### Nested Schema for `spec.remote_write.queue`

Optional:

- `batch_send_deadline` (String) Maximum time a sample will wait in buffer.
- `capacity` (Number) Number of samples to buffer per shard before we block reading of more
- `max_back_off_period` (String) Maximum retry delay.
- `max_samples_per_send` (Number) Maximum number of samples per send.
- `max_shards` (Number) Maximum number of shards, i.e. amount of concurrency.
- `min_back_off_period` (String) Initial retry delay. Gets doubled for every retry.
- `min_shards` (Number) Minimum number of shards, i.e. amount of concurrency.