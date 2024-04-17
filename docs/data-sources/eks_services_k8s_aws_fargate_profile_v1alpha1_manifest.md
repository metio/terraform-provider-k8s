---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_eks_services_k8s_aws_fargate_profile_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "eks.services.k8s.aws"
description: |-
  FargateProfile is the Schema for the FargateProfiles API
---

# k8s_eks_services_k8s_aws_fargate_profile_v1alpha1_manifest (Data Source)

FargateProfile is the Schema for the FargateProfiles API

## Example Usage

```terraform
data "k8s_eks_services_k8s_aws_fargate_profile_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) FargateProfileSpec defines the desired state of FargateProfile.An object representing an Fargate profile. (see [below for nested schema](#nestedatt--spec))

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

- `name` (String) The name of the Fargate profile.

Optional:

- `client_request_token` (String) A unique, case-sensitive identifier that you provide to ensure the idempotencyof the request.
- `cluster_name` (String) The name of your cluster.
- `cluster_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--cluster_ref))
- `pod_execution_role_arn` (String) The Amazon Resource Name (ARN) of the Pod execution role to use for a Podthat matches the selectors in the Fargate profile. The Pod execution roleallows Fargate infrastructure to register with your cluster as a node, andit provides read access to Amazon ECR image repositories. For more information,see Pod execution role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html)in the Amazon EKS User Guide.
- `pod_execution_role_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--pod_execution_role_ref))
- `selectors` (Attributes List) The selectors to match for a Pod to use this Fargate profile. Each selectormust have an associated Kubernetes namespace. Optionally, you can also specifylabels for a namespace. You may specify up to five selectors in a Fargateprofile. (see [below for nested schema](#nestedatt--spec--selectors))
- `subnet_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--subnet_refs))
- `subnets` (List of String) The IDs of subnets to launch a Pod into. A Pod running on Fargate isn't assigneda public IP address, so only private subnets (with no direct route to anInternet Gateway) are accepted for this parameter.
- `tags` (Map of String) Metadata that assists with categorization and organization. Each tag consistsof a key and an optional value. You define both. Tags don't propagate toany other cluster or Amazon Web Services resources.

<a id="nestedatt--spec--cluster_ref"></a>
### Nested Schema for `spec.cluster_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--cluster_ref--from))

<a id="nestedatt--spec--cluster_ref--from"></a>
### Nested Schema for `spec.cluster_ref.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--pod_execution_role_ref"></a>
### Nested Schema for `spec.pod_execution_role_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--pod_execution_role_ref--from))

<a id="nestedatt--spec--pod_execution_role_ref--from"></a>
### Nested Schema for `spec.pod_execution_role_ref.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--selectors"></a>
### Nested Schema for `spec.selectors`

Optional:

- `labels` (Map of String)
- `namespace` (String)


<a id="nestedatt--spec--subnet_refs"></a>
### Nested Schema for `spec.subnet_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--subnet_refs--from))

<a id="nestedatt--spec--subnet_refs--from"></a>
### Nested Schema for `spec.subnet_refs.from`

Optional:

- `name` (String)