---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_application_networking_k8s_aws_access_log_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "application-networking.k8s.aws"
description: |-
  
---

# k8s_application_networking_k8s_aws_access_log_policy_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_application_networking_k8s_aws_access_log_policy_v1alpha1_manifest" "example" {
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
- `spec` (Attributes) AccessLogPolicySpec defines the desired state of AccessLogPolicy. (see [below for nested schema](#nestedatt--spec))

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

- `destination_arn` (String) The Amazon Resource Name (ARN) of the destination that will store access logs. Supported values are S3 Bucket, CloudWatch Log Group, and Firehose Delivery Stream ARNs. Changes to this value results in replacement of the VPC Lattice Access Log Subscription.
- `target_ref` (Attributes) TargetRef points to the Kubernetes Gateway, HTTPRoute, or GRPCRoute resource that will have this policy attached. This field is following the guidelines of Kubernetes Gateway API policy attachment. (see [below for nested schema](#nestedatt--spec--target_ref))

<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Required:

- `group` (String) Group is the group of the target resource.
- `kind` (String) Kind is kind of the target resource.
- `name` (String) Name is the name of the target resource.

Optional:

- `namespace` (String) Namespace is the namespace of the referent. When unspecified, the local namespace is inferred. Even when policy targets a resource in a different namespace, it MUST only apply to traffic originating from the same namespace as the policy.
