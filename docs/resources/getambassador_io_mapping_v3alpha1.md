---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_getambassador_io_mapping_v3alpha1 Resource - terraform-provider-k8s"
subcategory: "getambassador.io"
description: |-
  Mapping is the Schema for the mappings API
---

# k8s_getambassador_io_mapping_v3alpha1 (Resource)

Mapping is the Schema for the mappings API

## Example Usage

```terraform
resource "k8s_getambassador_io_mapping_v3alpha1" "minimal" {
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

- `spec` (Attributes) MappingSpec defines the desired state of Mapping (see [below for nested schema](#nestedatt--spec))

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

Required:

- `prefix` (String)
- `service` (String)

Optional:

- `add_linkerd_headers` (Boolean)
- `add_request_headers` (Attributes) (see [below for nested schema](#nestedatt--spec--add_request_headers))
- `add_response_headers` (Attributes) (see [below for nested schema](#nestedatt--spec--add_response_headers))
- `allow_upgrade` (List of String) A case-insensitive list of the non-HTTP protocols to allow 'upgrading' to from HTTP via the 'Connection: upgrade' mechanism[1].  After the upgrade, Ambassador does not interpret the traffic, and behaves similarly to how it does for TCPMappings.  [1]: https://tools.ietf.org/html/rfc7230#section-6.7  For example, if your upstream service supports WebSockets, you would write     allow_upgrade:    - websocket  Or if your upstream service supports upgrading from HTTP to SPDY (as the Kubernetes apiserver does for 'kubectl exec' functionality), you would write     allow_upgrade:    - spdy/3.1
- `ambassador_id` (List of String) AmbassadorID declares which Ambassador instances should pay attention to this resource. If no value is provided, the default is:  	ambassador_id: 	- 'default'  TODO(lukeshu): In v3alpha2, consider renaming all of the 'ambassador_id' (singular) fields to 'ambassador_ids' (plural).
- `auth_context_extensions` (Map of String)
- `auto_host_rewrite` (Boolean)
- `bypass_auth` (Boolean)
- `bypass_error_response_overrides` (Boolean) If true, bypasses any 'error_response_overrides' set on the Ambassador module.
- `case_sensitive` (Boolean)
- `circuit_breakers` (Attributes List) (see [below for nested schema](#nestedatt--spec--circuit_breakers))
- `cluster_idle_timeout_ms` (Number) TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.
- `cluster_max_connection_lifetime_ms` (Number) TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.
- `cluster_tag` (String)
- `connect_timeout_ms` (Number) TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.
- `cors` (Attributes) (see [below for nested schema](#nestedatt--spec--cors))
- `dns_type` (String)
- `docs` (Attributes) DocsInfo provides some extra information about the docs for the Mapping. Docs is used by both the agent and the DevPortal. (see [below for nested schema](#nestedatt--spec--docs))
- `enable_ipv4` (Boolean)
- `enable_ipv6` (Boolean)
- `envoy_override` (Dynamic) UntypedDict is relatively opaque as a Go type, but it preserves its contents in a roundtrippable way.
- `error_response_overrides` (Attributes List) Error response overrides for this Mapping. Replaces all of the 'error_response_overrides' set on the Ambassador module, if any. (see [below for nested schema](#nestedatt--spec--error_response_overrides))
- `grpc` (Boolean)
- `headers` (Map of String)
- `host` (String) Exact match for the hostname of a request if HostRegex is false; regex match for the hostname if HostRegex is true.  Host specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Host will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.  DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.
- `host_redirect` (Boolean)
- `host_regex` (Boolean) DEPRECATED: Host is either an exact match or a regex, depending on HostRegex. Use HostName instead.  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.host and MappingSpec.host_regex in favor of a MappingSpec.deprecated_hostname_regex.
- `host_rewrite` (String)
- `hostname` (String) Hostname is a DNS glob specifying the hosts to which this Mapping applies.  Hostname specifies both a match for the ':authority' header of a request, as well as a match criterion for Host CRDs: a Mapping that specifies Hostname will not associate with a Host that doesn't have a matching Hostname.  If both Host and Hostname are set, an error is logged, Host is ignored, and Hostname is used.
- `idle_timeout_ms` (Number) TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.
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
- `remove_request_headers` (List of String)
- `remove_response_headers` (List of String)
- `resolver` (String)
- `respect_dns_ttl` (Boolean)
- `retry_policy` (Attributes) (see [below for nested schema](#nestedatt--spec--retry_policy))
- `rewrite` (String)
- `shadow` (Boolean)
- `stats_name` (String)
- `timeout_ms` (Number) The timeout for requests that use this Mapping. Overrides 'cluster_request_timeout_ms' set on the Ambassador Module, if it exists.
- `tls` (String)
- `use_websocket` (Boolean) use_websocket is deprecated, and is equivlaent to setting 'allow_upgrade: ['websocket']'  TODO(lukeshu): In v3alpha2, get rid of MappingSpec.DeprecatedUseWebsocket.
- `v2_bool_headers` (List of String)
- `v2_bool_query_parameters` (List of String)
- `v2_explicit_tls` (Attributes) V2ExplicitTLS controls some vanity/stylistic elements when converting from v3alpha1 to v2.  The values in an V2ExplicitTLS should not in any way affect the runtime operation of Emissary; except that it may affect internal names in the Envoy config, which may in turn affect stats names.  But it should not affect any end-user observable behavior. (see [below for nested schema](#nestedatt--spec--v2_explicit_tls))
- `weight` (Number)

<a id="nestedatt--spec--add_request_headers"></a>
### Nested Schema for `spec.add_request_headers`

Optional:

- `append` (Boolean)
- `v2_representation` (String)
- `value` (String)


<a id="nestedatt--spec--add_response_headers"></a>
### Nested Schema for `spec.add_response_headers`

Optional:

- `append` (Boolean)
- `v2_representation` (String)
- `value` (String)


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
- `exposed_headers` (List of String)
- `headers` (List of String)
- `max_age` (String)
- `methods` (List of String)
- `origins` (List of String)
- `v2_comma_separated_origins` (Boolean)


<a id="nestedatt--spec--docs"></a>
### Nested Schema for `spec.docs`

Optional:

- `display_name` (String)
- `ignored` (Boolean)
- `path` (String)
- `timeout_ms` (Number) TODO(lukeshu): In v3alpha2, change all of the '{foo}_ms'/'MillisecondDuration' fields to '{foo}'/'metav1.Duration'.
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


<a id="nestedatt--spec--v2_explicit_tls"></a>
### Nested Schema for `spec.v2_explicit_tls`

Optional:

- `service_scheme` (String) ServiceScheme specifies how to spell and capitalize the scheme-part of the service URL.  Acceptable values are 'http://' (case-insensitive), 'https://' (case-insensitive), or ''.  The value is used if it agrees with whether or not this resource enables TLS origination, or if something else in the resource overrides the scheme.
- `tls` (String) TLS controls whether and how to represent the 'tls' field when its value could be implied by the 'service' field.  In v2, there were a lot of different ways to spell an 'empty' value, and this field specifies which way to spell it (and will therefore only be used if the value will indeed be empty).   | Value        | Representation                        | Meaning of representation          |  |--------------+---------------------------------------+------------------------------------|  | ''           | omit the field                        | defer to service (no TLSContext)   |  | 'null'       | store an explicit 'null' in the field | defer to service (no TLSContext)   |  | 'string'     | store an empty string in the field    | defer to service (no TLSContext)   |  | 'bool:false' | store a Boolean 'false' in the field  | defer to service (no TLSContext)   |  | 'bool:true'  | store a Boolean 'true' in the field   | originate TLS (no TLSContext)      |  If the meaning of the representation contradicts anything else (if a TLSContext is to be used, or in the case of 'bool:true' if TLS is not to be originated), then this field is ignored.


