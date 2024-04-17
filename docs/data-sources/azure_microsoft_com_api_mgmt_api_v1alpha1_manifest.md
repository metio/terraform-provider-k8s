---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "azure.microsoft.com"
description: |-
  
---

# k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) APIMgmtSpec defines the desired state of APIMgmt (see [below for nested schema](#nestedatt--spec))

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

- `api_id` (String)
- `api_service` (String)
- `location` (String)
- `properties` (Attributes) (see [below for nested schema](#nestedatt--spec--properties))
- `resource_group` (String)

<a id="nestedatt--spec--properties"></a>
### Nested Schema for `spec.properties`

Optional:

- `api_revision` (String) APIRevision - Describes the Revision of the Api. If no value is provided, default revision 1 is created
- `api_revision_description` (String) APIRevisionDescription - Description of the Api Revision.
- `api_version` (String) APIVersion - Indicates the Version identifier of the API if the API is versioned
- `api_version_description` (String) APIVersionDescription - Description of the Api Version.
- `api_version_set_id` (String) APIVersionSetID - A resource identifier for the related ApiVersionSet.
- `api_version_sets` (Attributes) APIVersionSet - APIVersionSetContractDetails an API Version Set contains the common configuration for a set of API versions. (see [below for nested schema](#nestedatt--spec--properties--api_version_sets))
- `description` (String) Description - Description of the API. May include HTML formatting tags.
- `display_name` (String) DisplayName - API name. Must be 1 to 300 characters long.
- `format` (String) Format - Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'
- `is_current` (Boolean) IsCurrent - Indicates if API revision is current api revision.
- `is_online` (Boolean) IsOnline - READ-ONLY; Indicates if API revision is accessible via the gateway.
- `path` (String) Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
- `protocols` (List of String) Protocols - Describes on which protocols the operations in this API can be invoked.
- `service_url` (String) ServiceURL - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
- `source_api_id` (String) SourceAPIID - API identifier of the source API.
- `subscription_required` (Boolean) SubscriptionRequired - Specifies whether an API or Product subscription is required for accessing the API.

<a id="nestedatt--spec--properties--api_version_sets"></a>
### Nested Schema for `spec.properties.api_version_sets`

Optional:

- `description` (String) Description - Description of API Version Set.
- `id` (String) ID - Identifier for existing API Version Set. Omit this value to create a new Version Set.
- `name` (String) Name - The display Name of the API Version Set.