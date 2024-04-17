---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "ec2.services.k8s.aws"
description: |-
  VPCEndpoint is the Schema for the VPCEndpoints API
---

# k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest (Data Source)

VPCEndpoint is the Schema for the VPCEndpoints API

## Example Usage

```terraform
data "k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) VpcEndpointSpec defines the desired state of VpcEndpoint.Describes a VPC endpoint. (see [below for nested schema](#nestedatt--spec))

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

- `service_name` (String) The service name. To get a list of available services, use the DescribeVpcEndpointServicesrequest, or get the name from the service provider.

Optional:

- `dns_options` (Attributes) The DNS options for the endpoint. (see [below for nested schema](#nestedatt--spec--dns_options))
- `ip_address_type` (String) The IP address type for the endpoint.
- `policy_document` (String) (Interface and gateway endpoints) A policy to attach to the endpoint thatcontrols access to the service. The policy must be in valid JSON format.If this parameter is not specified, we attach a default policy that allowsfull access to the service.
- `private_dns_enabled` (Boolean) (Interface endpoint) Indicates whether to associate a private hosted zonewith the specified VPC. The private hosted zone contains a record set forthe default public DNS name for the service for the Region (for example,kinesis.us-east-1.amazonaws.com), which resolves to the private IP addressesof the endpoint network interfaces in the VPC. This enables you to make requeststo the default public DNS name for the service instead of the public DNSnames that are automatically generated by the VPC endpoint service.To use a private hosted zone, you must set the following VPC attributes totrue: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute toset the VPC attributes.Default: true
- `route_table_i_ds` (List of String) (Gateway endpoint) One or more route table IDs.
- `route_table_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--route_table_refs))
- `security_group_i_ds` (List of String) (Interface endpoint) The ID of one or more security groups to associate withthe endpoint network interface.
- `security_group_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--security_group_refs))
- `subnet_i_ds` (List of String) (Interface and Gateway Load Balancer endpoints) The ID of one or more subnetsin which to create an endpoint network interface. For a Gateway Load Balancerendpoint, you can specify one subnet only.
- `subnet_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--subnet_refs))
- `tags` (Attributes List) The tags. The value parameter is required, but if you don't want the tagto have a value, specify the parameter with no value, and we set the valueto an empty string. (see [below for nested schema](#nestedatt--spec--tags))
- `vpc_endpoint_type` (String) The type of endpoint.Default: Gateway
- `vpc_id` (String) The ID of the VPC in which the endpoint will be used.
- `vpc_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--vpc_ref))

<a id="nestedatt--spec--dns_options"></a>
### Nested Schema for `spec.dns_options`

Optional:

- `dns_record_ip_type` (String)


<a id="nestedatt--spec--route_table_refs"></a>
### Nested Schema for `spec.route_table_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--route_table_refs--from))

<a id="nestedatt--spec--route_table_refs--from"></a>
### Nested Schema for `spec.route_table_refs.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--security_group_refs"></a>
### Nested Schema for `spec.security_group_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--security_group_refs--from))

<a id="nestedatt--spec--security_group_refs--from"></a>
### Nested Schema for `spec.security_group_refs.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--subnet_refs"></a>
### Nested Schema for `spec.subnet_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--subnet_refs--from))

<a id="nestedatt--spec--subnet_refs--from"></a>
### Nested Schema for `spec.subnet_refs.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--spec--vpc_ref"></a>
### Nested Schema for `spec.vpc_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--vpc_ref--from))

<a id="nestedatt--spec--vpc_ref--from"></a>
### Nested Schema for `spec.vpc_ref.from`

Optional:

- `name` (String)