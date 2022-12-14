---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_gateway_networking_k8s_io_tls_route_v1alpha2 Resource - terraform-provider-k8s"
subcategory: "gateway.networking.k8s.io"
description: |-
  The TLSRoute resource is similar to TCPRoute, but can be configured to match against TLS-specific metadata. This allows more flexibility in matching streams for a given TLS listener.  If you need to forward traffic to a single target for a TLS listener, you could choose to use a TCPRoute with a TLS listener.
---

# k8s_gateway_networking_k8s_io_tls_route_v1alpha2 (Resource)

The TLSRoute resource is similar to TCPRoute, but can be configured to match against TLS-specific metadata. This allows more flexibility in matching streams for a given TLS listener.  If you need to forward traffic to a single target for a TLS listener, you could choose to use a TCPRoute with a TLS listener.

## Example Usage

```terraform
resource "k8s_gateway_networking_k8s_io_tls_route_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    rules = []
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of TLSRoute. (see [below for nested schema](#nestedatt--spec))

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

- `rules` (Attributes List) Rules are a list of TLS matchers and actions. (see [below for nested schema](#nestedatt--spec--rules))

Optional:

- `hostnames` (List of String) Hostnames defines a set of SNI names that should match against the SNI attribute of TLS ClientHello message in TLS handshake. This matches the RFC 1123 definition of a hostname with 2 notable exceptions:  1. IPs are not allowed in SNI names per RFC 6066. 2. A hostname may be prefixed with a wildcard label ('*.'). The wildcard    label must appear by itself as the first label.  If a hostname is specified by both the Listener and TLSRoute, there must be at least one intersecting hostname for the TLSRoute to be attached to the Listener. For example:  * A Listener with 'test.example.com' as the hostname matches TLSRoutes   that have either not specified any hostnames, or have specified at   least one of 'test.example.com' or '*.example.com'. * A Listener with '*.example.com' as the hostname matches TLSRoutes   that have either not specified any hostnames or have specified at least   one hostname that matches the Listener hostname. For example,   'test.example.com' and '*.example.com' would both match. On the other   hand, 'example.com' and 'test.example.net' would not match.  If both the Listener and TLSRoute have specified hostnames, any TLSRoute hostnames that do not match the Listener hostname MUST be ignored. For example, if a Listener specified '*.example.com', and the TLSRoute specified 'test.example.com' and 'test.example.net', 'test.example.net' must not be considered for a match.  If both the Listener and TLSRoute have specified hostnames, and none match with the criteria above, then the TLSRoute is not accepted. The implementation must raise an 'Accepted' Condition with a status of 'False' in the corresponding RouteParentStatus.  Support: Core
- `parent_refs` (Attributes List) ParentRefs references the resources (usually Gateways) that a Route wants to be attached to. Note that the referenced parent resource needs to allow this for the attachment to be complete. For Gateways, that means the Gateway needs to allow attachment from Routes of this kind and namespace.  The only kind of parent resource with 'Core' support is Gateway. This API may be extended in the future to support additional kinds of parent resources such as one of the route kinds.  It is invalid to reference an identical parent more than once. It is valid to reference multiple distinct sections within the same parent resource, such as 2 Listeners within a Gateway.  It is possible to separately reference multiple distinct objects that may be collapsed by an implementation. For example, some implementations may choose to merge compatible Gateway Listeners together. If that is the case, the list of routes attached to those resources should also be merged. (see [below for nested schema](#nestedatt--spec--parent_refs))

<a id="nestedatt--spec--rules"></a>
### Nested Schema for `spec.rules`

Optional:

- `backend_refs` (Attributes List) BackendRefs defines the backend(s) where matching requests should be sent. If unspecified or invalid (refers to a non-existent resource or a Service with no endpoints), the rule performs no forwarding; if no filters are specified that would result in a response being sent, the underlying implementation must actively reject request attempts to this backend, by rejecting the connection or returning a 503 status code. Request rejections must respect weight; if an invalid backend is requested to have 80% of requests, then 80% of requests must be rejected instead.  Support: Core for Kubernetes Service Support: Custom for any other resource  Support for weight: Extended (see [below for nested schema](#nestedatt--spec--rules--backend_refs))

<a id="nestedatt--spec--rules--backend_refs"></a>
### Nested Schema for `spec.rules.backend_refs`

Required:

- `name` (String) Name is the name of the referent.

Optional:

- `group` (String) Group is the group of the referent. For example, 'networking.k8s.io'. When unspecified (empty string), core API group is inferred.
- `kind` (String) Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.
- `namespace` (String) Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferencePolicy object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferencePolicy documentation for details.  Support: Core
- `port` (Number) Port specifies the destination port number to use for this resource. Port is required when the referent is a Kubernetes Service. For other resources, destination port might be derived from the referent resource or this field.
- `weight` (Number) Weight specifies the proportion of requests forwarded to the referenced backend. This is computed as weight/(sum of all weights in this BackendRefs list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. Weight is not a percentage and the sum of weights does not need to equal 100.  If only one backend is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weight is set to 0, no traffic should be forwarded for this entry. If unspecified, weight defaults to 1.  Support for this field varies based on the context where used.



<a id="nestedatt--spec--parent_refs"></a>
### Nested Schema for `spec.parent_refs`

Required:

- `name` (String) Name is the name of the referent.  Support: Core

Optional:

- `group` (String) Group is the group of the referent.  Support: Core
- `kind` (String) Kind is kind of the referent.  Support: Core (Gateway) Support: Custom (Other Resources)
- `namespace` (String) Namespace is the namespace of the referent. When unspecified (or empty string), this refers to the local namespace of the Route.  Support: Core
- `section_name` (String) SectionName is the name of a section within the target resource. In the following resources, SectionName is interpreted as the following:  * Gateway: Listener Name  Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.  When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.  Support: Core


