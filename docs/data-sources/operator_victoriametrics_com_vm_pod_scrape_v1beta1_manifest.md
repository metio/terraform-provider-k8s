---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "operator.victoriametrics.com"
description: |-
  VMPodScrape is scrape configuration for pods,it generates vmagent's config for scraping pod targetsbased on selectors.
---

# k8s_operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest (Data Source)

VMPodScrape is scrape configuration for pods,it generates vmagent's config for scraping pod targetsbased on selectors.

## Example Usage

```terraform
data "k8s_operator_victoriametrics_com_vm_pod_scrape_v1beta1_manifest" "example" {
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

- `spec` (Attributes) VMPodScrapeSpec defines the desired state of VMPodScrape (see [below for nested schema](#nestedatt--spec))

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

Required:

- `pod_metrics_endpoints` (Attributes List) A list of endpoints allowed as part of this PodMonitor. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints))

Optional:

- `attach_metadata` (Attributes) AttachMetadata configures metadata attaching from service discovery (see [below for nested schema](#nestedatt--spec--attach_metadata))
- `job_label` (String) The label to use to retrieve the job name from.
- `namespace_selector` (Attributes) Selector to select which namespaces the Endpoints objects are discovered from. (see [below for nested schema](#nestedatt--spec--namespace_selector))
- `pod_target_labels` (List of String) PodTargetLabels transfers labels on the Kubernetes Pod onto the target.
- `sample_limit` (Number) SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
- `selector` (Attributes) Selector to select Pod objects. (see [below for nested schema](#nestedatt--spec--selector))

<a id="nestedatt--spec--pod_metrics_endpoints"></a>
### Nested Schema for `spec.pod_metrics_endpoints`

Optional:

- `attach_metadata` (Attributes) AttachMetadata configures metadata attaching from service discovery (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--attach_metadata))
- `authorization` (Attributes) Authorization with http header Authorization (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--authorization))
- `basic_auth` (Attributes) BasicAuth allow an endpoint to authenticate over basic authenticationMore info: https://prometheus.io/docs/operating/configuration/#endpoints (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--basic_auth))
- `bearer_token_file` (String) File to read bearer token for scraping targets.
- `bearer_token_secret` (Attributes) Secret to mount to read bearer token for scraping targets. The secretneeds to be in the same namespace as the service scrape and accessible bythe victoria-metrics operator. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--bearer_token_secret))
- `filter_running` (Boolean) FilterRunning applies filter with pod status == runningit prevents from scrapping metrics at failed or succeed state pods.enabled by default
- `follow_redirects` (Boolean) FollowRedirects controls redirects for scraping.
- `honor_labels` (Boolean) HonorLabels chooses the metric's labels on collisions with target labels.
- `honor_timestamps` (Boolean) HonorTimestamps controls whether vmagent respects the timestamps present in scraped data.
- `interval` (String) Interval at which metrics should be scraped
- `metric_relabel_configs` (Attributes List) MetricRelabelConfigs to apply to samples before ingestion. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--metric_relabel_configs))
- `oauth2` (Attributes) OAuth2 defines auth configuration (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--oauth2))
- `params` (Map of List of String) Optional HTTP URL parameters
- `path` (String) HTTP path to scrape for metrics.
- `port` (String) Name of the pod port this endpoint refers to. Mutually exclusive with targetPort.
- `proxy_url` (String) ProxyURL eg http://proxyserver:2195 Directs scrapes to proxy through this endpoint.
- `relabel_configs` (Attributes List) RelabelConfigs to apply to samples before ingestion.More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--relabel_configs))
- `sample_limit` (Number) SampleLimit defines per-podEndpoint limit on number of scraped samples that will be accepted.
- `scheme` (String) HTTP scheme to use for scraping.
- `scrape_interval` (String) ScrapeInterval is the same as Interval and has priority over it.one of scrape_interval or interval can be used
- `scrape_timeout` (String) Timeout after which the scrape is ended
- `target_port` (String) Deprecated: Use 'port' instead.
- `tls_config` (Attributes) TLSConfig configuration to use when scraping the endpoint (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config))
- `vm_scrape_params` (Attributes) VMScrapeParams defines VictoriaMetrics specific scrape parametrs (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params))

<a id="nestedatt--spec--pod_metrics_endpoints--attach_metadata"></a>
### Nested Schema for `spec.pod_metrics_endpoints.attach_metadata`

Optional:

- `node` (Boolean) Node instructs vmagent to add node specific metadata from service discoveryValid for roles: pod, endpoints, endpointslice.


<a id="nestedatt--spec--pod_metrics_endpoints--authorization"></a>
### Nested Schema for `spec.pod_metrics_endpoints.authorization`

Optional:

- `credentials` (Attributes) Reference to the secret with value for authorization (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--authorization--credentials))
- `credentials_file` (String) File with value for authorization
- `type` (String) Type of authorization, default to bearer

<a id="nestedatt--spec--pod_metrics_endpoints--authorization--credentials"></a>
### Nested Schema for `spec.pod_metrics_endpoints.authorization.credentials`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--basic_auth"></a>
### Nested Schema for `spec.pod_metrics_endpoints.basic_auth`

Optional:

- `password` (Attributes) The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--basic_auth--password))
- `password_file` (String) PasswordFile defines path to password file at disk
- `username` (Attributes) The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--basic_auth--username))

<a id="nestedatt--spec--pod_metrics_endpoints--basic_auth--password"></a>
### Nested Schema for `spec.pod_metrics_endpoints.basic_auth.password`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--basic_auth--username"></a>
### Nested Schema for `spec.pod_metrics_endpoints.basic_auth.username`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--bearer_token_secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.bearer_token_secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--metric_relabel_configs"></a>
### Nested Schema for `spec.pod_metrics_endpoints.metric_relabel_configs`

Optional:

- `action` (String) Action to perform based on regex matching. Default is 'replace'
- `if` (Map of String) If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'
- `labels` (Map of String) Labels is used together with Match for 'action: graphite'
- `match` (String) Match is used together with Labels for 'action: graphite'
- `modulus` (Number) Modulus to take of the hash of the source label values.
- `regex` (String) Regular expression against which the extracted value is matched. Default is '(.*)'
- `replacement` (String) Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'
- `separator` (String) Separator placed between concatenated source label values. default is ';'.
- `source_labels` (List of String) The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.
- `target_label` (String) Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.


<a id="nestedatt--spec--pod_metrics_endpoints--oauth2"></a>
### Nested Schema for `spec.pod_metrics_endpoints.oauth2`

Required:

- `client_id` (Attributes) The secret or configmap containing the OAuth2 client id (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--oauth2--client_id))
- `token_url` (String) The URL to fetch the token from

Optional:

- `client_secret` (Attributes) The secret containing the OAuth2 client secret (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--oauth2--client_secret))
- `client_secret_file` (String) ClientSecretFile defines path for client secret file.
- `endpoint_params` (Map of String) Parameters to append to the token URL
- `scopes` (List of String) OAuth2 scopes used for the token request

<a id="nestedatt--spec--pod_metrics_endpoints--oauth2--client_id"></a>
### Nested Schema for `spec.pod_metrics_endpoints.oauth2.client_id`

Optional:

- `config_map` (Attributes) ConfigMap containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--oauth2--scopes--config_map))
- `secret` (Attributes) Secret containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--oauth2--scopes--secret))

<a id="nestedatt--spec--pod_metrics_endpoints--oauth2--scopes--config_map"></a>
### Nested Schema for `spec.pod_metrics_endpoints.oauth2.scopes.config_map`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--oauth2--scopes--secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.oauth2.scopes.secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--oauth2--client_secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.oauth2.client_secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--relabel_configs"></a>
### Nested Schema for `spec.pod_metrics_endpoints.relabel_configs`

Optional:

- `action` (String) Action to perform based on regex matching. Default is 'replace'
- `if` (Map of String) If represents metricsQL match expression (or list of expressions): '{__name__=~'foo_.*'}'
- `labels` (Map of String) Labels is used together with Match for 'action: graphite'
- `match` (String) Match is used together with Labels for 'action: graphite'
- `modulus` (Number) Modulus to take of the hash of the source label values.
- `regex` (String) Regular expression against which the extracted value is matched. Default is '(.*)'
- `replacement` (String) Replacement value against which a regex replace is performed if theregular expression matches. Regex capture groups are available. Default is '$1'
- `separator` (String) Separator placed between concatenated source label values. default is ';'.
- `source_labels` (List of String) The source labels select values from existing labels. Their content is concatenatedusing the configured separator and matched against the configured regular expressionfor the replace, keep, and drop actions.
- `target_label` (String) Label to which the resulting value is written in a replace action.It is mandatory for replace actions. Regex capture groups are available.


<a id="nestedatt--spec--pod_metrics_endpoints--tls_config"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config`

Optional:

- `ca` (Attributes) Stuct containing the CA cert to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--ca))
- `ca_file` (String) Path to the CA cert in the container to use for the targets.
- `cert` (Attributes) Struct containing the client cert file for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--cert))
- `cert_file` (String) Path to the client cert file in the container for the targets.
- `insecure_skip_verify` (Boolean) Disable target certificate validation.
- `key_file` (String) Path to the client key file in the container for the targets.
- `key_secret` (Attributes) Secret containing the client key file for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--key_secret))
- `server_name` (String) Used to verify the hostname for the targets.

<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--ca"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.ca`

Optional:

- `config_map` (Attributes) ConfigMap containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--config_map))
- `secret` (Attributes) Secret containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--secret))

<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--config_map"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.server_name.config_map`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.server_name.secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--cert"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.cert`

Optional:

- `config_map` (Attributes) ConfigMap containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--config_map))
- `secret` (Attributes) Secret containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--secret))

<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--config_map"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.server_name.config_map`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--server_name--secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.server_name.secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--tls_config--key_secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.tls_config.key_secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params`

Optional:

- `disable_compression` (Boolean)
- `disable_keep_alive` (Boolean)
- `headers` (List of String) Headers allows sending custom headers to scrape targetsmust be in of semicolon separated header with it's valueeg:headerName: headerValuevmagent supports since 1.79.0 version
- `metric_relabel_debug` (Boolean)
- `no_stale_markers` (Boolean)
- `proxy_client_config` (Attributes) ProxyClientConfig configures proxy auth settings for scrapingSee feature description https://docs.victoriametrics.com/vmagent.html#scraping-targets-via-a-proxy (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--proxy_client_config))
- `relabel_debug` (Boolean)
- `scrape_align_interval` (String)
- `scrape_offset` (String)
- `stream_parse` (Boolean)

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--proxy_client_config"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.proxy_client_config`

Optional:

- `basic_auth` (Attributes) BasicAuth allow an endpoint to authenticate over basic authentication (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth))
- `bearer_token` (Attributes) SecretKeySelector selects a key of a Secret. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--bearer_token))
- `bearer_token_file` (String)
- `tls_config` (Attributes) TLSConfig specifies TLSConfig configuration parameters. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config))

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.basic_auth`

Optional:

- `password` (Attributes) The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth--password))
- `password_file` (String) PasswordFile defines path to password file at disk
- `username` (Attributes) The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth--username))

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth--password"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.basic_auth.password`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--basic_auth--username"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.basic_auth.username`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--bearer_token"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.bearer_token`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config`

Optional:

- `ca` (Attributes) Stuct containing the CA cert to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--ca))
- `ca_file` (String) Path to the CA cert in the container to use for the targets.
- `cert` (Attributes) Struct containing the client cert file for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--cert))
- `cert_file` (String) Path to the client cert file in the container for the targets.
- `insecure_skip_verify` (Boolean) Disable target certificate validation.
- `key_file` (String) Path to the client key file in the container for the targets.
- `key_secret` (Attributes) Secret containing the client key file for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--key_secret))
- `server_name` (String) Used to verify the hostname for the targets.

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--ca"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.ca`

Optional:

- `config_map` (Attributes) ConfigMap containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--config_map))
- `secret` (Attributes) Secret containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--secret))

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--config_map"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.server_name.config_map`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.server_name.secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--cert"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.cert`

Optional:

- `config_map` (Attributes) ConfigMap containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--config_map))
- `secret` (Attributes) Secret containing data to use for the targets. (see [below for nested schema](#nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--secret))

<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--config_map"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.server_name.config_map`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--server_name--secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.server_name.secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--pod_metrics_endpoints--vm_scrape_params--stream_parse--tls_config--key_secret"></a>
### Nested Schema for `spec.pod_metrics_endpoints.vm_scrape_params.stream_parse.tls_config.key_secret`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined






<a id="nestedatt--spec--attach_metadata"></a>
### Nested Schema for `spec.attach_metadata`

Optional:

- `node` (Boolean) Node instructs vmagent to add node specific metadata from service discoveryValid for roles: pod, endpoints, endpointslice.


<a id="nestedatt--spec--namespace_selector"></a>
### Nested Schema for `spec.namespace_selector`

Optional:

- `any` (Boolean) Boolean describing whether all namespaces are selected in contrast to alist restricting them.
- `match_names` (List of String) List of namespace names.


<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.