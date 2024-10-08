---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "cloudfront.services.k8s.aws"
description: |-
  OriginRequestPolicy is the Schema for the OriginRequestPolicies API
---

# k8s_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest (Data Source)

OriginRequestPolicy is the Schema for the OriginRequestPolicies API

## Example Usage

```terraform
data "k8s_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) OriginRequestPolicySpec defines the desired state of OriginRequestPolicy. An origin request policy. When it's attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following: * The request body and the URL path (without the domain name) from the viewer request. * The headers that CloudFront automatically includes in every origin request, including Host, User-Agent, and X-Amz-Cf-Id. * All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront. CloudFront sends a request when it can't find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use CachePolicy. (see [below for nested schema](#nestedatt--spec))

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

- `origin_request_policy_config` (Attributes) An origin request policy configuration. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config))

<a id="nestedatt--spec--origin_request_policy_config"></a>
### Nested Schema for `spec.origin_request_policy_config`

Optional:

- `comment` (String)
- `cookies_config` (Attributes) An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--cookies_config))
- `headers_config` (Attributes) An object that determines whether any HTTP headers (and if so, which headers) are included in requests that CloudFront sends to the origin. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--headers_config))
- `name` (String)
- `query_strings_config` (Attributes) An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--query_strings_config))

<a id="nestedatt--spec--origin_request_policy_config--cookies_config"></a>
### Nested Schema for `spec.origin_request_policy_config.cookies_config`

Optional:

- `cookie_behavior` (String)
- `cookies` (Attributes) Contains a list of cookie names. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--cookies_config--cookies))

<a id="nestedatt--spec--origin_request_policy_config--cookies_config--cookies"></a>
### Nested Schema for `spec.origin_request_policy_config.cookies_config.cookies`

Optional:

- `items` (List of String)



<a id="nestedatt--spec--origin_request_policy_config--headers_config"></a>
### Nested Schema for `spec.origin_request_policy_config.headers_config`

Optional:

- `header_behavior` (String)
- `headers` (Attributes) Contains a list of HTTP header names. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--headers_config--headers))

<a id="nestedatt--spec--origin_request_policy_config--headers_config--headers"></a>
### Nested Schema for `spec.origin_request_policy_config.headers_config.headers`

Optional:

- `items` (List of String)



<a id="nestedatt--spec--origin_request_policy_config--query_strings_config"></a>
### Nested Schema for `spec.origin_request_policy_config.query_strings_config`

Optional:

- `query_string_behavior` (String)
- `query_strings` (Attributes) Contains a list of query string names. (see [below for nested schema](#nestedatt--spec--origin_request_policy_config--query_strings_config--query_strings))

<a id="nestedatt--spec--origin_request_policy_config--query_strings_config--query_strings"></a>
### Nested Schema for `spec.origin_request_policy_config.query_strings_config.query_strings`

Optional:

- `items` (List of String)
