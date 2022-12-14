---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networking_istio_io_destination_rule_v1beta1 Resource - terraform-provider-k8s"
subcategory: "networking.istio.io"
description: |-
  
---

# k8s_networking_istio_io_destination_rule_v1beta1 (Resource)



## Example Usage

```terraform
resource "k8s_networking_istio_io_destination_rule_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `export_to` (List of String) A list of namespaces to which this destination rule is exported.
- `host` (String) The name of a service from the service registry.
- `subsets` (Attributes List) (see [below for nested schema](#nestedatt--spec--subsets))
- `traffic_policy` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy))
- `workload_selector` (Attributes) (see [below for nested schema](#nestedatt--spec--workload_selector))

<a id="nestedatt--spec--subsets"></a>
### Nested Schema for `spec.subsets`

Optional:

- `labels` (Map of String)
- `name` (String) Name of the subset.
- `traffic_policy` (Attributes) Traffic policies that apply to this subset. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy))

<a id="nestedatt--spec--subsets--traffic_policy"></a>
### Nested Schema for `spec.subsets.traffic_policy`

Optional:

- `connection_pool` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--connection_pool))
- `load_balancer` (Attributes) Settings controlling the load balancer algorithms. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--load_balancer))
- `outlier_detection` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--outlier_detection))
- `port_level_settings` (Attributes List) Traffic policies specific to individual ports. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--port_level_settings))
- `tls` (Attributes) TLS related settings for connections to the upstream service. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tls))
- `tunnel` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel))

<a id="nestedatt--spec--subsets--traffic_policy--connection_pool"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `http` (Attributes) HTTP connection pool settings. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--http))
- `tcp` (Attributes) Settings common to both HTTP and TCP upstream connections. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--tcp))

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--http"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.http`

Optional:

- `h2_upgrade_policy` (String) Specify if http1.1 connection should be upgraded to http2 for the associated destination.
- `http1_max_pending_requests` (Number)
- `http2_max_requests` (Number) Maximum number of active requests to a destination.
- `idle_timeout` (String) The idle timeout for upstream connection pool connections.
- `max_requests_per_connection` (Number) Maximum number of requests per connection to a backend.
- `max_retries` (Number)
- `use_client_protocol` (Boolean) If set to true, client protocol will be preserved while initiating connection to backend.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--tcp"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.tcp`

Optional:

- `connect_timeout` (String) TCP connection timeout.
- `max_connection_duration` (String) The maximum duration of a connection.
- `max_connections` (Number) Maximum number of HTTP1 /TCP connections to a destination host.
- `tcp_keepalive` (Attributes) If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--tcp--tcp_keepalive))

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--tcp--tcp_keepalive"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.tcp.tcp_keepalive`

Optional:

- `interval` (String) The time duration between keep-alive probes.
- `probes` (Number)
- `time` (String)




<a id="nestedatt--spec--subsets--traffic_policy--load_balancer"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `consistent_hash` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash))
- `locality_lb_setting` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting))
- `simple` (String)
- `warmup_duration_secs` (String) Represents the warmup duration of Service.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.consistent_hash`

Optional:

- `http_cookie` (Attributes) Hash based on HTTP cookie. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--http_cookie))
- `http_header_name` (String) Hash based on a specific HTTP header.
- `http_query_parameter_name` (String) Hash based on a specific HTTP query parameter.
- `maglev` (Attributes) The Maglev load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--maglev))
- `minimum_ring_size` (Number) Deprecated.
- `ring_hash` (Attributes) The ring/modulo hash load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--ring_hash))
- `use_source_ip` (Boolean) Hash based on the source IP address.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--http_cookie"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.consistent_hash.use_source_ip`

Optional:

- `name` (String) Name of the cookie.
- `path` (String) Path to set for the cookie.
- `ttl` (String) Lifetime of the cookie.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--maglev"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.consistent_hash.use_source_ip`

Optional:

- `table_size` (Number) The table size for Maglev hashing.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--consistent_hash--ring_hash"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.consistent_hash.use_source_ip`

Optional:

- `minimum_ring_size` (Number)



<a id="nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.locality_lb_setting`

Optional:

- `distribute` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting--distribute))
- `enabled` (Boolean) enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
- `failover` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting--failover))
- `failover_priority` (List of String) failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting--distribute"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.locality_lb_setting.failover_priority`

Optional:

- `from` (String) Originating locality, '/' separated, e.g.
- `to` (Map of String) Map of upstream localities to traffic distribution weights.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--locality_lb_setting--failover"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.locality_lb_setting.failover_priority`

Optional:

- `from` (String) Originating region.
- `to` (String)




<a id="nestedatt--spec--subsets--traffic_policy--outlier_detection"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `base_ejection_time` (String) Minimum ejection duration.
- `consecutive5xx_errors` (Number) Number of 5xx errors before a host is ejected from the connection pool.
- `consecutive_errors` (Number)
- `consecutive_gateway_errors` (Number) Number of gateway errors before a host is ejected from the connection pool.
- `consecutive_local_origin_failures` (Number)
- `interval` (String) Time interval between ejection sweep analysis.
- `max_ejection_percent` (Number)
- `min_health_percent` (Number)
- `split_external_local_origin_errors` (Boolean) Determines whether to distinguish local origin failures from external errors.


<a id="nestedatt--spec--subsets--traffic_policy--port_level_settings"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `connection_pool` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool))
- `load_balancer` (Attributes) Settings controlling the load balancer algorithms. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer))
- `outlier_detection` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--outlier_detection))
- `port` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--port))
- `tls` (Attributes) TLS related settings for connections to the upstream service. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--tls))

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.connection_pool`

Optional:

- `http` (Attributes) HTTP connection pool settings. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--http))
- `tcp` (Attributes) Settings common to both HTTP and TCP upstream connections. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--tcp))

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--http"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.connection_pool.tcp`

Optional:

- `h2_upgrade_policy` (String) Specify if http1.1 connection should be upgraded to http2 for the associated destination.
- `http1_max_pending_requests` (Number)
- `http2_max_requests` (Number) Maximum number of active requests to a destination.
- `idle_timeout` (String) The idle timeout for upstream connection pool connections.
- `max_requests_per_connection` (Number) Maximum number of requests per connection to a backend.
- `max_retries` (Number)
- `use_client_protocol` (Boolean) If set to true, client protocol will be preserved while initiating connection to backend.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--tcp"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.connection_pool.tcp`

Optional:

- `connect_timeout` (String) TCP connection timeout.
- `max_connection_duration` (String) The maximum duration of a connection.
- `max_connections` (Number) Maximum number of HTTP1 /TCP connections to a destination host.
- `tcp_keepalive` (Attributes) If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--tcp--tcp_keepalive))

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--connection_pool--tcp--tcp_keepalive"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.connection_pool.tcp.tcp_keepalive`

Optional:

- `interval` (String) The time duration between keep-alive probes.
- `probes` (Number)
- `time` (String)




<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer`

Optional:

- `consistent_hash` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--consistent_hash))
- `locality_lb_setting` (Attributes) (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--locality_lb_setting))
- `simple` (String)
- `warmup_duration_secs` (String) Represents the warmup duration of Service.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--consistent_hash"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs`

Optional:

- `http_cookie` (Attributes) Hash based on HTTP cookie. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--http_cookie))
- `http_header_name` (String) Hash based on a specific HTTP header.
- `http_query_parameter_name` (String) Hash based on a specific HTTP query parameter.
- `maglev` (Attributes) The Maglev load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--maglev))
- `minimum_ring_size` (Number) Deprecated.
- `ring_hash` (Attributes) The ring/modulo hash load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--ring_hash))
- `use_source_ip` (Boolean) Hash based on the source IP address.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--http_cookie"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs.use_source_ip`

Optional:

- `name` (String) Name of the cookie.
- `path` (String) Path to set for the cookie.
- `ttl` (String) Lifetime of the cookie.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--maglev"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs.use_source_ip`

Optional:

- `table_size` (Number) The table size for Maglev hashing.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--ring_hash"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs.use_source_ip`

Optional:

- `minimum_ring_size` (Number)



<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--locality_lb_setting"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs`

Optional:

- `distribute` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--distribute))
- `enabled` (Boolean) enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
- `failover` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--failover))
- `failover_priority` (List of String) failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.

<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--distribute"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs.failover_priority`

Optional:

- `from` (String) Originating locality, '/' separated, e.g.
- `to` (Map of String) Map of upstream localities to traffic distribution weights.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--load_balancer--warmup_duration_secs--failover"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.load_balancer.warmup_duration_secs.failover_priority`

Optional:

- `from` (String) Originating region.
- `to` (String)




<a id="nestedatt--spec--subsets--traffic_policy--tunnel--outlier_detection"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.outlier_detection`

Optional:

- `base_ejection_time` (String) Minimum ejection duration.
- `consecutive5xx_errors` (Number) Number of 5xx errors before a host is ejected from the connection pool.
- `consecutive_errors` (Number)
- `consecutive_gateway_errors` (Number) Number of gateway errors before a host is ejected from the connection pool.
- `consecutive_local_origin_failures` (Number)
- `interval` (String) Time interval between ejection sweep analysis.
- `max_ejection_percent` (Number)
- `min_health_percent` (Number)
- `split_external_local_origin_errors` (Boolean) Determines whether to distinguish local origin failures from external errors.


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--port"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.port`

Optional:

- `number` (Number)


<a id="nestedatt--spec--subsets--traffic_policy--tunnel--tls"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel.tls`

Optional:

- `ca_certificates` (String)
- `client_certificate` (String) REQUIRED if mode is 'MUTUAL'.
- `credential_name` (String)
- `insecure_skip_verify` (Boolean)
- `mode` (String)
- `private_key` (String) REQUIRED if mode is 'MUTUAL'.
- `sni` (String) SNI string to present to the server during TLS handshake.
- `subject_alt_names` (List of String)



<a id="nestedatt--spec--subsets--traffic_policy--tls"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `ca_certificates` (String)
- `client_certificate` (String) REQUIRED if mode is 'MUTUAL'.
- `credential_name` (String)
- `insecure_skip_verify` (Boolean)
- `mode` (String)
- `private_key` (String) REQUIRED if mode is 'MUTUAL'.
- `sni` (String) SNI string to present to the server during TLS handshake.
- `subject_alt_names` (List of String)


<a id="nestedatt--spec--subsets--traffic_policy--tunnel"></a>
### Nested Schema for `spec.subsets.traffic_policy.tunnel`

Optional:

- `protocol` (String) Specifies which protocol to use for tunneling the downstream connection.
- `target_host` (String) Specifies a host to which the downstream connection is tunneled.
- `target_port` (Number) Specifies a port to which the downstream connection is tunneled.




<a id="nestedatt--spec--traffic_policy"></a>
### Nested Schema for `spec.traffic_policy`

Optional:

- `connection_pool` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--connection_pool))
- `load_balancer` (Attributes) Settings controlling the load balancer algorithms. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer))
- `outlier_detection` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--outlier_detection))
- `port_level_settings` (Attributes List) Traffic policies specific to individual ports. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings))
- `tls` (Attributes) TLS related settings for connections to the upstream service. (see [below for nested schema](#nestedatt--spec--traffic_policy--tls))
- `tunnel` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--tunnel))

<a id="nestedatt--spec--traffic_policy--connection_pool"></a>
### Nested Schema for `spec.traffic_policy.connection_pool`

Optional:

- `http` (Attributes) HTTP connection pool settings. (see [below for nested schema](#nestedatt--spec--traffic_policy--connection_pool--http))
- `tcp` (Attributes) Settings common to both HTTP and TCP upstream connections. (see [below for nested schema](#nestedatt--spec--traffic_policy--connection_pool--tcp))

<a id="nestedatt--spec--traffic_policy--connection_pool--http"></a>
### Nested Schema for `spec.traffic_policy.connection_pool.tcp`

Optional:

- `h2_upgrade_policy` (String) Specify if http1.1 connection should be upgraded to http2 for the associated destination.
- `http1_max_pending_requests` (Number)
- `http2_max_requests` (Number) Maximum number of active requests to a destination.
- `idle_timeout` (String) The idle timeout for upstream connection pool connections.
- `max_requests_per_connection` (Number) Maximum number of requests per connection to a backend.
- `max_retries` (Number)
- `use_client_protocol` (Boolean) If set to true, client protocol will be preserved while initiating connection to backend.


<a id="nestedatt--spec--traffic_policy--connection_pool--tcp"></a>
### Nested Schema for `spec.traffic_policy.connection_pool.tcp`

Optional:

- `connect_timeout` (String) TCP connection timeout.
- `max_connection_duration` (String) The maximum duration of a connection.
- `max_connections` (Number) Maximum number of HTTP1 /TCP connections to a destination host.
- `tcp_keepalive` (Attributes) If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives. (see [below for nested schema](#nestedatt--spec--traffic_policy--connection_pool--tcp--tcp_keepalive))

<a id="nestedatt--spec--traffic_policy--connection_pool--tcp--tcp_keepalive"></a>
### Nested Schema for `spec.traffic_policy.connection_pool.tcp.tcp_keepalive`

Optional:

- `interval` (String) The time duration between keep-alive probes.
- `probes` (Number)
- `time` (String)




<a id="nestedatt--spec--traffic_policy--load_balancer"></a>
### Nested Schema for `spec.traffic_policy.load_balancer`

Optional:

- `consistent_hash` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--consistent_hash))
- `locality_lb_setting` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--locality_lb_setting))
- `simple` (String)
- `warmup_duration_secs` (String) Represents the warmup duration of Service.

<a id="nestedatt--spec--traffic_policy--load_balancer--consistent_hash"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs`

Optional:

- `http_cookie` (Attributes) Hash based on HTTP cookie. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--http_cookie))
- `http_header_name` (String) Hash based on a specific HTTP header.
- `http_query_parameter_name` (String) Hash based on a specific HTTP query parameter.
- `maglev` (Attributes) The Maglev load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--maglev))
- `minimum_ring_size` (Number) Deprecated.
- `ring_hash` (Attributes) The ring/modulo hash load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--ring_hash))
- `use_source_ip` (Boolean) Hash based on the source IP address.

<a id="nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--http_cookie"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs.http_cookie`

Optional:

- `name` (String) Name of the cookie.
- `path` (String) Path to set for the cookie.
- `ttl` (String) Lifetime of the cookie.


<a id="nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--maglev"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs.maglev`

Optional:

- `table_size` (Number) The table size for Maglev hashing.


<a id="nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--ring_hash"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs.ring_hash`

Optional:

- `minimum_ring_size` (Number)



<a id="nestedatt--spec--traffic_policy--load_balancer--locality_lb_setting"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs`

Optional:

- `distribute` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--distribute))
- `enabled` (Boolean) enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
- `failover` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--failover))
- `failover_priority` (List of String) failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.

<a id="nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--distribute"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs.distribute`

Optional:

- `from` (String) Originating locality, '/' separated, e.g.
- `to` (Map of String) Map of upstream localities to traffic distribution weights.


<a id="nestedatt--spec--traffic_policy--load_balancer--warmup_duration_secs--failover"></a>
### Nested Schema for `spec.traffic_policy.load_balancer.warmup_duration_secs.failover`

Optional:

- `from` (String) Originating region.
- `to` (String)




<a id="nestedatt--spec--traffic_policy--outlier_detection"></a>
### Nested Schema for `spec.traffic_policy.outlier_detection`

Optional:

- `base_ejection_time` (String) Minimum ejection duration.
- `consecutive5xx_errors` (Number) Number of 5xx errors before a host is ejected from the connection pool.
- `consecutive_errors` (Number)
- `consecutive_gateway_errors` (Number) Number of gateway errors before a host is ejected from the connection pool.
- `consecutive_local_origin_failures` (Number)
- `interval` (String) Time interval between ejection sweep analysis.
- `max_ejection_percent` (Number)
- `min_health_percent` (Number)
- `split_external_local_origin_errors` (Boolean) Determines whether to distinguish local origin failures from external errors.


<a id="nestedatt--spec--traffic_policy--port_level_settings"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings`

Optional:

- `connection_pool` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--connection_pool))
- `load_balancer` (Attributes) Settings controlling the load balancer algorithms. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--load_balancer))
- `outlier_detection` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--outlier_detection))
- `port` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--port))
- `tls` (Attributes) TLS related settings for connections to the upstream service. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls))

<a id="nestedatt--spec--traffic_policy--port_level_settings--connection_pool"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls`

Optional:

- `http` (Attributes) HTTP connection pool settings. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--http))
- `tcp` (Attributes) Settings common to both HTTP and TCP upstream connections. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--tcp))

<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--http"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.http`

Optional:

- `h2_upgrade_policy` (String) Specify if http1.1 connection should be upgraded to http2 for the associated destination.
- `http1_max_pending_requests` (Number)
- `http2_max_requests` (Number) Maximum number of active requests to a destination.
- `idle_timeout` (String) The idle timeout for upstream connection pool connections.
- `max_requests_per_connection` (Number) Maximum number of requests per connection to a backend.
- `max_retries` (Number)
- `use_client_protocol` (Boolean) If set to true, client protocol will be preserved while initiating connection to backend.


<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--tcp"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.tcp`

Optional:

- `connect_timeout` (String) TCP connection timeout.
- `max_connection_duration` (String) The maximum duration of a connection.
- `max_connections` (Number) Maximum number of HTTP1 /TCP connections to a destination host.
- `tcp_keepalive` (Attributes) If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--tcp--tcp_keepalive))

<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--tcp--tcp_keepalive"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.tcp.tcp_keepalive`

Optional:

- `interval` (String) The time duration between keep-alive probes.
- `probes` (Number)
- `time` (String)




<a id="nestedatt--spec--traffic_policy--port_level_settings--load_balancer"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls`

Optional:

- `consistent_hash` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash))
- `locality_lb_setting` (Attributes) (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting))
- `simple` (String)
- `warmup_duration_secs` (String) Represents the warmup duration of Service.

<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.consistent_hash`

Optional:

- `http_cookie` (Attributes) Hash based on HTTP cookie. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--http_cookie))
- `http_header_name` (String) Hash based on a specific HTTP header.
- `http_query_parameter_name` (String) Hash based on a specific HTTP query parameter.
- `maglev` (Attributes) The Maglev load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--maglev))
- `minimum_ring_size` (Number) Deprecated.
- `ring_hash` (Attributes) The ring/modulo hash load balancer implements consistent hashing to backend hosts. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--ring_hash))
- `use_source_ip` (Boolean) Hash based on the source IP address.

<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--http_cookie"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.consistent_hash.use_source_ip`

Optional:

- `name` (String) Name of the cookie.
- `path` (String) Path to set for the cookie.
- `ttl` (String) Lifetime of the cookie.


<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--maglev"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.consistent_hash.use_source_ip`

Optional:

- `table_size` (Number) The table size for Maglev hashing.


<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--consistent_hash--ring_hash"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.consistent_hash.use_source_ip`

Optional:

- `minimum_ring_size` (Number)



<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.locality_lb_setting`

Optional:

- `distribute` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting--distribute))
- `enabled` (Boolean) enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.
- `failover` (Attributes List) Optional: only one of distribute, failover or failoverPriority can be set. (see [below for nested schema](#nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting--failover))
- `failover_priority` (List of String) failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.

<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting--distribute"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.locality_lb_setting.failover_priority`

Optional:

- `from` (String) Originating locality, '/' separated, e.g.
- `to` (Map of String) Map of upstream localities to traffic distribution weights.


<a id="nestedatt--spec--traffic_policy--port_level_settings--tls--locality_lb_setting--failover"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls.locality_lb_setting.failover_priority`

Optional:

- `from` (String) Originating region.
- `to` (String)




<a id="nestedatt--spec--traffic_policy--port_level_settings--outlier_detection"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls`

Optional:

- `base_ejection_time` (String) Minimum ejection duration.
- `consecutive5xx_errors` (Number) Number of 5xx errors before a host is ejected from the connection pool.
- `consecutive_errors` (Number)
- `consecutive_gateway_errors` (Number) Number of gateway errors before a host is ejected from the connection pool.
- `consecutive_local_origin_failures` (Number)
- `interval` (String) Time interval between ejection sweep analysis.
- `max_ejection_percent` (Number)
- `min_health_percent` (Number)
- `split_external_local_origin_errors` (Boolean) Determines whether to distinguish local origin failures from external errors.


<a id="nestedatt--spec--traffic_policy--port_level_settings--port"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls`

Optional:

- `number` (Number)


<a id="nestedatt--spec--traffic_policy--port_level_settings--tls"></a>
### Nested Schema for `spec.traffic_policy.port_level_settings.tls`

Optional:

- `ca_certificates` (String)
- `client_certificate` (String) REQUIRED if mode is 'MUTUAL'.
- `credential_name` (String)
- `insecure_skip_verify` (Boolean)
- `mode` (String)
- `private_key` (String) REQUIRED if mode is 'MUTUAL'.
- `sni` (String) SNI string to present to the server during TLS handshake.
- `subject_alt_names` (List of String)



<a id="nestedatt--spec--traffic_policy--tls"></a>
### Nested Schema for `spec.traffic_policy.tls`

Optional:

- `ca_certificates` (String)
- `client_certificate` (String) REQUIRED if mode is 'MUTUAL'.
- `credential_name` (String)
- `insecure_skip_verify` (Boolean)
- `mode` (String)
- `private_key` (String) REQUIRED if mode is 'MUTUAL'.
- `sni` (String) SNI string to present to the server during TLS handshake.
- `subject_alt_names` (List of String)


<a id="nestedatt--spec--traffic_policy--tunnel"></a>
### Nested Schema for `spec.traffic_policy.tunnel`

Optional:

- `protocol` (String) Specifies which protocol to use for tunneling the downstream connection.
- `target_host` (String) Specifies a host to which the downstream connection is tunneled.
- `target_port` (Number) Specifies a port to which the downstream connection is tunneled.



<a id="nestedatt--spec--workload_selector"></a>
### Nested Schema for `spec.workload_selector`

Optional:

- `match_labels` (Map of String)


