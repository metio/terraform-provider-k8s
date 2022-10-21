---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_chaos_mesh_org_pod_http_chaos_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "chaos-mesh.org/v1alpha1"
description: |-
  PodHttpChaos is the Schema for the podhttpchaos API
---

# k8s_chaos_mesh_org_pod_http_chaos_v1alpha1 (Resource)

PodHttpChaos is the Schema for the podhttpchaos API

## Example Usage

```terraform
resource "k8s_chaos_mesh_org_pod_http_chaos_v1alpha1" "minimal" {
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

- `spec` (Attributes) PodHttpChaosSpec defines the desired state of PodHttpChaos. (see [below for nested schema](#nestedatt--spec))

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

- `rules` (Attributes List) Rules are a list of injection rule for http request. (see [below for nested schema](#nestedatt--spec--rules))

<a id="nestedatt--spec--rules"></a>
### Nested Schema for `spec.rules`

Required:

- `actions` (Attributes) Actions contains rules to inject target. (see [below for nested schema](#nestedatt--spec--rules--actions))
- `port` (Number) Port represents the target port to be proxy of.
- `selector` (Attributes) Selector contains the rules to select target. (see [below for nested schema](#nestedatt--spec--rules--selector))
- `target` (String) Target is the object to be selected and injected, <Request|Response>.

Optional:

- `source` (String) Source represents the source of current rules

<a id="nestedatt--spec--rules--actions"></a>
### Nested Schema for `spec.rules.actions`

Optional:

- `abort` (Boolean) Abort is a rule to abort a http session.
- `delay` (String) Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.
- `patch` (Attributes) Patch is a rule to patch some contents in target. (see [below for nested schema](#nestedatt--spec--rules--actions--patch))
- `replace` (Attributes) Replace is a rule to replace some contents in target. (see [below for nested schema](#nestedatt--spec--rules--actions--replace))

<a id="nestedatt--spec--rules--actions--patch"></a>
### Nested Schema for `spec.rules.actions.replace`

Optional:

- `body` (Attributes) Body is a rule to patch message body of target. (see [below for nested schema](#nestedatt--spec--rules--actions--replace--body))
- `headers` (List of String) Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.
- `queries` (List of String) Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.

<a id="nestedatt--spec--rules--actions--replace--body"></a>
### Nested Schema for `spec.rules.actions.replace.body`

Required:

- `type` (String) Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.
- `value` (String) Value is the patch contents.



<a id="nestedatt--spec--rules--actions--replace"></a>
### Nested Schema for `spec.rules.actions.replace`

Optional:

- `body` (String) Body is a rule to replace http message body in target.
- `code` (Number) Code is a rule to replace http status code in response.
- `headers` (Map of String) Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.
- `method` (String) Method is a rule to replace http method in request.
- `path` (String) Path is rule to to replace uri path in http request.
- `queries` (Map of String) Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',



<a id="nestedatt--spec--rules--selector"></a>
### Nested Schema for `spec.rules.selector`

Optional:

- `code` (Number) Code is a rule to select target by http status code in response.
- `method` (String) Method is a rule to select target by http method in request.
- `path` (String) Path is a rule to select target by uri path in http request.
- `port` (Number) Port is a rule to select server listening on specific port.
- `request_headers` (Map of String) RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.
- `response_headers` (Map of String) ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.

