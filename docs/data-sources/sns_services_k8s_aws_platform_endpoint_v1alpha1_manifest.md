---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "sns.services.k8s.aws"
description: |-
  PlatformEndpoint is the Schema for the PlatformEndpoints API
---

# k8s_sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest (Data Source)

PlatformEndpoint is the Schema for the PlatformEndpoints API

## Example Usage

```terraform
data "k8s_sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) PlatformEndpointSpec defines the desired state of PlatformEndpoint. (see [below for nested schema](#nestedatt--spec))

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

- `platform_application_arn` (String) PlatformApplicationArn returned from CreatePlatformApplication is used to create a an endpoint.
- `token` (String) Unique identifier created by the notification service for an app on a device. The specific name for Token will vary, depending on which notification service is being used. For example, when using APNS as the notification service, you need the device token. Alternatively, when using GCM (Firebase Cloud Messaging) or ADM, the device token equivalent is called the registration ID.

Optional:

- `custom_user_data` (String) Arbitrary user data to associate with the endpoint. Amazon SNS does not use this data. The data must be in UTF-8 format and less than 2KB.
- `enabled` (String)
