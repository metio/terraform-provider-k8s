---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "gateway.nginx.org"
description: |-
  NginxProxy is a configuration object that is attached to a GatewayClass parametersRef. It provides a wayto configure global settings for all Gateways defined from the GatewayClass.
---

# k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest (Data Source)

NginxProxy is a configuration object that is attached to a GatewayClass parametersRef. It provides a wayto configure global settings for all Gateways defined from the GatewayClass.

## Example Usage

```terraform
data "k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of the NginxProxy. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `disable_http2` (Boolean) DisableHTTP2 defines if http2 should be disabled for all servers.Default is false, meaning http2 will be enabled for all servers.
- `ip_family` (String) IPFamily specifies the IP family to be used by the NGINX.Default is 'dual', meaning the server will use both IPv4 and IPv6.
- `telemetry` (Attributes) Telemetry specifies the OpenTelemetry configuration. (see [below for nested schema](#nestedatt--spec--telemetry))

<a id="nestedatt--spec--telemetry"></a>
### Nested Schema for `spec.telemetry`

Optional:

- `exporter` (Attributes) Exporter specifies OpenTelemetry export parameters. (see [below for nested schema](#nestedatt--spec--telemetry--exporter))
- `service_name` (String) ServiceName is the 'service.name' attribute of the OpenTelemetry resource.Default is 'ngf:<gateway-namespace>:<gateway-name>'. If a value is provided by the user,then the default becomes a prefix to that value.
- `span_attributes` (Attributes List) SpanAttributes are custom key/value attributes that are added to each span. (see [below for nested schema](#nestedatt--spec--telemetry--span_attributes))

<a id="nestedatt--spec--telemetry--exporter"></a>
### Nested Schema for `spec.telemetry.exporter`

Required:

- `endpoint` (String) Endpoint is the address of OTLP/gRPC endpoint that will accept telemetry data.Format: alphanumeric hostname with optional http scheme and optional port.

Optional:

- `batch_count` (Number) BatchCount is the number of pending batches per worker, spans exceeding the limit are dropped.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter
- `batch_size` (Number) BatchSize is the maximum number of spans to be sent in one batch per worker.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter
- `interval` (String) Interval is the maximum interval between two exports.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter


<a id="nestedatt--spec--telemetry--span_attributes"></a>
### Nested Schema for `spec.telemetry.span_attributes`

Required:

- `key` (String) Key is the key for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''
- `value` (String) Value is the value for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''
