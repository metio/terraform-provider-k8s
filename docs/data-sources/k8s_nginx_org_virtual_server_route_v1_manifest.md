---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_k8s_nginx_org_virtual_server_route_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "k8s.nginx.org"
description: |-
  VirtualServerRoute defines the VirtualServerRoute resource.
---

# k8s_k8s_nginx_org_virtual_server_route_v1_manifest (Data Source)

VirtualServerRoute defines the VirtualServerRoute resource.

## Example Usage

```terraform
data "k8s_k8s_nginx_org_virtual_server_route_v1_manifest" "example" {
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

- `spec` (Attributes) VirtualServerRouteSpec is the spec of the VirtualServerRoute resource. (see [below for nested schema](#nestedatt--spec))

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

- `host` (String)
- `ingress_class_name` (String)
- `subroutes` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes))
- `upstreams` (Attributes List) (see [below for nested schema](#nestedatt--spec--upstreams))

<a id="nestedatt--spec--subroutes"></a>
### Nested Schema for `spec.subroutes`

Optional:

- `action` (Attributes) Action defines an action. (see [below for nested schema](#nestedatt--spec--subroutes--action))
- `dos` (String)
- `error_pages` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--error_pages))
- `location_snippets` (String)
- `matches` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches))
- `path` (String)
- `policies` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--policies))
- `route` (String)
- `splits` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--splits))

<a id="nestedatt--spec--subroutes--action"></a>
### Nested Schema for `spec.subroutes.action`

Optional:

- `pass` (String)
- `proxy` (Attributes) ActionProxy defines a proxy in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--action--proxy))
- `redirect` (Attributes) ActionRedirect defines a redirect in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--action--redirect))
- `return` (Attributes) ActionReturn defines a return in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--action--return))

<a id="nestedatt--spec--subroutes--action--proxy"></a>
### Nested Schema for `spec.subroutes.action.proxy`

Optional:

- `request_headers` (Attributes) ProxyRequestHeaders defines the request headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--action--return--request_headers))
- `response_headers` (Attributes) ProxyResponseHeaders defines the response headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--action--return--response_headers))
- `rewrite_path` (String)
- `upstream` (String)

<a id="nestedatt--spec--subroutes--action--return--request_headers"></a>
### Nested Schema for `spec.subroutes.action.return.request_headers`

Optional:

- `pass` (Boolean)
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--action--return--request_headers--set))

<a id="nestedatt--spec--subroutes--action--return--request_headers--set"></a>
### Nested Schema for `spec.subroutes.action.return.request_headers.set`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--subroutes--action--return--response_headers"></a>
### Nested Schema for `spec.subroutes.action.return.response_headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--action--return--response_headers--add))
- `hide` (List of String)
- `ignore` (List of String)
- `pass` (List of String)

<a id="nestedatt--spec--subroutes--action--return--response_headers--add"></a>
### Nested Schema for `spec.subroutes.action.return.response_headers.add`

Optional:

- `always` (Boolean)
- `name` (String)
- `value` (String)




<a id="nestedatt--spec--subroutes--action--redirect"></a>
### Nested Schema for `spec.subroutes.action.redirect`

Optional:

- `code` (Number)
- `url` (String)


<a id="nestedatt--spec--subroutes--action--return"></a>
### Nested Schema for `spec.subroutes.action.return`

Optional:

- `body` (String)
- `code` (Number)
- `type` (String)



<a id="nestedatt--spec--subroutes--error_pages"></a>
### Nested Schema for `spec.subroutes.error_pages`

Optional:

- `codes` (List of String)
- `redirect` (Attributes) ErrorPageRedirect defines a redirect for an ErrorPage. (see [below for nested schema](#nestedatt--spec--subroutes--error_pages--redirect))
- `return` (Attributes) ErrorPageReturn defines a return for an ErrorPage. (see [below for nested schema](#nestedatt--spec--subroutes--error_pages--return))

<a id="nestedatt--spec--subroutes--error_pages--redirect"></a>
### Nested Schema for `spec.subroutes.error_pages.redirect`

Optional:

- `code` (Number)
- `url` (String)


<a id="nestedatt--spec--subroutes--error_pages--return"></a>
### Nested Schema for `spec.subroutes.error_pages.return`

Optional:

- `body` (String)
- `code` (Number)
- `headers` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--error_pages--return--headers))
- `type` (String)

<a id="nestedatt--spec--subroutes--error_pages--return--headers"></a>
### Nested Schema for `spec.subroutes.error_pages.return.headers`

Optional:

- `name` (String)
- `value` (String)




<a id="nestedatt--spec--subroutes--matches"></a>
### Nested Schema for `spec.subroutes.matches`

Optional:

- `action` (Attributes) Action defines an action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--action))
- `conditions` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--conditions))
- `splits` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits))

<a id="nestedatt--spec--subroutes--matches--action"></a>
### Nested Schema for `spec.subroutes.matches.action`

Optional:

- `pass` (String)
- `proxy` (Attributes) ActionProxy defines a proxy in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--proxy))
- `redirect` (Attributes) ActionRedirect defines a redirect in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--redirect))
- `return` (Attributes) ActionReturn defines a return in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--return))

<a id="nestedatt--spec--subroutes--matches--splits--proxy"></a>
### Nested Schema for `spec.subroutes.matches.splits.proxy`

Optional:

- `request_headers` (Attributes) ProxyRequestHeaders defines the request headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--proxy--request_headers))
- `response_headers` (Attributes) ProxyResponseHeaders defines the response headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--proxy--response_headers))
- `rewrite_path` (String)
- `upstream` (String)

<a id="nestedatt--spec--subroutes--matches--splits--proxy--request_headers"></a>
### Nested Schema for `spec.subroutes.matches.splits.proxy.request_headers`

Optional:

- `pass` (Boolean)
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--proxy--upstream--set))

<a id="nestedatt--spec--subroutes--matches--splits--proxy--upstream--set"></a>
### Nested Schema for `spec.subroutes.matches.splits.proxy.upstream.set`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--subroutes--matches--splits--proxy--response_headers"></a>
### Nested Schema for `spec.subroutes.matches.splits.proxy.response_headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--proxy--upstream--add))
- `hide` (List of String)
- `ignore` (List of String)
- `pass` (List of String)

<a id="nestedatt--spec--subroutes--matches--splits--proxy--upstream--add"></a>
### Nested Schema for `spec.subroutes.matches.splits.proxy.upstream.add`

Optional:

- `always` (Boolean)
- `name` (String)
- `value` (String)




<a id="nestedatt--spec--subroutes--matches--splits--redirect"></a>
### Nested Schema for `spec.subroutes.matches.splits.redirect`

Optional:

- `code` (Number)
- `url` (String)


<a id="nestedatt--spec--subroutes--matches--splits--return"></a>
### Nested Schema for `spec.subroutes.matches.splits.return`

Optional:

- `body` (String)
- `code` (Number)
- `type` (String)



<a id="nestedatt--spec--subroutes--matches--conditions"></a>
### Nested Schema for `spec.subroutes.matches.conditions`

Optional:

- `argument` (String)
- `cookie` (String)
- `header` (String)
- `value` (String)
- `variable` (String)


<a id="nestedatt--spec--subroutes--matches--splits"></a>
### Nested Schema for `spec.subroutes.matches.splits`

Optional:

- `action` (Attributes) Action defines an action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action))
- `weight` (Number)

<a id="nestedatt--spec--subroutes--matches--splits--action"></a>
### Nested Schema for `spec.subroutes.matches.splits.action`

Optional:

- `pass` (String)
- `proxy` (Attributes) ActionProxy defines a proxy in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--proxy))
- `redirect` (Attributes) ActionRedirect defines a redirect in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--redirect))
- `return` (Attributes) ActionReturn defines a return in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--return))

<a id="nestedatt--spec--subroutes--matches--splits--action--proxy"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.proxy`

Optional:

- `request_headers` (Attributes) ProxyRequestHeaders defines the request headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--return--request_headers))
- `response_headers` (Attributes) ProxyResponseHeaders defines the response headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--return--response_headers))
- `rewrite_path` (String)
- `upstream` (String)

<a id="nestedatt--spec--subroutes--matches--splits--action--return--request_headers"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.return.request_headers`

Optional:

- `pass` (Boolean)
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--return--upstream--set))

<a id="nestedatt--spec--subroutes--matches--splits--action--return--upstream--set"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.return.upstream.set`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--subroutes--matches--splits--action--return--response_headers"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.return.response_headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--matches--splits--action--return--upstream--add))
- `hide` (List of String)
- `ignore` (List of String)
- `pass` (List of String)

<a id="nestedatt--spec--subroutes--matches--splits--action--return--upstream--add"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.return.upstream.add`

Optional:

- `always` (Boolean)
- `name` (String)
- `value` (String)




<a id="nestedatt--spec--subroutes--matches--splits--action--redirect"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.redirect`

Optional:

- `code` (Number)
- `url` (String)


<a id="nestedatt--spec--subroutes--matches--splits--action--return"></a>
### Nested Schema for `spec.subroutes.matches.splits.action.return`

Optional:

- `body` (String)
- `code` (Number)
- `type` (String)





<a id="nestedatt--spec--subroutes--policies"></a>
### Nested Schema for `spec.subroutes.policies`

Optional:

- `name` (String)
- `namespace` (String)


<a id="nestedatt--spec--subroutes--splits"></a>
### Nested Schema for `spec.subroutes.splits`

Optional:

- `action` (Attributes) Action defines an action. (see [below for nested schema](#nestedatt--spec--subroutes--splits--action))
- `weight` (Number)

<a id="nestedatt--spec--subroutes--splits--action"></a>
### Nested Schema for `spec.subroutes.splits.action`

Optional:

- `pass` (String)
- `proxy` (Attributes) ActionProxy defines a proxy in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--proxy))
- `redirect` (Attributes) ActionRedirect defines a redirect in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--redirect))
- `return` (Attributes) ActionReturn defines a return in an Action. (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--return))

<a id="nestedatt--spec--subroutes--splits--weight--proxy"></a>
### Nested Schema for `spec.subroutes.splits.weight.proxy`

Optional:

- `request_headers` (Attributes) ProxyRequestHeaders defines the request headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--proxy--request_headers))
- `response_headers` (Attributes) ProxyResponseHeaders defines the response headers manipulation in an ActionProxy. (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--proxy--response_headers))
- `rewrite_path` (String)
- `upstream` (String)

<a id="nestedatt--spec--subroutes--splits--weight--proxy--request_headers"></a>
### Nested Schema for `spec.subroutes.splits.weight.proxy.request_headers`

Optional:

- `pass` (Boolean)
- `set` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--proxy--upstream--set))

<a id="nestedatt--spec--subroutes--splits--weight--proxy--upstream--set"></a>
### Nested Schema for `spec.subroutes.splits.weight.proxy.upstream.set`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--subroutes--splits--weight--proxy--response_headers"></a>
### Nested Schema for `spec.subroutes.splits.weight.proxy.response_headers`

Optional:

- `add` (Attributes List) (see [below for nested schema](#nestedatt--spec--subroutes--splits--weight--proxy--upstream--add))
- `hide` (List of String)
- `ignore` (List of String)
- `pass` (List of String)

<a id="nestedatt--spec--subroutes--splits--weight--proxy--upstream--add"></a>
### Nested Schema for `spec.subroutes.splits.weight.proxy.upstream.add`

Optional:

- `always` (Boolean)
- `name` (String)
- `value` (String)




<a id="nestedatt--spec--subroutes--splits--weight--redirect"></a>
### Nested Schema for `spec.subroutes.splits.weight.redirect`

Optional:

- `code` (Number)
- `url` (String)


<a id="nestedatt--spec--subroutes--splits--weight--return"></a>
### Nested Schema for `spec.subroutes.splits.weight.return`

Optional:

- `body` (String)
- `code` (Number)
- `type` (String)





<a id="nestedatt--spec--upstreams"></a>
### Nested Schema for `spec.upstreams`

Optional:

- `backup` (String)
- `backup_port` (Number)
- `buffer_size` (String)
- `buffering` (Boolean)
- `buffers` (Attributes) UpstreamBuffers defines Buffer Configuration for an Upstream. (see [below for nested schema](#nestedatt--spec--upstreams--buffers))
- `client_max_body_size` (String)
- `connect_timeout` (String)
- `fail_timeout` (String)
- `health_check` (Attributes) HealthCheck defines the parameters for active Upstream HealthChecks. (see [below for nested schema](#nestedatt--spec--upstreams--health_check))
- `keepalive` (Number)
- `lb_method` (String)
- `max_conns` (Number)
- `max_fails` (Number)
- `name` (String)
- `next_upstream` (String)
- `next_upstream_timeout` (String)
- `next_upstream_tries` (Number)
- `ntlm` (Boolean)
- `port` (Number)
- `queue` (Attributes) UpstreamQueue defines Queue Configuration for an Upstream. (see [below for nested schema](#nestedatt--spec--upstreams--queue))
- `read_timeout` (String)
- `send_timeout` (String)
- `service` (String)
- `session_cookie` (Attributes) SessionCookie defines the parameters for session persistence. (see [below for nested schema](#nestedatt--spec--upstreams--session_cookie))
- `slow_start` (String)
- `subselector` (Map of String)
- `tls` (Attributes) UpstreamTLS defines a TLS configuration for an Upstream. (see [below for nested schema](#nestedatt--spec--upstreams--tls))
- `type` (String)
- `use_cluster_ip` (Boolean)

<a id="nestedatt--spec--upstreams--buffers"></a>
### Nested Schema for `spec.upstreams.buffers`

Optional:

- `number` (Number)
- `size` (String)


<a id="nestedatt--spec--upstreams--health_check"></a>
### Nested Schema for `spec.upstreams.health_check`

Optional:

- `connect_timeout` (String)
- `enable` (Boolean)
- `fails` (Number)
- `grpc_service` (String)
- `grpc_status` (Number)
- `headers` (Attributes List) (see [below for nested schema](#nestedatt--spec--upstreams--health_check--headers))
- `interval` (String)
- `jitter` (String)
- `keepalive_time` (String)
- `mandatory` (Boolean)
- `passes` (Number)
- `path` (String)
- `persistent` (Boolean)
- `port` (Number)
- `read_timeout` (String)
- `send_timeout` (String)
- `status_match` (String)
- `tls` (Attributes) UpstreamTLS defines a TLS configuration for an Upstream. (see [below for nested schema](#nestedatt--spec--upstreams--health_check--tls))

<a id="nestedatt--spec--upstreams--health_check--headers"></a>
### Nested Schema for `spec.upstreams.health_check.headers`

Optional:

- `name` (String)
- `value` (String)


<a id="nestedatt--spec--upstreams--health_check--tls"></a>
### Nested Schema for `spec.upstreams.health_check.tls`

Optional:

- `enable` (Boolean)



<a id="nestedatt--spec--upstreams--queue"></a>
### Nested Schema for `spec.upstreams.queue`

Optional:

- `size` (Number)
- `timeout` (String)


<a id="nestedatt--spec--upstreams--session_cookie"></a>
### Nested Schema for `spec.upstreams.session_cookie`

Optional:

- `domain` (String)
- `enable` (Boolean)
- `expires` (String)
- `http_only` (Boolean)
- `name` (String)
- `path` (String)
- `samesite` (String)
- `secure` (Boolean)


<a id="nestedatt--spec--upstreams--tls"></a>
### Nested Schema for `spec.upstreams.tls`

Optional:

- `enable` (Boolean)