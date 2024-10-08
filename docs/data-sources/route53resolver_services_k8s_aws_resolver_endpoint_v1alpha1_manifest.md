---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "route53resolver.services.k8s.aws"
description: |-
  ResolverEndpoint is the Schema for the ResolverEndpoints API
---

# k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest (Data Source)

ResolverEndpoint is the Schema for the ResolverEndpoints API

## Example Usage

```terraform
data "k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ResolverEndpointSpec defines the desired state of ResolverEndpoint. In the response to a CreateResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html), DeleteResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverEndpoint.html), GetResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html), Updates the name, or ResolverEndpointType for an endpoint, or UpdateResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverEndpoint.html) request, a complex type that contains settings for an existing inbound or outbound Resolver endpoint. (see [below for nested schema](#nestedatt--spec))

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

- `direction` (String) Specify the applicable value: * INBOUND: Resolver forwards DNS queries to the DNS service for a VPC from your network * OUTBOUND: Resolver forwards DNS queries from the DNS service for a VPC to your network
- `ip_addresses` (Attributes List) The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints). The subnet ID uniquely identifies a VPC. (see [below for nested schema](#nestedatt--spec--ip_addresses))

Optional:

- `name` (String) A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
- `resolver_endpoint_type` (String) For the endpoint type you can choose either IPv4, IPv6. or dual-stack. A dual-stack endpoint means that it will resolve via both IPv4 and IPv6. This endpoint type is applied to all IP addresses.
- `security_group_i_ds` (List of String) The ID of one or more security groups that you want to use to control access to this VPC. The security group that you specify must include one or more inbound rules (for inbound Resolver endpoints) or outbound rules (for outbound Resolver endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
- `security_group_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--security_group_refs))
- `tags` (Attributes List) A list of the tag keys and values that you want to associate with the endpoint. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--ip_addresses"></a>
### Nested Schema for `spec.ip_addresses`

Optional:

- `ip` (String)
- `ipv6` (String)
- `subnet_id` (String)
- `subnet_ref` (Attributes) Reference field for SubnetID (see [below for nested schema](#nestedatt--spec--ip_addresses--subnet_ref))

<a id="nestedatt--spec--ip_addresses--subnet_ref"></a>
### Nested Schema for `spec.ip_addresses.subnet_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--ip_addresses--subnet_ref--from))

<a id="nestedatt--spec--ip_addresses--subnet_ref--from"></a>
### Nested Schema for `spec.ip_addresses.subnet_ref.from`

Optional:

- `name` (String)
- `namespace` (String)




<a id="nestedatt--spec--security_group_refs"></a>
### Nested Schema for `spec.security_group_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--security_group_refs--from))

<a id="nestedatt--spec--security_group_refs--from"></a>
### Nested Schema for `spec.security_group_refs.from`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
