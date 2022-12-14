---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "sagemaker.services.k8s.aws"
description: |-
  EndpointConfig is the Schema for the EndpointConfigs API
---

# k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1 (Resource)

EndpointConfig is the Schema for the EndpointConfigs API

## Example Usage

```terraform
resource "k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1" "minimal" {
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

- `spec` (Attributes) EndpointConfigSpec defines the desired state of EndpointConfig. (see [below for nested schema](#nestedatt--spec))

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

- `endpoint_config_name` (String) The name of the endpoint configuration. You specify this name in a CreateEndpoint request.
- `production_variants` (Attributes List) An list of ProductionVariant objects, one for each model that you want to host at this endpoint. (see [below for nested schema](#nestedatt--spec--production_variants))

Optional:

- `async_inference_config` (Attributes) Specifies configuration for how an endpoint performs asynchronous inference. This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html). (see [below for nested schema](#nestedatt--spec--async_inference_config))
- `data_capture_config` (Attributes) Configuration to control how SageMaker captures inference data. (see [below for nested schema](#nestedatt--spec--data_capture_config))
- `kms_key_id` (String) The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.  The KmsKeyId can be any of the following formats:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  * Alias name: alias/ExampleAlias  * Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias  The KMS key policy must grant permission to the IAM role that you specify in your CreateEndpoint, UpdateEndpoint requests. For more information, refer to the Amazon Web Services Key Management Service section Using Key Policies in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)  Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a KmsKeyId when using an instance type with local storage. If any of the models that you specify in the ProductionVariants parameter use nitro-based instances with local storage, do not specify a value for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any nitro-based instances with local storage, the call to CreateEndpointConfig fails.  For a list of instance types that support local instance storage, see Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).  For more information about local instance storage encryption, see SSD Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).
- `tags` (Attributes List) An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--production_variants"></a>
### Nested Schema for `spec.production_variants`

Optional:

- `accelerator_type` (String)
- `container_startup_health_check_timeout_in_seconds` (Number)
- `core_dump_config` (Attributes) Specifies configuration for a core dump from the model container when the process crashes. (see [below for nested schema](#nestedatt--spec--production_variants--core_dump_config))
- `initial_instance_count` (Number)
- `initial_variant_weight` (Number)
- `instance_type` (String)
- `model_data_download_timeout_in_seconds` (Number)
- `model_name` (String)
- `variant_name` (String)
- `volume_size_in_gb` (Number)

<a id="nestedatt--spec--production_variants--core_dump_config"></a>
### Nested Schema for `spec.production_variants.core_dump_config`

Optional:

- `destination_s3_uri` (String)
- `kms_key_id` (String)



<a id="nestedatt--spec--async_inference_config"></a>
### Nested Schema for `spec.async_inference_config`

Optional:

- `client_config` (Attributes) Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference. (see [below for nested schema](#nestedatt--spec--async_inference_config--client_config))
- `output_config` (Attributes) Specifies the configuration for asynchronous inference invocation outputs. (see [below for nested schema](#nestedatt--spec--async_inference_config--output_config))

<a id="nestedatt--spec--async_inference_config--client_config"></a>
### Nested Schema for `spec.async_inference_config.client_config`

Optional:

- `max_concurrent_invocations_per_instance` (Number)


<a id="nestedatt--spec--async_inference_config--output_config"></a>
### Nested Schema for `spec.async_inference_config.output_config`

Optional:

- `kms_key_id` (String)
- `notification_config` (Attributes) Specifies the configuration for notifications of inference results for asynchronous inference. (see [below for nested schema](#nestedatt--spec--async_inference_config--output_config--notification_config))
- `s3_output_path` (String)

<a id="nestedatt--spec--async_inference_config--output_config--notification_config"></a>
### Nested Schema for `spec.async_inference_config.output_config.s3_output_path`

Optional:

- `error_topic` (String)
- `success_topic` (String)




<a id="nestedatt--spec--data_capture_config"></a>
### Nested Schema for `spec.data_capture_config`

Optional:

- `capture_content_type_header` (Attributes) Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data. (see [below for nested schema](#nestedatt--spec--data_capture_config--capture_content_type_header))
- `capture_options` (Attributes List) (see [below for nested schema](#nestedatt--spec--data_capture_config--capture_options))
- `destination_s3_uri` (String)
- `enable_capture` (Boolean)
- `initial_sampling_percentage` (Number)
- `kms_key_id` (String)

<a id="nestedatt--spec--data_capture_config--capture_content_type_header"></a>
### Nested Schema for `spec.data_capture_config.capture_content_type_header`

Optional:

- `csv_content_types` (List of String)
- `json_content_types` (List of String)


<a id="nestedatt--spec--data_capture_config--capture_options"></a>
### Nested Schema for `spec.data_capture_config.capture_options`

Optional:

- `capture_mode` (String)



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)


