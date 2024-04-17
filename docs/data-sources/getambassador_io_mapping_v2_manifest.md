---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_getambassador_io_mapping_v2_manifest Data Source - terraform-provider-k8s"
subcategory: "getambassador.io"
description: |-
  Mapping is the Schema for the mappings API
---

# k8s_getambassador_io_mapping_v2_manifest (Data Source)

Mapping is the Schema for the mappings API

## Example Usage

```terraform
data "k8s_getambassador_io_mapping_v2_manifest" "example" {
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

- `spec` (Attributes) MappingSpec defines the desired state of Mapping (see [below for nested schema](#nestedatt--spec))

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

- `prefix` (String)
- `service` (String)

Optional:

- `add_linkerd_headers` (Boolean)
- `add_request_headers` (Map of String)
- `add_response_headers` (Map of String)
- `allow_upgrade` (List of String) A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write  allow_upgrade: - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write  allow_upgrade: - spdy/3.1
- `ambassador_id` (List of String) AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  ambassador_id: - 'default'
- `auth_context_extensions` (Map of String)
- `auto_host_rewrite` (Boolean)
- `bypass_auth` (Boolean)
- `bypass_error_response_overrides` (Boolean) If true, bypasses any 'error_response_overrides' set on the Ambassador module.
- `case_sensitive` (Boolean)
- `circuit_breakers` (Attributes List) (see [below for nested schema](#nestedatt--spec--circuit_breakers))
- `cluster_idle_timeout_ms` (Number)
- `cluster_max_connection_lifetime_ms` (Number)
- `cluster_tag` (String)
- `connect_timeout_ms` (Number)
- `cors` (Attributes) (see [below for nested schema](#nestedatt--spec--cors))
- `dns_type` (String)
- `docs` (Attributes) DocsInfo provides some extra information about the docs for the Mapping (used by the Dev Portal) (see [below for nested schema](#nestedatt--spec--docs))
- `enable_ipv4` (Boolean)
- `enable_ipv6` (Boolean)
- `envoy_override` (Map of String)
- `error_response_overrides` (Attributes List) Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any. (see [below for nested schema](#nestedatt--spec--error_response_overrides))
- `grpc` (Boolean)
- `headers` (Map of String)
- `host` (String)
- `host_redirect` (Boolean)
- `host_regex` (Boolean)
- `host_rewrite` (String)
- `idle_timeout_ms` (Number)
- `keepalive` (Attributes) (see [below for nested schema](#nestedatt--spec--keepalive))
- `labels` (Map of String) A DomainMap is the overall Mapping.spec.Labels type. It maps domains (kind of like namespaces for Mapping labels) to arrays of label groups.
- `load_balancer` (Attributes) (see [below for nested schema](#nestedatt--spec--load_balancer))
- `method` (String)
- `method_regex` (Boolean)
- `modules` (List of Map of String)
- `outlier_detection` (String)
- `path_redirect` (String) Path replacement to use when generating an HTTP redirect. Used with 'host_redirect'.
- `precedence` (Number)
- `prefix_exact` (Boolean)
- `prefix_redirect` (String) Prefix rewrite to use when generating an HTTP redirect. Used with 'host_redirect'.
- `prefix_regex` (Boolean)
- `priority` (String)
- `query_parameters` (Map of String)
- `redirect_response_code` (Number) The response code to use when generating an HTTP redirect. Defaults to 301. Used with 'host_redirect'.
- `regex_headers` (Map of String)
- `regex_query_parameters` (Map of String)
- `regex_redirect` (Attributes) Prefix regex rewrite to use when generating an HTTP redirect. Used with 'host_redirect'. (see [below for nested schema](#nestedatt--spec--regex_redirect))
- `regex_rewrite` (Attributes) (see [below for nested schema](#nestedatt--spec--regex_rewrite))
- `remove_request_headers` (List of String) StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.
- `remove_response_headers` (List of String) StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.
- `resolver` (String)
- `respect_dns_ttl` (Boolean)
- `retry_policy` (Attributes) (see [below for nested schema](#nestedatt--spec--retry_policy))
- `rewrite` (String)
- `shadow` (Boolean)
- `timeout_ms` (Number) The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.
- `tls` (Boolean)
- `use_websocket` (Boolean) use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'
- `v3_stats_name` (String)
- `v3health_checks` (Attributes List) (see [below for nested schema](#nestedatt--spec--v3health_checks))
- `weight` (Number)

<a id="nestedatt--spec--circuit_breakers"></a>
### Nested Schema for `spec.circuit_breakers`

Optional:

- `max_connections` (Number)
- `max_pending_requests` (Number)
- `max_requests` (Number)
- `max_retries` (Number)
- `priority` (String)


<a id="nestedatt--spec--cors"></a>
### Nested Schema for `spec.cors`

Optional:

- `credentials` (Boolean)
- `exposed_headers` (List of String) StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.
- `headers` (List of String) StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.
- `max_age` (String)
- `methods` (List of String) StringOrStringList is just what it says on the tin, but note that it will always marshal as a list of strings right now.
- `origins` (List of String)


<a id="nestedatt--spec--docs"></a>
### Nested Schema for `spec.docs`

Optional:

- `display_name` (String)
- `ignored` (Boolean)
- `path` (String)
- `timeout_ms` (Number)
- `url` (String)


<a id="nestedatt--spec--error_response_overrides"></a>
### Nested Schema for `spec.error_response_overrides`

Required:

- `body` (Attributes) The new response body (see [below for nested schema](#nestedatt--spec--error_response_overrides--body))
- `on_status_code` (Number) The status code to match on -- not a pointer because it's required.

<a id="nestedatt--spec--error_response_overrides--body"></a>
### Nested Schema for `spec.error_response_overrides.body`

Optional:

- `content_type` (String) The content type to set on the error response body when using text_format or text_format_source. Defaults to 'text/plain'.
- `json_format` (Map of String) A JSON response with content-type: application/json. The values can contain format text like in text_format.
- `text_format` (String) A format string representing a text response body. Content-Type can be set using the 'content_type' field below.
- `text_format_source` (Attributes) A format string sourced from a file on the Ambassador container. Useful for larger response bodies that should not be placed inline in configuration. (see [below for nested schema](#nestedatt--spec--error_response_overrides--body--text_format_source))

<a id="nestedatt--spec--error_response_overrides--body--text_format_source"></a>
### Nested Schema for `spec.error_response_overrides.body.text_format_source`

Optional:

- `filename` (String) The name of a file on the Ambassador pod that contains a format text string.




<a id="nestedatt--spec--keepalive"></a>
### Nested Schema for `spec.keepalive`

Optional:

- `idle_time` (Number)
- `interval` (Number)
- `probes` (Number)


<a id="nestedatt--spec--load_balancer"></a>
### Nested Schema for `spec.load_balancer`

Required:

- `policy` (String)

Optional:

- `cookie` (Attributes) (see [below for nested schema](#nestedatt--spec--load_balancer--cookie))
- `header` (String)
- `source_ip` (Boolean)

<a id="nestedatt--spec--load_balancer--cookie"></a>
### Nested Schema for `spec.load_balancer.cookie`

Required:

- `name` (String)

Optional:

- `path` (String)
- `ttl` (String)



<a id="nestedatt--spec--regex_redirect"></a>
### Nested Schema for `spec.regex_redirect`

Optional:

- `pattern` (String)
- `substitution` (String)


<a id="nestedatt--spec--regex_rewrite"></a>
### Nested Schema for `spec.regex_rewrite`

Optional:

- `pattern` (String)
- `substitution` (String)


<a id="nestedatt--spec--retry_policy"></a>
### Nested Schema for `spec.retry_policy`

Optional:

- `num_retries` (Number)
- `per_try_timeout` (String)
- `retry_on` (String)


<a id="nestedatt--spec--v3health_checks"></a>
### Nested Schema for `spec.v3health_checks`

Required:

- `health_check` (Attributes) Configuration for where the healthcheck request should be made to (see [below for nested schema](#nestedatt--spec--v3health_checks--health_check))

Optional:

- `healthy_threshold` (Number) Number of expected responses for the upstream to be considered healthy. Defaults to 1.
- `interval` (String) Interval between health checks. Defaults to every 5 seconds.
- `timeout` (String) Timeout for connecting to the health checking endpoint. Defaults to 3 seconds.
- `unhealthy_threshold` (Number) Number of non-expected responses for the upstream to be considered unhealthy. A single 503 will mark the upstream as unhealthy regardless of the threshold. Defaults to 2.

<a id="nestedatt--spec--v3health_checks--health_check"></a>
### Nested Schema for `spec.v3health_checks.health_check`

Optional:

- `grpc` (Attributes) HealthCheck for gRPC upstreams. Only one of grpc_health_check or http_health_check may be specified (see [below for nested schema](#nestedatt--spec--v3health_checks--health_check--grpc))
- `http` (Attributes) HealthCheck for HTTP upstreams. Only one of http_health_check or grpc_health_check may be specified (see [below for nested schema](#nestedatt--spec--v3health_checks--health_check--http))

<a id="nestedatt--spec--v3health_checks--health_check--grpc"></a>
### Nested Schema for `spec.v3health_checks.health_check.grpc`

Required:

- `upstream_name` (String) The upstream name parameter which will be sent to gRPC service in the health check message

Optional:

- `authority` (String) The value of the :authority header in the gRPC health check request. If left empty the upstream name will be used.


<a id="nestedatt--spec--v3health_checks--health_check--http"></a>
### Nested Schema for `spec.v3health_checks.health_check.http`

Required:

- `path` (String)

Optional:

- `add_request_headers` (Attributes) (see [below for nested schema](#nestedatt--spec--v3health_checks--health_check--http--add_request_headers))
- `expected_statuses` (Attributes List) (see [below for nested schema](#nestedatt--spec--v3health_checks--health_check--http--expected_statuses))
- `hostname` (String)
- `remove_request_headers` (List of String)

<a id="nestedatt--spec--v3health_checks--health_check--http--add_request_headers"></a>
### Nested Schema for `spec.v3health_checks.health_check.http.add_request_headers`

Optional:

- `append` (Boolean)
- `v2_representation` (String)
- `value` (String)


<a id="nestedatt--spec--v3health_checks--health_check--http--expected_statuses"></a>
### Nested Schema for `spec.v3health_checks.health_check.http.expected_statuses`

Required:

- `max` (Number) End of the statuses to include. Must be between 100 and 599 (inclusive)
- `min` (Number) Start of the statuses to include. Must be between 100 and 599 (inclusive)