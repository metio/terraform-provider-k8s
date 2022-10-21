---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cilium_io_cilium_clusterwide_envoy_config_v2 Resource - terraform-provider-k8s"
subcategory: "cilium.io/v2"
description: |-
  
---

# k8s_cilium_io_cilium_clusterwide_envoy_config_v2 (Resource)



## Example Usage

```terraform
resource "k8s_cilium_io_cilium_clusterwide_envoy_config_v2" "minimal" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `backend_services` (Attributes List) BackendServices specifies Kubernetes services whose backends are automatically synced to Envoy using EDS.  Traffic for these services is not forwarded to an Envoy listener. This allows an Envoy listener load balance traffic to these backends while normal Cilium service load balancing takes care of balancing traffic for these services at the same time. (see [below for nested schema](#nestedatt--spec--backend_services))
- `resources` (List of Map of String) Envoy xDS resources, a list of the following Envoy resource types: type.googleapis.com/envoy.config.listener.v3.Listener, type.googleapis.com/envoy.config.route.v3.RouteConfiguration, type.googleapis.com/envoy.config.cluster.v3.Cluster, type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment, and type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret.
- `services` (Attributes List) Services specifies Kubernetes services for which traffic is forwarded to an Envoy listener for L7 load balancing. Backends of these services are automatically synced to Envoy usign EDS. (see [below for nested schema](#nestedatt--spec--services))

<a id="nestedatt--spec--backend_services"></a>
### Nested Schema for `spec.backend_services`

Required:

- `name` (String) Name is the name of a destination Kubernetes service that identifies traffic to be redirected.

Optional:

- `namespace` (String) Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace defaults to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.
- `number` (List of String) Port is the port number, which can be used for filtering in case of underlying is exposing multiple port numbers.


<a id="nestedatt--spec--services"></a>
### Nested Schema for `spec.services`

Required:

- `name` (String) Name is the name of a destination Kubernetes service that identifies traffic to be redirected.

Optional:

- `listener` (String) Listener specifies the name of the Envoy listener the service traffic is redirected to. The listener must be specified in the Envoy 'resources' of the same CiliumEnvoyConfig.  If omitted, the first listener specified in 'resources' is used.
- `namespace` (String) Namespace is the Kubernetes service namespace. In CiliumEnvoyConfig namespace this is overridden to the namespace of the CEC, In CiliumClusterwideEnvoyConfig namespace defaults to 'default'.

