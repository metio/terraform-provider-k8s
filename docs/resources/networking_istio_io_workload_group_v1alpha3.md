---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networking_istio_io_workload_group_v1alpha3 Resource - terraform-provider-k8s"
subcategory: "networking.istio.io"
description: |-
  
---

# k8s_networking_istio_io_workload_group_v1alpha3 (Resource)



## Example Usage

```terraform
resource "k8s_networking_istio_io_workload_group_v1alpha3" "minimal" {
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

- `spec` (Attributes) Describes a collection of workload instances. See more details at: https://istio.io/docs/reference/config/networking/workload-group.html (see [below for nested schema](#nestedatt--spec))

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

- `metadata` (Attributes) Metadata that will be used for all corresponding 'WorkloadEntries'. (see [below for nested schema](#nestedatt--spec--metadata))
- `probe` (Attributes) 'ReadinessProbe' describes the configuration the user must provide for healthchecking on their workload. (see [below for nested schema](#nestedatt--spec--probe))
- `template` (Attributes) Template to be used for the generation of 'WorkloadEntry' resources that belong to this 'WorkloadGroup'. (see [below for nested schema](#nestedatt--spec--template))

<a id="nestedatt--spec--metadata"></a>
### Nested Schema for `spec.metadata`

Optional:

- `annotations` (Map of String)
- `labels` (Map of String)


<a id="nestedatt--spec--probe"></a>
### Nested Schema for `spec.probe`

Optional:

- `exec` (Attributes) Health is determined by how the command that is executed exited. (see [below for nested schema](#nestedatt--spec--probe--exec))
- `failure_threshold` (Number) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- `http_get` (Attributes) (see [below for nested schema](#nestedatt--spec--probe--http_get))
- `initial_delay_seconds` (Number) Number of seconds after the container has started before readiness probes are initiated.
- `period_seconds` (Number) How often (in seconds) to perform the probe.
- `success_threshold` (Number) Minimum consecutive successes for the probe to be considered successful after having failed.
- `tcp_socket` (Attributes) Health is determined by if the proxy is able to connect. (see [below for nested schema](#nestedatt--spec--probe--tcp_socket))
- `timeout_seconds` (Number) Number of seconds after which the probe times out.

<a id="nestedatt--spec--probe--exec"></a>
### Nested Schema for `spec.probe.exec`

Optional:

- `command` (List of String) Command to run.


<a id="nestedatt--spec--probe--http_get"></a>
### Nested Schema for `spec.probe.http_get`

Optional:

- `host` (String) Host name to connect to, defaults to the pod IP.
- `http_headers` (Attributes List) Headers the proxy will pass on to make the request. (see [below for nested schema](#nestedatt--spec--probe--http_get--http_headers))
- `path` (String) Path to access on the HTTP server.
- `port` (Number) Port on which the endpoint lives.
- `scheme` (String)

<a id="nestedatt--spec--probe--http_get--http_headers"></a>
### Nested Schema for `spec.probe.http_get.scheme`

Optional:

- `name` (String)
- `value` (String)



<a id="nestedatt--spec--probe--tcp_socket"></a>
### Nested Schema for `spec.probe.tcp_socket`

Optional:

- `host` (String)
- `port` (Number)



<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Optional:

- `address` (String)
- `labels` (Map of String) One or more labels associated with the endpoint.
- `locality` (String) The locality associated with the endpoint.
- `network` (String)
- `ports` (Map of String) Set of ports associated with the endpoint.
- `service_account` (String)
- `weight` (Number) The load balancing weight associated with the endpoint.


