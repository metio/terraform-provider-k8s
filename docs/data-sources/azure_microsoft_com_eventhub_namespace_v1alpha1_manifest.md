---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  EventhubNamespace is the Schema for the eventhubnamespaces API
---

# k8s_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest (Data Source)

EventhubNamespace is the Schema for the eventhubnamespaces API

## Example Usage

```terraform
data "k8s_azure_microsoft_com_eventhub_namespace_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) EventhubNamespaceSpec defines the desired state of EventhubNamespace (see [below for nested schema](#nestedatt--spec))

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

- `location` (String) INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'make' to regenerate code after modifying this file
- `resource_group` (String)

Optional:

- `network_rule` (Attributes) EventhubNamespaceNetworkRule defines the namespace network rule (see [below for nested schema](#nestedatt--spec--network_rule))
- `properties` (Attributes) EventhubNamespaceProperties defines the namespace properties (see [below for nested schema](#nestedatt--spec--properties))
- `sku` (Attributes) EventhubNamespaceSku defines the sku (see [below for nested schema](#nestedatt--spec--sku))

<a id="nestedatt--spec--network_rule"></a>
### Nested Schema for `spec.network_rule`

Optional:

- `default_action` (String) DefaultAction defined as a string
- `ip_rules` (Attributes List) IPRules - List of IpRules (see [below for nested schema](#nestedatt--spec--network_rule--ip_rules))
- `virtual_network_rules` (Attributes List) VirtualNetworkRules - List VirtualNetwork Rules (see [below for nested schema](#nestedatt--spec--network_rule--virtual_network_rules))

<a id="nestedatt--spec--network_rule--ip_rules"></a>
### Nested Schema for `spec.network_rule.ip_rules`

Optional:

- `ip_mask` (String) IPMask - IPv4 address 1.1.1.1 or CIDR notation 1.1.0.0/24


<a id="nestedatt--spec--network_rule--virtual_network_rules"></a>
### Nested Schema for `spec.network_rule.virtual_network_rules`

Optional:

- `ignore_missing_service_endpoint` (Boolean) IgnoreMissingVnetServiceEndpoint - Value that indicates whether to ignore missing VNet Service Endpoint
- `subnet_id` (String) Subnet - Full Resource ID of Virtual Network Subnet



<a id="nestedatt--spec--properties"></a>
### Nested Schema for `spec.properties`

Optional:

- `is_auto_inflate_enabled` (Boolean)
- `kafka_enabled` (Boolean)
- `maximum_throughput_units` (Number)


<a id="nestedatt--spec--sku"></a>
### Nested Schema for `spec.sku`

Optional:

- `capacity` (Number)
- `name` (String)
- `tier` (String)
