---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ec2_services_k8s_aws_subnet_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "ec2.services.k8s.aws/v1alpha1"
description: |-
  Subnet is the Schema for the Subnets API
---

# k8s_ec2_services_k8s_aws_subnet_v1alpha1 (Resource)

Subnet is the Schema for the Subnets API

## Example Usage

```terraform
resource "k8s_ec2_services_k8s_aws_subnet_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) SubnetSpec defines the desired state of Subnet.  Describes a subnet. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `assign_i_pv6_address_on_creation` (Boolean)
- `availability_zone` (String) The Availability Zone or Local Zone for the subnet.  Default: Amazon Web Services selects one for you. If you create more than one subnet in your VPC, we do not necessarily select a different zone for each subnet.  To create a subnet in a Local Zone, set this value to the Local Zone ID, for example us-west-2-lax-1a. For information about the Regions that support Local Zones, see Available Regions (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions) in the Amazon Elastic Compute Cloud User Guide.  To create a subnet in an Outpost, set this value to the Availability Zone for the Outpost and specify the Outpost ARN.
- `availability_zone_id` (String) The AZ ID or the Local Zone ID of the subnet.
- `cidr_block` (String) The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24. We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18.  This parameter is not supported for an IPv6 only subnet.
- `customer_owned_i_pv4_pool` (String)
- `enable_dns64` (Boolean)
- `enable_resource_name_dnsa_record` (Boolean)
- `enable_resource_name_dnsaaaa_record` (Boolean)
- `hostname_type` (String)
- `ipv6_cidr_block` (String) The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length.  This parameter is required for an IPv6 only subnet.
- `ipv6_native` (Boolean) Indicates whether to create an IPv6 only subnet.
- `map_public_ip_on_launch` (Boolean)
- `outpost_arn` (String) The Amazon Resource Name (ARN) of the Outpost. If you specify an Outpost ARN, you must also specify the Availability Zone of the Outpost subnet.
- `route_table_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--route_table_refs))
- `route_tables` (List of String)
- `tags` (Attributes List) The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string. (see [below for nested schema](#nestedatt--spec--tags))
- `vpc_id` (String) The ID of the VPC.
- `vpc_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api (see [below for nested schema](#nestedatt--spec--vpc_ref))

<a id="nestedatt--spec--route_table_refs"></a>
### Nested Schema for `spec.route_table_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--route_table_refs--from))

<a id="nestedatt--spec--route_table_refs--from"></a>
### Nested Schema for `spec.route_table_refs.from`

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

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--vpc_ref--from))

<a id="nestedatt--spec--vpc_ref--from"></a>
### Nested Schema for `spec.vpc_ref.from`

Optional:

- `name` (String)

