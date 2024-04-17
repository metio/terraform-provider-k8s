---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_storage_account_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  StorageAccount is the Schema for the storages API
---

# k8s_azure_microsoft_com_storage_account_v1alpha1_manifest (Data Source)

StorageAccount is the Schema for the storages API

## Example Usage

```terraform
data "k8s_azure_microsoft_com_storage_account_v1alpha1_manifest" "example" {
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

- `additional_resources` (Attributes) StorageAccountAdditionalResources holds the additional resources (see [below for nested schema](#nestedatt--additional_resources))
- `output` (Attributes) StorageAccountOutput is the object that contains the output from creating a Storage Account object (see [below for nested schema](#nestedatt--output))
- `spec` (Attributes) StorageAccountSpec defines the desired state of Storage (see [below for nested schema](#nestedatt--spec))

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


<a id="nestedatt--additional_resources"></a>
### Nested Schema for `additional_resources`

Optional:

- `secrets` (List of String)


<a id="nestedatt--output"></a>
### Nested Schema for `output`

Optional:

- `connection_string1` (String)
- `connection_string2` (String)
- `key1` (String)
- `key2` (String)
- `storage_account_name` (String)


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `resource_group` (String)

Optional:

- `access_tier` (String) StorageAccountAccessTier enumerates the values for access tier. Only one of the following access tiers may be specified. If none of the following access tiers is specified, the default one is Hot.
- `data_lake_enabled` (Boolean)
- `kind` (String) StorageAccountKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is StorageV2.
- `location` (String)
- `network_rule` (Attributes) (see [below for nested schema](#nestedatt--spec--network_rule))
- `sku` (Attributes) StorageAccountSku the SKU of the storage account. (see [below for nested schema](#nestedatt--spec--sku))
- `supports_https_traffic_only` (Boolean)

<a id="nestedatt--spec--network_rule"></a>
### Nested Schema for `spec.network_rule`

Optional:

- `bypass` (String) Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, 'Logging, Metrics'), or None to bypass none of those traffics. Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'
- `default_action` (String) DefaultAction - Specifies the default action of allow or deny when no other rules match. Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'
- `ip_rules` (Attributes List) IPRules - Sets the IP ACL rules (see [below for nested schema](#nestedatt--spec--network_rule--ip_rules))
- `virtual_network_rules` (Attributes List) VirtualNetworkRules - Sets the virtual network rules (see [below for nested schema](#nestedatt--spec--network_rule--virtual_network_rules))

<a id="nestedatt--spec--network_rule--ip_rules"></a>
### Nested Schema for `spec.network_rule.ip_rules`

Optional:

- `ip_address_or_range` (String) IPAddressOrRange - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.


<a id="nestedatt--spec--network_rule--virtual_network_rules"></a>
### Nested Schema for `spec.network_rule.virtual_network_rules`

Optional:

- `subnet_id` (String) SubnetId - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.



<a id="nestedatt--spec--sku"></a>
### Nested Schema for `spec.sku`

Optional:

- `name` (String) Name - The SKU name. Required for account creation; optional for update. Possible values include: 'Standard_LRS', 'Standard_GRS', 'Standard_RAGRS', 'Standard_ZRS', 'Premium_LRS', 'Premium_ZRS', 'Standard_GZRS', 'Standard_RAGZRS'. For the full list of allowed options, see: https://docs.microsoft.com/en-us/rest/api/storagerp/storageaccounts/create#skuname