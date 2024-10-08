---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_elbv2_k8s_aws_ingress_class_params_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "elbv2.k8s.aws"
description: |-
  IngressClassParams is the Schema for the IngressClassParams API
---

# k8s_elbv2_k8s_aws_ingress_class_params_v1beta1_manifest (Data Source)

IngressClassParams is the Schema for the IngressClassParams API

## Example Usage

```terraform
data "k8s_elbv2_k8s_aws_ingress_class_params_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) IngressClassParamsSpec defines the desired state of IngressClassParams (see [below for nested schema](#nestedatt--spec))

### Read-Only

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

- `certificate_arn` (List of String) CertificateArn specifies the ARN of the certificates for all Ingresses that belong to IngressClass with this IngressClassParams.
- `group` (Attributes) Group defines the IngressGroup for all Ingresses that belong to IngressClass with this IngressClassParams. (see [below for nested schema](#nestedatt--spec--group))
- `inbound_cidrs` (List of String) InboundCIDRs specifies the CIDRs that are allowed to access the Ingresses that belong to IngressClass with this IngressClassParams.
- `ip_address_type` (String) IPAddressType defines the ip address type for all Ingresses that belong to IngressClass with this IngressClassParams.
- `listeners` (Attributes List) Listeners define a list of listeners with their protocol, port and attributes. (see [below for nested schema](#nestedatt--spec--listeners))
- `load_balancer_attributes` (Attributes List) LoadBalancerAttributes define the custom attributes to LoadBalancers for all Ingress that that belong to IngressClass with this IngressClassParams. (see [below for nested schema](#nestedatt--spec--load_balancer_attributes))
- `namespace_selector` (Attributes) NamespaceSelector restrict the namespaces of Ingresses that are allowed to specify the IngressClass with this IngressClassParams. * if absent or present but empty, it selects all namespaces. (see [below for nested schema](#nestedatt--spec--namespace_selector))
- `scheme` (String) Scheme defines the scheme for all Ingresses that belong to IngressClass with this IngressClassParams.
- `ssl_policy` (String) SSLPolicy specifies the SSL Policy for all Ingresses that belong to IngressClass with this IngressClassParams.
- `subnets` (Attributes) Subnets defines the subnets for all Ingresses that belong to IngressClass with this IngressClassParams. (see [below for nested schema](#nestedatt--spec--subnets))
- `tags` (Attributes List) Tags defines list of Tags on AWS resources provisioned for Ingresses that belong to IngressClass with this IngressClassParams. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--group"></a>
### Nested Schema for `spec.group`

Required:

- `name` (String) Name is the name of IngressGroup.


<a id="nestedatt--spec--listeners"></a>
### Nested Schema for `spec.listeners`

Optional:

- `listener_attributes` (Attributes List) The attributes of the listener (see [below for nested schema](#nestedatt--spec--listeners--listener_attributes))
- `port` (Number) The port of the listener
- `protocol` (String) The protocol of the listener

<a id="nestedatt--spec--listeners--listener_attributes"></a>
### Nested Schema for `spec.listeners.listener_attributes`

Required:

- `key` (String) The key of the attribute.
- `value` (String) The value of the attribute.



<a id="nestedatt--spec--load_balancer_attributes"></a>
### Nested Schema for `spec.load_balancer_attributes`

Required:

- `key` (String) The key of the attribute.
- `value` (String) The value of the attribute.


<a id="nestedatt--spec--namespace_selector"></a>
### Nested Schema for `spec.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--subnets"></a>
### Nested Schema for `spec.subnets`

Optional:

- `ids` (List of String) IDs specify the resource IDs of subnets. Exactly one of this or 'tags' must be specified.
- `tags` (Map of List of String) Tags specifies subnets in the load balancer's VPC where each tag specified in the map key contains one of the values in the corresponding value list. Exactly one of this or 'ids' must be specified.


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Required:

- `key` (String) The key of the tag.
- `value` (String) The value of the tag.
