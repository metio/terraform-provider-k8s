---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_externaldns_k8s_io_dns_endpoint_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "externaldns.k8s.io"
description: |-
  
---

# k8s_externaldns_k8s_io_dns_endpoint_v1alpha1 (Resource)



## Example Usage

```terraform
resource "k8s_externaldns_k8s_io_dns_endpoint_v1alpha1" "minimal" {
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

- `spec` (Attributes) DNSEndpointSpec defines the desired state of DNSEndpoint (see [below for nested schema](#nestedatt--spec))

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

- `endpoints` (Attributes List) (see [below for nested schema](#nestedatt--spec--endpoints))

<a id="nestedatt--spec--endpoints"></a>
### Nested Schema for `spec.endpoints`

Optional:

- `dns_name` (String) The hostname of the DNS record
- `labels` (Map of String) Labels stores labels defined for the Endpoint
- `provider_specific` (Attributes List) ProviderSpecific stores provider specific config (see [below for nested schema](#nestedatt--spec--endpoints--provider_specific))
- `record_ttl` (Number) TTL for the record
- `record_type` (String) RecordType type of record, e.g. CNAME, A, SRV, TXT etc
- `set_identifier` (String) Identifier to distinguish multiple records with the same name and type (e.g. Route53 records with routing policies other than 'simple')
- `targets` (List of String) The targets the DNS record points to

<a id="nestedatt--spec--endpoints--provider_specific"></a>
### Nested Schema for `spec.endpoints.provider_specific`

Optional:

- `name` (String)
- `value` (String)


