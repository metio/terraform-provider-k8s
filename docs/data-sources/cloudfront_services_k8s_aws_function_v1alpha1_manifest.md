---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cloudfront_services_k8s_aws_function_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "cloudfront.services.k8s.aws"
description: |-
  Function is the Schema for the Functions API
---

# k8s_cloudfront_services_k8s_aws_function_v1alpha1_manifest (Data Source)

Function is the Schema for the Functions API

## Example Usage

```terraform
data "k8s_cloudfront_services_k8s_aws_function_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) FunctionSpec defines the desired state of Function. (see [below for nested schema](#nestedatt--spec))

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

- `function_code` (String) The function code. For more information about writing a CloudFront function,see Writing function code for CloudFront Functions (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html)in the Amazon CloudFront Developer Guide.
- `function_config` (Attributes) Configuration information about the function, including an optional commentand the function's runtime. (see [below for nested schema](#nestedatt--spec--function_config))
- `name` (String) A name to identify the function.

<a id="nestedatt--spec--function_config"></a>
### Nested Schema for `spec.function_config`

Optional:

- `comment` (String)
- `runtime` (String)