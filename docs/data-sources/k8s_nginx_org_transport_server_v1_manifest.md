---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_k8s_nginx_org_transport_server_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "k8s.nginx.org"
description: |-
  TransportServer defines the TransportServer resource.
---

# k8s_k8s_nginx_org_transport_server_v1_manifest (Data Source)

TransportServer defines the TransportServer resource.

## Example Usage

```terraform
data "k8s_k8s_nginx_org_transport_server_v1_manifest" "example" {
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

- `spec` (Attributes) TransportServerSpec is the spec of the TransportServer resource. (see [below for nested schema](#nestedatt--spec))

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

- `action` (Attributes) TransportServerAction defines an action. (see [below for nested schema](#nestedatt--spec--action))
- `host` (String)
- `ingress_class_name` (String)
- `listener` (Attributes) TransportServerListener defines a listener for a TransportServer. (see [below for nested schema](#nestedatt--spec--listener))
- `server_snippets` (String)
- `session_parameters` (Attributes) SessionParameters defines session parameters. (see [below for nested schema](#nestedatt--spec--session_parameters))
- `stream_snippets` (String)
- `tls` (Attributes) TransportServerTLS defines TransportServerTLS configuration for a TransportServer. (see [below for nested schema](#nestedatt--spec--tls))
- `upstream_parameters` (Attributes) UpstreamParameters defines parameters for an upstream. (see [below for nested schema](#nestedatt--spec--upstream_parameters))
- `upstreams` (Attributes List) (see [below for nested schema](#nestedatt--spec--upstreams))

<a id="nestedatt--spec--action"></a>
### Nested Schema for `spec.action`

Optional:

- `pass` (String)


<a id="nestedatt--spec--listener"></a>
### Nested Schema for `spec.listener`

Optional:

- `name` (String)
- `protocol` (String)


<a id="nestedatt--spec--session_parameters"></a>
### Nested Schema for `spec.session_parameters`

Optional:

- `timeout` (String)


<a id="nestedatt--spec--tls"></a>
### Nested Schema for `spec.tls`

Optional:

- `secret` (String)


<a id="nestedatt--spec--upstream_parameters"></a>
### Nested Schema for `spec.upstream_parameters`

Optional:

- `connect_timeout` (String)
- `next_upstream` (Boolean)
- `next_upstream_timeout` (String)
- `next_upstream_tries` (Number)
- `udp_requests` (Number)
- `udp_responses` (Number)


<a id="nestedatt--spec--upstreams"></a>
### Nested Schema for `spec.upstreams`

Optional:

- `backup` (String)
- `backup_port` (Number)
- `fail_timeout` (String)
- `health_check` (Attributes) TransportServerHealthCheck defines the parameters for active Upstream HealthChecks. (see [below for nested schema](#nestedatt--spec--upstreams--health_check))
- `load_balancing_method` (String)
- `max_conns` (Number)
- `max_fails` (Number)
- `name` (String)
- `port` (Number)
- `service` (String)

<a id="nestedatt--spec--upstreams--health_check"></a>
### Nested Schema for `spec.upstreams.health_check`

Optional:

- `enable` (Boolean)
- `fails` (Number)
- `interval` (String)
- `jitter` (String)
- `match` (Attributes) TransportServerMatch defines the parameters of a custom health check. (see [below for nested schema](#nestedatt--spec--upstreams--health_check--match))
- `passes` (Number)
- `port` (Number)
- `timeout` (String)

<a id="nestedatt--spec--upstreams--health_check--match"></a>
### Nested Schema for `spec.upstreams.health_check.match`

Optional:

- `expect` (String)
- `send` (String)
