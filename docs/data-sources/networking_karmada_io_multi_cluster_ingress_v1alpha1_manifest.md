---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "networking.karmada.io"
description: |-
  MultiClusterIngress is a collection of rules that allow inbound connections to reach theendpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress,indicates the Ingress in multi-clusters.
---

# k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest (Data Source)

MultiClusterIngress is a collection of rules that allow inbound connections to reach theendpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress,indicates the Ingress in multi-clusters.

## Example Usage

```terraform
data "k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) Spec is the desired state of the MultiClusterIngress. (see [below for nested schema](#nestedatt--spec))

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

- `default_backend` (Attributes) defaultBackend is the backend that should handle requests that don'tmatch any rule. If Rules are not specified, DefaultBackend must be specified.If DefaultBackend is not set, the handling of requests that do not match anyof the rules will be up to the Ingress controller. (see [below for nested schema](#nestedatt--spec--default_backend))
- `ingress_class_name` (String) ingressClassName is the name of an IngressClass cluster resource. Ingresscontroller implementations use this field to know whether they should beserving this Ingress resource, by a transitive connection(controller -> IngressClass -> Ingress resource). Although the'kubernetes.io/ingress.class' annotation (simple constant name) was neverformally defined, it was widely supported by Ingress controllers to createa direct binding between Ingress controller and Ingress resources. Newlycreated Ingress resources should prefer using the field. However, eventhough the annotation is officially deprecated, for backwards compatibilityreasons, ingress controllers should still honor that annotation if present.
- `rules` (Attributes List) rules is a list of host rules used to configure the Ingress. If unspecified,or no rule matches, all traffic is sent to the default backend. (see [below for nested schema](#nestedatt--spec--rules))
- `tls` (Attributes List) tls represents the TLS configuration. Currently the Ingress only supports asingle TLS port, 443. If multiple members of this list specify different hosts,they will be multiplexed on the same port according to the hostname specifiedthrough the SNI TLS extension, if the ingress controller fulfilling theingress supports SNI. (see [below for nested schema](#nestedatt--spec--tls))

<a id="nestedatt--spec--default_backend"></a>
### Nested Schema for `spec.default_backend`

Optional:

- `resource` (Attributes) resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'. (see [below for nested schema](#nestedatt--spec--default_backend--resource))
- `service` (Attributes) service references a service as a backend.This is a mutually exclusive setting with 'Resource'. (see [below for nested schema](#nestedatt--spec--default_backend--service))

<a id="nestedatt--spec--default_backend--resource"></a>
### Nested Schema for `spec.default_backend.resource`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--default_backend--service"></a>
### Nested Schema for `spec.default_backend.service`

Required:

- `name` (String) name is the referenced service. The service must exist inthe same namespace as the Ingress object.

Optional:

- `port` (Attributes) port of the referenced service. A port name or port numberis required for a IngressServiceBackend. (see [below for nested schema](#nestedatt--spec--default_backend--service--port))

<a id="nestedatt--spec--default_backend--service--port"></a>
### Nested Schema for `spec.default_backend.service.port`

Optional:

- `name` (String) name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.
- `number` (Number) number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.




<a id="nestedatt--spec--rules"></a>
### Nested Schema for `spec.rules`

Optional:

- `host` (String) host is the fully qualified domain name of a network host, as defined by RFC 3986.Note the following deviations from the 'host' part of theURI as defined in RFC 3986:1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future.Incoming requests are matched against the host before theIngressRuleValue. If the host is unspecified, the Ingress routes alltraffic based on the specified IngressRuleValue.host can be 'precise' which is a domain name without the terminating dot ofa network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain nameprefixed with a single wildcard label (e.g. '*.foo.com').The wildcard character '*' must appear by itself as the first DNS label andmatches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*').Requests will be matched against the Host field in the following way:1. If host is precise, the request matches this rule if the http host header is equal to Host.2. If host is a wildcard, then the request matches this rule if the http host headeris to equal to the suffix (removing the first label) of the wildcard rule.
- `http` (Attributes) HTTPIngressRuleValue is a list of http selectors pointing to backends.In the example: http://<host>/<path>?<searchpart> -> backend wherewhere parts of the url correspond to RFC 3986, this resource will be usedto match against everything after the last '/' and before the first '?'or '#'. (see [below for nested schema](#nestedatt--spec--rules--http))

<a id="nestedatt--spec--rules--http"></a>
### Nested Schema for `spec.rules.http`

Required:

- `paths` (Attributes List) paths is a collection of paths that map requests to backends. (see [below for nested schema](#nestedatt--spec--rules--http--paths))

<a id="nestedatt--spec--rules--http--paths"></a>
### Nested Schema for `spec.rules.http.paths`

Required:

- `backend` (Attributes) backend defines the referenced service endpoint to which the trafficwill be forwarded to. (see [below for nested schema](#nestedatt--spec--rules--http--paths--backend))
- `path_type` (String) pathType determines the interpretation of the path matching. PathType canbe one of the following values:* Exact: Matches the URL path exactly.* Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.

Optional:

- `path` (String) path is matched against the path of an incoming request. Currently it cancontain characters disallowed from the conventional 'path' part of a URLas defined by RFC 3986. Paths must begin with a '/' and must be presentwhen using PathType with value 'Exact' or 'Prefix'.

<a id="nestedatt--spec--rules--http--paths--backend"></a>
### Nested Schema for `spec.rules.http.paths.backend`

Optional:

- `resource` (Attributes) resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'. (see [below for nested schema](#nestedatt--spec--rules--http--paths--backend--resource))
- `service` (Attributes) service references a service as a backend.This is a mutually exclusive setting with 'Resource'. (see [below for nested schema](#nestedatt--spec--rules--http--paths--backend--service))

<a id="nestedatt--spec--rules--http--paths--backend--resource"></a>
### Nested Schema for `spec.rules.http.paths.backend.resource`

Required:

- `kind` (String) Kind is the type of resource being referenced
- `name` (String) Name is the name of resource being referenced

Optional:

- `api_group` (String) APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.


<a id="nestedatt--spec--rules--http--paths--backend--service"></a>
### Nested Schema for `spec.rules.http.paths.backend.service`

Required:

- `name` (String) name is the referenced service. The service must exist inthe same namespace as the Ingress object.

Optional:

- `port` (Attributes) port of the referenced service. A port name or port numberis required for a IngressServiceBackend. (see [below for nested schema](#nestedatt--spec--rules--http--paths--backend--service--port))

<a id="nestedatt--spec--rules--http--paths--backend--service--port"></a>
### Nested Schema for `spec.rules.http.paths.backend.service.port`

Optional:

- `name` (String) name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.
- `number` (Number) number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.







<a id="nestedatt--spec--tls"></a>
### Nested Schema for `spec.tls`

Optional:

- `hosts` (List of String) hosts is a list of hosts included in the TLS certificate. The values inthis list must match the name/s used in the tlsSecret. Defaults to thewildcard host setting for the loadbalancer controller fulfilling thisIngress, if left unspecified.
- `secret_name` (String) secretName is the name of the secret used to terminate TLS traffic onport 443. Field is left optional to allow TLS routing based on SNIhostname alone. If the SNI host in a listener conflicts with the 'Host'header field used by an IngressRule, the SNI host is used for terminationand value of the 'Host' header is used for routing.
