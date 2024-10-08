---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "ec2.services.k8s.aws"
description: |-
  ElasticIPAddress is the Schema for the ElasticIPAddresses API
---

# k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest (Data Source)

ElasticIPAddress is the Schema for the ElasticIPAddresses API

## Example Usage

```terraform
data "k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ElasticIPAddressSpec defines the desired state of ElasticIPAddress. (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `address` (String) [EC2-VPC] The Elastic IP address to recover or an IPv4 address from an address pool.
- `customer_owned_i_pv4_pool` (String) The ID of a customer-owned address pool. Use this parameter to let Amazon EC2 select an address from the address pool. Alternatively, specify a specific address from the address pool.
- `network_border_group` (String) A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups. Use DescribeAvailabilityZones (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups. You cannot use a network border group with EC2 Classic. If you attempt this operation on EC2 Classic, you receive an InvalidParameterCombination error.
- `public_i_pv4_pool` (String) The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool. To specify a specific address from the address pool, use the Address parameter instead.
- `tags` (Attributes List) The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
