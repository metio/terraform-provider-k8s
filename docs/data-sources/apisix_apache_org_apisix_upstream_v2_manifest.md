---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apisix_apache_org_apisix_upstream_v2_manifest Data Source - terraform-provider-k8s"
subcategory: "apisix.apache.org"
description: |-
  
---

# k8s_apisix_apache_org_apisix_upstream_v2_manifest (Data Source)



## Example Usage

```terraform
data "k8s_apisix_apache_org_apisix_upstream_v2_manifest" "example" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `discovery` (Attributes) Discovery is used to configure service discovery for upstream (see [below for nested schema](#nestedatt--spec--discovery))
- `external_nodes` (Attributes List) ExternalNodes contains external nodes the Upstream should use If this field is set, the upstream will use these nodes directly without any further resolves (see [below for nested schema](#nestedatt--spec--external_nodes))
- `health_check` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check))
- `ingress_class_name` (String)
- `loadbalancer` (Attributes) (see [below for nested schema](#nestedatt--spec--loadbalancer))
- `pass_host` (String)
- `port_level_settings` (Attributes List) (see [below for nested schema](#nestedatt--spec--port_level_settings))
- `retries` (Number)
- `scheme` (String)
- `subsets` (Attributes List) (see [below for nested schema](#nestedatt--spec--subsets))
- `timeout` (Attributes) (see [below for nested schema](#nestedatt--spec--timeout))
- `tls_secret` (Attributes) ApisixSecret describes the Kubernetes Secret name and namespace. (see [below for nested schema](#nestedatt--spec--tls_secret))
- `upstream_host` (String)

<a id="nestedatt--spec--discovery"></a>
### Nested Schema for `spec.discovery`

Optional:

- `args` (Map of String)
- `service_name` (String)
- `type` (String)


<a id="nestedatt--spec--external_nodes"></a>
### Nested Schema for `spec.external_nodes`

Optional:

- `name` (String)
- `port` (Number)
- `type` (String)
- `weight` (Number)


<a id="nestedatt--spec--health_check"></a>
### Nested Schema for `spec.health_check`

Optional:

- `active` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--active))
- `passive` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--passive))

<a id="nestedatt--spec--health_check--active"></a>
### Nested Schema for `spec.health_check.active`

Optional:

- `concurrency` (Number)
- `healthy` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--active--healthy))
- `host` (String)
- `http_path` (String)
- `port` (Number)
- `request_headers` (List of String)
- `strict_tls` (Boolean)
- `timeout` (Number)
- `type` (String)
- `unhealthy` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--active--unhealthy))

<a id="nestedatt--spec--health_check--active--healthy"></a>
### Nested Schema for `spec.health_check.active.healthy`

Optional:

- `http_codes` (List of String)
- `interval` (String)
- `successes` (Number)


<a id="nestedatt--spec--health_check--active--unhealthy"></a>
### Nested Schema for `spec.health_check.active.unhealthy`

Optional:

- `http_codes` (List of String)
- `http_failures` (Number)
- `interval` (String)
- `tcp_failures` (Number)
- `timeouts` (Number)



<a id="nestedatt--spec--health_check--passive"></a>
### Nested Schema for `spec.health_check.passive`

Optional:

- `healthy` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--passive--healthy))
- `type` (String)
- `unhealthy` (Attributes) (see [below for nested schema](#nestedatt--spec--health_check--passive--unhealthy))

<a id="nestedatt--spec--health_check--passive--healthy"></a>
### Nested Schema for `spec.health_check.passive.healthy`

Optional:

- `http_codes` (List of String)
- `successes` (Number)


<a id="nestedatt--spec--health_check--passive--unhealthy"></a>
### Nested Schema for `spec.health_check.passive.unhealthy`

Optional:

- `http_codes` (List of String)
- `http_failures` (Number)
- `tcp_failures` (Number)
- `timeouts` (Number)




<a id="nestedatt--spec--loadbalancer"></a>
### Nested Schema for `spec.loadbalancer`

Required:

- `type` (String)

Optional:

- `hash_on` (String)
- `key` (String)


<a id="nestedatt--spec--port_level_settings"></a>
### Nested Schema for `spec.port_level_settings`

Optional:

- `health_check` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check))
- `loadbalancer` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--loadbalancer))
- `port` (Number)
- `retries` (Number)
- `scheme` (String)
- `timeout` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--timeout))

<a id="nestedatt--spec--port_level_settings--health_check"></a>
### Nested Schema for `spec.port_level_settings.health_check`

Optional:

- `active` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--active))
- `passive` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--passive))

<a id="nestedatt--spec--port_level_settings--health_check--active"></a>
### Nested Schema for `spec.port_level_settings.health_check.active`

Optional:

- `concurrency` (Number)
- `healthy` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--active--healthy))
- `host` (String)
- `http_path` (String)
- `port` (Number)
- `request_headers` (List of String)
- `strict_tls` (Boolean)
- `timeout` (Number)
- `type` (String)
- `unhealthy` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--active--unhealthy))

<a id="nestedatt--spec--port_level_settings--health_check--active--healthy"></a>
### Nested Schema for `spec.port_level_settings.health_check.active.healthy`

Optional:

- `http_codes` (List of String)
- `interval` (String)
- `successes` (Number)


<a id="nestedatt--spec--port_level_settings--health_check--active--unhealthy"></a>
### Nested Schema for `spec.port_level_settings.health_check.active.unhealthy`

Optional:

- `http_codes` (List of String)
- `http_failures` (Number)
- `interval` (String)
- `tcp_failures` (Number)
- `timeout` (String)



<a id="nestedatt--spec--port_level_settings--health_check--passive"></a>
### Nested Schema for `spec.port_level_settings.health_check.passive`

Optional:

- `healthy` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--passive--healthy))
- `type` (String)
- `unhealthy` (Attributes) (see [below for nested schema](#nestedatt--spec--port_level_settings--health_check--passive--unhealthy))

<a id="nestedatt--spec--port_level_settings--health_check--passive--healthy"></a>
### Nested Schema for `spec.port_level_settings.health_check.passive.healthy`

Optional:

- `http_codes` (List of String)
- `successes` (Number)


<a id="nestedatt--spec--port_level_settings--health_check--passive--unhealthy"></a>
### Nested Schema for `spec.port_level_settings.health_check.passive.unhealthy`

Optional:

- `http_codes` (List of String)
- `http_failures` (Number)
- `tcp_failures` (Number)
- `timeout` (String)




<a id="nestedatt--spec--port_level_settings--loadbalancer"></a>
### Nested Schema for `spec.port_level_settings.loadbalancer`

Required:

- `type` (String)

Optional:

- `hash_on` (String)
- `key` (String)


<a id="nestedatt--spec--port_level_settings--timeout"></a>
### Nested Schema for `spec.port_level_settings.timeout`

Optional:

- `connect` (String)
- `read` (String)
- `send` (String)



<a id="nestedatt--spec--subsets"></a>
### Nested Schema for `spec.subsets`

Required:

- `labels` (Map of String)
- `name` (String)


<a id="nestedatt--spec--timeout"></a>
### Nested Schema for `spec.timeout`

Optional:

- `connect` (String)
- `read` (String)
- `send` (String)


<a id="nestedatt--spec--tls_secret"></a>
### Nested Schema for `spec.tls_secret`

Required:

- `name` (String)
- `namespace` (String)
