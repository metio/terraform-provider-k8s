---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_cosmos_db_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  CosmosDB is the Schema for the cosmosdbs API
---

# k8s_azure_microsoft_com_cosmos_db_v1alpha1_manifest (Data Source)

CosmosDB is the Schema for the cosmosdbs API

## Example Usage

```terraform
data "k8s_azure_microsoft_com_cosmos_db_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) CosmosDBSpec defines the desired state of CosmosDB (see [below for nested schema](#nestedatt--spec))

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

- `resource_group` (String)

Optional:

- `ip_rules` (List of String)
- `key_vault_to_store_secrets` (String)
- `kind` (String) CosmosDBKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is GlobalDocumentDBKind.
- `location` (String) Location is the Azure location where the CosmosDB exists
- `locations` (Attributes List) (see [below for nested schema](#nestedatt--spec--locations))
- `properties` (Attributes) CosmosDBProperties the CosmosDBProperties of CosmosDB. (see [below for nested schema](#nestedatt--spec--properties))
- `virtual_network_rules` (Attributes List) (see [below for nested schema](#nestedatt--spec--virtual_network_rules))

<a id="nestedatt--spec--locations"></a>
### Nested Schema for `spec.locations`

Required:

- `failover_priority` (Number)
- `location_name` (String)

Optional:

- `is_zone_redundant` (Boolean)


<a id="nestedatt--spec--properties"></a>
### Nested Schema for `spec.properties`

Optional:

- `capabilities` (Attributes List) (see [below for nested schema](#nestedatt--spec--properties--capabilities))
- `database_account_offer_type` (String) DatabaseAccountOfferType - The offer type for the Cosmos DB database account.
- `enable_multiple_write_locations` (Boolean)
- `is_virtual_network_filter_enabled` (Boolean) IsVirtualNetworkFilterEnabled - Flag to indicate whether to enable/disable Virtual Network ACL rules.
- `mongo_db_version` (String)

<a id="nestedatt--spec--properties--capabilities"></a>
### Nested Schema for `spec.properties.capabilities`

Optional:

- `name` (String) Name *CosmosCapability 'json:'name,omitempty''



<a id="nestedatt--spec--virtual_network_rules"></a>
### Nested Schema for `spec.virtual_network_rules`

Optional:

- `ignore_missing_v_net_service_endpoint` (Boolean) IgnoreMissingVNetServiceEndpoint - Create firewall rule before the virtual network has vnet service endpoint enabled.
- `subnet_id` (String) ID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.