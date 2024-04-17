---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "cloudfront.services.k8s.aws"
description: |-
  ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API
---

# k8s_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest (Data Source)

ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API

## Example Usage

```terraform
data "k8s_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy.A response headers policy.A response headers policy contains information about a set of HTTP responseheaders.After you create a response headers policy, you can use its ID to attachit to one or more cache behaviors in a CloudFront distribution. When it'sattached to a cache behavior, the response headers policy affects the HTTPheaders that CloudFront includes in HTTP responses to requests that matchthe cache behavior. CloudFront adds or removes response headers accordingto the configuration of the response headers policy.For more information, see Adding or removing HTTP headers in CloudFront responses(https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html)in the Amazon CloudFront Developer Guide. (see [below for nested schema](#nestedatt--spec))

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

- `response_headers_policy_config` (Attributes) Contains metadata about the response headers policy, and a set of configurationsthat specify the HTTP headers. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config))

<a id="nestedatt--spec--response_headers_policy_config"></a>
### Nested Schema for `spec.response_headers_policy_config`

Optional:

- `comment` (String)
- `cors_config` (Attributes) A configuration for a set of HTTP response headers that are used for cross-originresource sharing (CORS). CloudFront adds these headers to HTTP responsesthat it sends for CORS requests that match a cache behavior associated withthis response headers policy.For more information about CORS, see Cross-Origin Resource Sharing (CORS)(https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--cors_config))
- `custom_headers_config` (Attributes) A list of HTTP response header names and their values. CloudFront includesthese headers in HTTP responses that it sends for requests that match a cachebehavior that's associated with this response headers policy. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--custom_headers_config))
- `name` (String)
- `remove_headers_config` (Attributes) A list of HTTP header names that CloudFront removes from HTTP responses torequests that match the cache behavior that this response headers policyis attached to. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--remove_headers_config))
- `security_headers_config` (Attributes) A configuration for a set of security-related HTTP response headers. CloudFrontadds these headers to HTTP responses that it sends for requests that matcha cache behavior associated with this response headers policy. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config))
- `server_timing_headers_config` (Attributes) A configuration for enabling the Server-Timing header in HTTP responses sentfrom CloudFront. CloudFront adds this header to HTTP responses that it sendsin response to requests that match a cache behavior that's associated withthis response headers policy.You can use the Server-Timing header to view metrics that can help you gaininsights about the behavior and performance of CloudFront. For example, youcan see which cache layer served a cache hit, or the first byte latency fromthe origin when there was a cache miss. You can use the metrics in the Server-Timingheader to troubleshoot issues or test the efficiency of your CloudFront configuration.For more information, see Server-Timing header (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#server-timing-header)in the Amazon CloudFront Developer Guide. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--server_timing_headers_config))

<a id="nestedatt--spec--response_headers_policy_config--cors_config"></a>
### Nested Schema for `spec.response_headers_policy_config.cors_config`

Optional:

- `access_control_allow_credentials` (Boolean)
- `access_control_allow_headers` (Attributes) A list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-HeadersHTTP response header.For more information about the Access-Control-Allow-Headers HTTP responseheader, see Access-Control-Allow-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_headers))
- `access_control_allow_methods` (Attributes) A list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-MethodsHTTP response header.For more information about the Access-Control-Allow-Methods HTTP responseheader, see Access-Control-Allow-Methods (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_methods))
- `access_control_allow_origins` (Attributes) A list of origins (domain names) that CloudFront can use as the value forthe Access-Control-Allow-Origin HTTP response header.For more information about the Access-Control-Allow-Origin HTTP responseheader, see Access-Control-Allow-Origin (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_origins))
- `access_control_expose_headers` (Attributes) A list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-HeadersHTTP response header.For more information about the Access-Control-Expose-Headers HTTP responseheader, see Access-Control-Expose-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--cors_config--access_control_expose_headers))
- `access_control_max_age_sec` (Number)
- `origin_override` (Boolean)

<a id="nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_headers"></a>
### Nested Schema for `spec.response_headers_policy_config.cors_config.access_control_allow_headers`

Optional:

- `items` (List of String)


<a id="nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_methods"></a>
### Nested Schema for `spec.response_headers_policy_config.cors_config.access_control_allow_methods`

Optional:

- `items` (List of String)


<a id="nestedatt--spec--response_headers_policy_config--cors_config--access_control_allow_origins"></a>
### Nested Schema for `spec.response_headers_policy_config.cors_config.access_control_allow_origins`

Optional:

- `items` (List of String)


<a id="nestedatt--spec--response_headers_policy_config--cors_config--access_control_expose_headers"></a>
### Nested Schema for `spec.response_headers_policy_config.cors_config.access_control_expose_headers`

Optional:

- `items` (List of String)



<a id="nestedatt--spec--response_headers_policy_config--custom_headers_config"></a>
### Nested Schema for `spec.response_headers_policy_config.custom_headers_config`

Optional:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--custom_headers_config--items))

<a id="nestedatt--spec--response_headers_policy_config--custom_headers_config--items"></a>
### Nested Schema for `spec.response_headers_policy_config.custom_headers_config.items`

Optional:

- `header` (String)
- `override` (Boolean)
- `value` (String)



<a id="nestedatt--spec--response_headers_policy_config--remove_headers_config"></a>
### Nested Schema for `spec.response_headers_policy_config.remove_headers_config`

Optional:

- `items` (Attributes List) (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--remove_headers_config--items))

<a id="nestedatt--spec--response_headers_policy_config--remove_headers_config--items"></a>
### Nested Schema for `spec.response_headers_policy_config.remove_headers_config.items`

Optional:

- `header` (String)



<a id="nestedatt--spec--response_headers_policy_config--security_headers_config"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config`

Optional:

- `content_security_policy` (Attributes) The policy directives and their values that CloudFront includes as valuesfor the Content-Security-Policy HTTP response header.For more information about the Content-Security-Policy HTTP response header,see Content-Security-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--content_security_policy))
- `content_type_options` (Attributes) Determines whether CloudFront includes the X-Content-Type-Options HTTP responseheader with its value set to nosniff.For more information about the X-Content-Type-Options HTTP response header,see X-Content-Type-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--content_type_options))
- `frame_options` (Attributes) Determines whether CloudFront includes the X-Frame-Options HTTP responseheader and the header's value.For more information about the X-Frame-Options HTTP response header, seeX-Frame-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--frame_options))
- `referrer_policy` (Attributes) Determines whether CloudFront includes the Referrer-Policy HTTP responseheader and the header's value.For more information about the Referrer-Policy HTTP response header, seeReferrer-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--referrer_policy))
- `strict_transport_security` (Attributes) Determines whether CloudFront includes the Strict-Transport-Security HTTPresponse header and the header's value.For more information about the Strict-Transport-Security HTTP response header,see Strict-Transport-Security (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--strict_transport_security))
- `x_ss_protection` (Attributes) Determines whether CloudFront includes the X-XSS-Protection HTTP responseheader and the header's value.For more information about the X-XSS-Protection HTTP response header, seeX-XSS-Protection (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection)in the MDN Web Docs. (see [below for nested schema](#nestedatt--spec--response_headers_policy_config--security_headers_config--x_ss_protection))

<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--content_security_policy"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.content_security_policy`

Optional:

- `content_security_policy` (String)
- `override` (Boolean)


<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--content_type_options"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.content_type_options`

Optional:

- `override` (Boolean)


<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--frame_options"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.frame_options`

Optional:

- `frame_option` (String)
- `override` (Boolean)


<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--referrer_policy"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.referrer_policy`

Optional:

- `override` (Boolean)
- `referrer_policy` (String)


<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--strict_transport_security"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.strict_transport_security`

Optional:

- `access_control_max_age_sec` (Number)
- `include_subdomains` (Boolean)
- `override` (Boolean)
- `preload` (Boolean)


<a id="nestedatt--spec--response_headers_policy_config--security_headers_config--x_ss_protection"></a>
### Nested Schema for `spec.response_headers_policy_config.security_headers_config.x_ss_protection`

Optional:

- `mode_block` (Boolean)
- `override` (Boolean)
- `protection` (Boolean)
- `report_uri` (String)



<a id="nestedatt--spec--response_headers_policy_config--server_timing_headers_config"></a>
### Nested Schema for `spec.response_headers_policy_config.server_timing_headers_config`

Optional:

- `enabled` (Boolean)
- `sampling_rate` (Number)