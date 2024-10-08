---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "kuadrant.io"
description: |-
  RateLimitPolicy enables rate limiting for service workloads in a Gateway API network
---

# k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest (Data Source)

RateLimitPolicy enables rate limiting for service workloads in a Gateway API network

## Example Usage

```terraform
data "k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest" "example" {
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

- `spec` (Attributes) RateLimitPolicySpec defines the desired state of RateLimitPolicy (see [below for nested schema](#nestedatt--spec))

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

- `target_ref` (Attributes) TargetRef identifies an API object to apply policy to. (see [below for nested schema](#nestedatt--spec--target_ref))

Optional:

- `defaults` (Attributes) Defaults define explicit default values for this policy and for policies inheriting this policy. Defaults are mutually exclusive with implicit defaults defined by RateLimitPolicyCommonSpec. (see [below for nested schema](#nestedatt--spec--defaults))
- `limits` (Attributes) Limits holds the struct of limits indexed by a unique name (see [below for nested schema](#nestedatt--spec--limits))
- `overrides` (Attributes) Overrides define override values for this policy and for policies inheriting this policy. Overrides are mutually exclusive with implicit defaults and explicit Defaults defined by RateLimitPolicyCommonSpec. (see [below for nested schema](#nestedatt--spec--overrides))

<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Required:

- `group` (String) Group is the group of the target resource.
- `kind` (String) Kind is kind of the target resource.
- `name` (String) Name is the name of the target resource.


<a id="nestedatt--spec--defaults"></a>
### Nested Schema for `spec.defaults`

Optional:

- `limits` (Attributes) Limits holds the struct of limits indexed by a unique name (see [below for nested schema](#nestedatt--spec--defaults--limits))

<a id="nestedatt--spec--defaults--limits"></a>
### Nested Schema for `spec.defaults.limits`

Optional:

- `counters` (List of String) Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `rates` (Attributes List) Rates holds the list of limit rates (see [below for nested schema](#nestedatt--spec--defaults--limits--rates))
- `route_selectors` (Attributes List) RouteSelectors defines semantics for matching an HTTP request based on conditions (see [below for nested schema](#nestedatt--spec--defaults--limits--route_selectors))
- `when` (Attributes List) When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match (see [below for nested schema](#nestedatt--spec--defaults--limits--when))

<a id="nestedatt--spec--defaults--limits--rates"></a>
### Nested Schema for `spec.defaults.limits.rates`

Required:

- `duration` (Number) Duration defines the time period for which the Limit specified above applies.
- `limit` (Number) Limit defines the max value allowed for a given period of time
- `unit` (String) Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'


<a id="nestedatt--spec--defaults--limits--route_selectors"></a>
### Nested Schema for `spec.defaults.limits.route_selectors`

Optional:

- `hostnames` (List of String) Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec
- `matches` (Attributes List) Matches define conditions used for matching the rule against incoming HTTP requests. https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec (see [below for nested schema](#nestedatt--spec--defaults--limits--route_selectors--matches))

<a id="nestedatt--spec--defaults--limits--route_selectors--matches"></a>
### Nested Schema for `spec.defaults.limits.route_selectors.matches`

Optional:

- `headers` (Attributes List) Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route. (see [below for nested schema](#nestedatt--spec--defaults--limits--route_selectors--matches--headers))
- `method` (String) Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method. Support: Extended
- `path` (Attributes) Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided. (see [below for nested schema](#nestedatt--spec--defaults--limits--route_selectors--matches--path))
- `query_params` (Attributes List) QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route. Support: Extended (see [below for nested schema](#nestedatt--spec--defaults--limits--route_selectors--matches--query_params))

<a id="nestedatt--spec--defaults--limits--route_selectors--matches--headers"></a>
### Nested Schema for `spec.defaults.limits.route_selectors.matches.headers`

Required:

- `name` (String) Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2). If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent. When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.
- `value` (String) Value is the value of HTTP Header to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the header. Support: Core (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression HeaderMatchType has implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.


<a id="nestedatt--spec--defaults--limits--route_selectors--matches--path"></a>
### Nested Schema for `spec.defaults.limits.route_selectors.matches.path`

Optional:

- `type` (String) Type specifies how to match against the path Value. Support: Core (Exact, PathPrefix) Support: Implementation-specific (RegularExpression)
- `value` (String) Value of the HTTP path to match against.


<a id="nestedatt--spec--defaults--limits--route_selectors--matches--query_params"></a>
### Nested Schema for `spec.defaults.limits.route_selectors.matches.query_params`

Required:

- `name` (String) Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3). If multiple entries specify equivalent query param names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent query param name MUST be ignored. If a query param is repeated in an HTTP request, the behavior is purposely left undefined, since different data planes have different capabilities. However, it is *recommended* that implementations should match against the first value of the param if the data plane supports it, as this behavior is expected in other load balancing contexts outside of the Gateway API. Users SHOULD NOT route traffic based on repeated query params to guard themselves against potential differences in the implementations.
- `value` (String) Value is the value of HTTP query param to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the query parameter. Support: Extended (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression QueryParamMatchType has Implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.




<a id="nestedatt--spec--defaults--limits--when"></a>
### Nested Schema for `spec.defaults.limits.when`

Required:

- `operator` (String) The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)
- `selector` (String) Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `value` (String) The value of reference for the comparison.




<a id="nestedatt--spec--limits"></a>
### Nested Schema for `spec.limits`

Optional:

- `counters` (List of String) Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `rates` (Attributes List) Rates holds the list of limit rates (see [below for nested schema](#nestedatt--spec--limits--rates))
- `route_selectors` (Attributes List) RouteSelectors defines semantics for matching an HTTP request based on conditions (see [below for nested schema](#nestedatt--spec--limits--route_selectors))
- `when` (Attributes List) When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match (see [below for nested schema](#nestedatt--spec--limits--when))

<a id="nestedatt--spec--limits--rates"></a>
### Nested Schema for `spec.limits.rates`

Required:

- `duration` (Number) Duration defines the time period for which the Limit specified above applies.
- `limit` (Number) Limit defines the max value allowed for a given period of time
- `unit` (String) Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'


<a id="nestedatt--spec--limits--route_selectors"></a>
### Nested Schema for `spec.limits.route_selectors`

Optional:

- `hostnames` (List of String) Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec
- `matches` (Attributes List) Matches define conditions used for matching the rule against incoming HTTP requests. https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec (see [below for nested schema](#nestedatt--spec--limits--route_selectors--matches))

<a id="nestedatt--spec--limits--route_selectors--matches"></a>
### Nested Schema for `spec.limits.route_selectors.matches`

Optional:

- `headers` (Attributes List) Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route. (see [below for nested schema](#nestedatt--spec--limits--route_selectors--matches--headers))
- `method` (String) Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method. Support: Extended
- `path` (Attributes) Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided. (see [below for nested schema](#nestedatt--spec--limits--route_selectors--matches--path))
- `query_params` (Attributes List) QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route. Support: Extended (see [below for nested schema](#nestedatt--spec--limits--route_selectors--matches--query_params))

<a id="nestedatt--spec--limits--route_selectors--matches--headers"></a>
### Nested Schema for `spec.limits.route_selectors.matches.headers`

Required:

- `name` (String) Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2). If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent. When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.
- `value` (String) Value is the value of HTTP Header to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the header. Support: Core (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression HeaderMatchType has implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.


<a id="nestedatt--spec--limits--route_selectors--matches--path"></a>
### Nested Schema for `spec.limits.route_selectors.matches.path`

Optional:

- `type` (String) Type specifies how to match against the path Value. Support: Core (Exact, PathPrefix) Support: Implementation-specific (RegularExpression)
- `value` (String) Value of the HTTP path to match against.


<a id="nestedatt--spec--limits--route_selectors--matches--query_params"></a>
### Nested Schema for `spec.limits.route_selectors.matches.query_params`

Required:

- `name` (String) Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3). If multiple entries specify equivalent query param names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent query param name MUST be ignored. If a query param is repeated in an HTTP request, the behavior is purposely left undefined, since different data planes have different capabilities. However, it is *recommended* that implementations should match against the first value of the param if the data plane supports it, as this behavior is expected in other load balancing contexts outside of the Gateway API. Users SHOULD NOT route traffic based on repeated query params to guard themselves against potential differences in the implementations.
- `value` (String) Value is the value of HTTP query param to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the query parameter. Support: Extended (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression QueryParamMatchType has Implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.




<a id="nestedatt--spec--limits--when"></a>
### Nested Schema for `spec.limits.when`

Required:

- `operator` (String) The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)
- `selector` (String) Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `value` (String) The value of reference for the comparison.



<a id="nestedatt--spec--overrides"></a>
### Nested Schema for `spec.overrides`

Optional:

- `limits` (Attributes) Limits holds the struct of limits indexed by a unique name (see [below for nested schema](#nestedatt--spec--overrides--limits))

<a id="nestedatt--spec--overrides--limits"></a>
### Nested Schema for `spec.overrides.limits`

Optional:

- `counters` (List of String) Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `rates` (Attributes List) Rates holds the list of limit rates (see [below for nested schema](#nestedatt--spec--overrides--limits--rates))
- `route_selectors` (Attributes List) RouteSelectors defines semantics for matching an HTTP request based on conditions (see [below for nested schema](#nestedatt--spec--overrides--limits--route_selectors))
- `when` (Attributes List) When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match (see [below for nested schema](#nestedatt--spec--overrides--limits--when))

<a id="nestedatt--spec--overrides--limits--rates"></a>
### Nested Schema for `spec.overrides.limits.rates`

Required:

- `duration` (Number) Duration defines the time period for which the Limit specified above applies.
- `limit` (Number) Limit defines the max value allowed for a given period of time
- `unit` (String) Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'


<a id="nestedatt--spec--overrides--limits--route_selectors"></a>
### Nested Schema for `spec.overrides.limits.route_selectors`

Optional:

- `hostnames` (List of String) Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the request https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec
- `matches` (Attributes List) Matches define conditions used for matching the rule against incoming HTTP requests. https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec (see [below for nested schema](#nestedatt--spec--overrides--limits--route_selectors--matches))

<a id="nestedatt--spec--overrides--limits--route_selectors--matches"></a>
### Nested Schema for `spec.overrides.limits.route_selectors.matches`

Optional:

- `headers` (Attributes List) Headers specifies HTTP request header matchers. Multiple match values are ANDed together, meaning, a request must match all the specified headers to select the route. (see [below for nested schema](#nestedatt--spec--overrides--limits--route_selectors--matches--headers))
- `method` (String) Method specifies HTTP method matcher. When specified, this route will be matched only if the request has the specified method. Support: Extended
- `path` (Attributes) Path specifies a HTTP request path matcher. If this field is not specified, a default prefix match on the '/' path is provided. (see [below for nested schema](#nestedatt--spec--overrides--limits--route_selectors--matches--path))
- `query_params` (Attributes List) QueryParams specifies HTTP query parameter matchers. Multiple match values are ANDed together, meaning, a request must match all the specified query parameters to select the route. Support: Extended (see [below for nested schema](#nestedatt--spec--overrides--limits--route_selectors--matches--query_params))

<a id="nestedatt--spec--overrides--limits--route_selectors--matches--headers"></a>
### Nested Schema for `spec.overrides.limits.route_selectors.matches.headers`

Required:

- `name` (String) Name is the name of the HTTP Header to be matched. Name matching MUST be case insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2). If multiple entries specify equivalent header names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent header name MUST be ignored. Due to the case-insensitivity of header names, 'foo' and 'Foo' are considered equivalent. When a header is repeated in an HTTP request, it is implementation-specific behavior as to how this is represented. Generally, proxies should follow the guidance from the RFC: https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regarding processing a repeated header, with special handling for 'Set-Cookie'.
- `value` (String) Value is the value of HTTP Header to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the header. Support: Core (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression HeaderMatchType has implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.


<a id="nestedatt--spec--overrides--limits--route_selectors--matches--path"></a>
### Nested Schema for `spec.overrides.limits.route_selectors.matches.path`

Optional:

- `type` (String) Type specifies how to match against the path Value. Support: Core (Exact, PathPrefix) Support: Implementation-specific (RegularExpression)
- `value` (String) Value of the HTTP path to match against.


<a id="nestedatt--spec--overrides--limits--route_selectors--matches--query_params"></a>
### Nested Schema for `spec.overrides.limits.route_selectors.matches.query_params`

Required:

- `name` (String) Name is the name of the HTTP query param to be matched. This must be an exact string match. (See https://tools.ietf.org/html/rfc7230#section-2.7.3). If multiple entries specify equivalent query param names, only the first entry with an equivalent name MUST be considered for a match. Subsequent entries with an equivalent query param name MUST be ignored. If a query param is repeated in an HTTP request, the behavior is purposely left undefined, since different data planes have different capabilities. However, it is *recommended* that implementations should match against the first value of the param if the data plane supports it, as this behavior is expected in other load balancing contexts outside of the Gateway API. Users SHOULD NOT route traffic based on repeated query params to guard themselves against potential differences in the implementations.
- `value` (String) Value is the value of HTTP query param to be matched.

Optional:

- `type` (String) Type specifies how to match against the value of the query parameter. Support: Extended (Exact) Support: Implementation-specific (RegularExpression) Since RegularExpression QueryParamMatchType has Implementation-specific conformance, implementations can support POSIX, PCRE or any other dialects of regular expressions. Please read the implementation's documentation to determine the supported dialect.




<a id="nestedatt--spec--overrides--limits--when"></a>
### Nested Schema for `spec.overrides.limits.when`

Required:

- `operator` (String) The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)
- `selector` (String) Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors
- `value` (String) The value of reference for the comparison.
