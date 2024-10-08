---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apigatewayv2_services_k8s_aws_route_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "apigatewayv2.services.k8s.aws"
description: |-
  Route is the Schema for the Routes API
---

# k8s_apigatewayv2_services_k8s_aws_route_v1alpha1_manifest (Data Source)

Route is the Schema for the Routes API

## Example Usage

```terraform
data "k8s_apigatewayv2_services_k8s_aws_route_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) RouteSpec defines the desired state of Route. Represents a route. (see [below for nested schema](#nestedatt--spec))

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

- `route_key` (String)

Optional:

- `api_id` (String)
- `api_key_required` (Boolean)
- `api_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api (see [below for nested schema](#nestedatt--spec--api_ref))
- `authorization_scopes` (List of String)
- `authorization_type` (String)
- `authorizer_id` (String)
- `authorizer_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api (see [below for nested schema](#nestedatt--spec--authorizer_ref))
- `model_selection_expression` (String)
- `operation_name` (String)
- `request_models` (Map of String)
- `request_parameters` (Attributes) (see [below for nested schema](#nestedatt--spec--request_parameters))
- `route_response_selection_expression` (String)
- `target` (String)
- `target_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api (see [below for nested schema](#nestedatt--spec--target_ref))

<a id="nestedatt--spec--api_ref"></a>
### Nested Schema for `spec.api_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--api_ref--from))

<a id="nestedatt--spec--api_ref--from"></a>
### Nested Schema for `spec.api_ref.from`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--authorizer_ref"></a>
### Nested Schema for `spec.authorizer_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--authorizer_ref--from))

<a id="nestedatt--spec--authorizer_ref--from"></a>
### Nested Schema for `spec.authorizer_ref.from`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--request_parameters"></a>
### Nested Schema for `spec.request_parameters`

Optional:

- `required` (Boolean)


<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--target_ref--from))

<a id="nestedatt--spec--target_ref--from"></a>
### Nested Schema for `spec.target_ref.from`

Optional:

- `name` (String)
- `namespace` (String)
