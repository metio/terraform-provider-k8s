---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sagemaker_services_k8s_aws_app_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "sagemaker.services.k8s.aws/v1alpha1"
description: |-
  App is the Schema for the Apps API
---

# k8s_sagemaker_services_k8s_aws_app_v1alpha1 (Resource)

App is the Schema for the Apps API

## Example Usage

```terraform
resource "k8s_sagemaker_services_k8s_aws_app_v1alpha1" "minimal" {
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

- `spec` (Attributes) AppSpec defines the desired state of App. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `app_name` (String) The name of the app.
- `app_type` (String) The type of app. Supported apps are JupyterServer and KernelGateway. TensorBoard is not supported.
- `domain_id` (String) The domain ID.
- `user_profile_name` (String) The user profile name.

Optional:

- `resource_spec` (Attributes) The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.  The value of InstanceType passed as part of the ResourceSpec in the CreateApp call overrides the value passed as part of the ResourceSpec configured for the user profile or the domain. If InstanceType is not specified in any of those three ResourceSpec values for a KernelGateway app, the CreateApp call fails with a request validation error. (see [below for nested schema](#nestedatt--spec--resource_spec))
- `tags` (Attributes List) Each tag consists of a key and an optional value. Tag keys must be unique per resource. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--resource_spec"></a>
### Nested Schema for `spec.resource_spec`

Optional:

- `instance_type` (String)
- `lifecycle_config_arn` (String)
- `sage_maker_image_arn` (String)
- `sage_maker_image_version_arn` (String)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)

