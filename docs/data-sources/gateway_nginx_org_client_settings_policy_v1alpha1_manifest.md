---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "gateway.nginx.org"
description: |-
  ClientSettingsPolicy is an Inherited Attached Policy. It provides a way to configure the behavior of the connection between the client and NGINX Gateway Fabric.
---

# k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest (Data Source)

ClientSettingsPolicy is an Inherited Attached Policy. It provides a way to configure the behavior of the connection between the client and NGINX Gateway Fabric.

## Example Usage

```terraform
data "k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      kind  = "Service"
      group = "v1"
      name  = "some-service"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of the ClientSettingsPolicy. (see [below for nested schema](#nestedatt--spec))

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

- `target_ref` (Attributes) TargetRef identifies an API object to apply the policy to. Object must be in the same namespace as the policy. Support: Gateway, HTTPRoute, GRPCRoute. (see [below for nested schema](#nestedatt--spec--target_ref))

Optional:

- `body` (Attributes) Body defines the client request body settings. (see [below for nested schema](#nestedatt--spec--body))
- `keep_alive` (Attributes) KeepAlive defines the keep-alive settings. (see [below for nested schema](#nestedatt--spec--keep_alive))

<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Required:

- `group` (String) Group is the group of the target resource.
- `kind` (String) Kind is kind of the target resource.
- `name` (String) Name is the name of the target resource.


<a id="nestedatt--spec--body"></a>
### Nested Schema for `spec.body`

Optional:

- `max_size` (String) MaxSize sets the maximum allowed size of the client request body. If the size in a request exceeds the configured value, the 413 (Request Entity Too Large) error is returned to the client. Setting size to 0 disables checking of client request body size. Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size.
- `timeout` (String) Timeout defines a timeout for reading client request body. The timeout is set only for a period between two successive read operations, not for the transmission of the whole request body. If a client does not transmit anything within this time, the request is terminated with the 408 (Request Time-out) error. Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout.


<a id="nestedatt--spec--keep_alive"></a>
### Nested Schema for `spec.keep_alive`

Optional:

- `requests` (Number) Requests sets the maximum number of requests that can be served through one keep-alive connection. After the maximum number of requests are made, the connection is closed. Closing connections periodically is necessary to free per-connection memory allocations. Therefore, using too high maximum number of requests is not recommended as it can lead to excessive memory usage. Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests.
- `time` (String) Time defines the maximum time during which requests can be processed through one keep-alive connection. After this time is reached, the connection is closed following the subsequent request processing. Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_time.
- `timeout` (Attributes) Timeout defines the keep-alive timeouts for clients. (see [below for nested schema](#nestedatt--spec--keep_alive--timeout))

<a id="nestedatt--spec--keep_alive--timeout"></a>
### Nested Schema for `spec.keep_alive.timeout`

Optional:

- `header` (String) Header sets the timeout in the 'Keep-Alive: timeout=time' response header field.
- `server` (String) Server sets the timeout during which a keep-alive client connection will stay open on the server side. Setting this value to 0 disables keep-alive client connections.
