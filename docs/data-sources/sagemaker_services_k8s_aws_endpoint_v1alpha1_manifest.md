---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "sagemaker.services.k8s.aws"
description: |-
  Endpoint is the Schema for the Endpoints API
---

# k8s_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest (Data Source)

Endpoint is the Schema for the Endpoints API

## Example Usage

```terraform
data "k8s_sagemaker_services_k8s_aws_endpoint_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) EndpointSpec defines the desired state of Endpoint. A hosted endpoint for real-time inference. (see [below for nested schema](#nestedatt--spec))

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

- `endpoint_config_name` (String) The name of an endpoint configuration. For more information, see CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html).
- `endpoint_name` (String) The name of the endpoint.The name must be unique within an Amazon Web Services Region in your Amazon Web Services account. The name is case-insensitive in CreateEndpoint, but the case is preserved and must be matched in InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html).

Optional:

- `deployment_config` (Attributes) The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations. (see [below for nested schema](#nestedatt--spec--deployment_config))
- `tags` (Attributes List) An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--deployment_config"></a>
### Nested Schema for `spec.deployment_config`

Optional:

- `auto_rollback_configuration` (Attributes) Automatic rollback configuration for handling endpoint deployment failures and recovery. (see [below for nested schema](#nestedatt--spec--deployment_config--auto_rollback_configuration))
- `blue_green_update_policy` (Attributes) Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default. (see [below for nested schema](#nestedatt--spec--deployment_config--blue_green_update_policy))
- `rolling_update_policy` (Attributes) Specifies a rolling deployment strategy for updating a SageMaker endpoint. (see [below for nested schema](#nestedatt--spec--deployment_config--rolling_update_policy))

<a id="nestedatt--spec--deployment_config--auto_rollback_configuration"></a>
### Nested Schema for `spec.deployment_config.auto_rollback_configuration`

Optional:

- `alarms` (Attributes List) (see [below for nested schema](#nestedatt--spec--deployment_config--auto_rollback_configuration--alarms))

<a id="nestedatt--spec--deployment_config--auto_rollback_configuration--alarms"></a>
### Nested Schema for `spec.deployment_config.auto_rollback_configuration.alarms`

Optional:

- `alarm_name` (String)



<a id="nestedatt--spec--deployment_config--blue_green_update_policy"></a>
### Nested Schema for `spec.deployment_config.blue_green_update_policy`

Optional:

- `maximum_execution_timeout_in_seconds` (Number)
- `termination_wait_in_seconds` (Number)
- `traffic_routing_configuration` (Attributes) Defines the traffic routing strategy during an endpoint deployment to shift traffic from the old fleet to the new fleet. (see [below for nested schema](#nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration))

<a id="nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration"></a>
### Nested Schema for `spec.deployment_config.blue_green_update_policy.traffic_routing_configuration`

Optional:

- `canary_size` (Attributes) Specifies the type and size of the endpoint capacity to activate for a blue/green deployment, a rolling deployment, or a rollback strategy. You can specify your batches as either instance count or the overall percentage or your fleet. For a rollback strategy, if you don't specify the fields in this object, or if you set the Value to 100%, then SageMaker uses a blue/green rollback strategy and rolls all traffic back to the blue fleet. (see [below for nested schema](#nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration--canary_size))
- `linear_step_size` (Attributes) Specifies the type and size of the endpoint capacity to activate for a blue/green deployment, a rolling deployment, or a rollback strategy. You can specify your batches as either instance count or the overall percentage or your fleet. For a rollback strategy, if you don't specify the fields in this object, or if you set the Value to 100%, then SageMaker uses a blue/green rollback strategy and rolls all traffic back to the blue fleet. (see [below for nested schema](#nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration--linear_step_size))
- `type_` (String)
- `wait_interval_in_seconds` (Number)

<a id="nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration--canary_size"></a>
### Nested Schema for `spec.deployment_config.blue_green_update_policy.traffic_routing_configuration.canary_size`

Optional:

- `type_` (String)
- `value` (Number)


<a id="nestedatt--spec--deployment_config--blue_green_update_policy--traffic_routing_configuration--linear_step_size"></a>
### Nested Schema for `spec.deployment_config.blue_green_update_policy.traffic_routing_configuration.linear_step_size`

Optional:

- `type_` (String)
- `value` (Number)




<a id="nestedatt--spec--deployment_config--rolling_update_policy"></a>
### Nested Schema for `spec.deployment_config.rolling_update_policy`

Optional:

- `maximum_batch_size` (Attributes) Specifies the type and size of the endpoint capacity to activate for a blue/green deployment, a rolling deployment, or a rollback strategy. You can specify your batches as either instance count or the overall percentage or your fleet. For a rollback strategy, if you don't specify the fields in this object, or if you set the Value to 100%, then SageMaker uses a blue/green rollback strategy and rolls all traffic back to the blue fleet. (see [below for nested schema](#nestedatt--spec--deployment_config--rolling_update_policy--maximum_batch_size))
- `maximum_execution_timeout_in_seconds` (Number)
- `rollback_maximum_batch_size` (Attributes) Specifies the type and size of the endpoint capacity to activate for a blue/green deployment, a rolling deployment, or a rollback strategy. You can specify your batches as either instance count or the overall percentage or your fleet. For a rollback strategy, if you don't specify the fields in this object, or if you set the Value to 100%, then SageMaker uses a blue/green rollback strategy and rolls all traffic back to the blue fleet. (see [below for nested schema](#nestedatt--spec--deployment_config--rolling_update_policy--rollback_maximum_batch_size))
- `wait_interval_in_seconds` (Number)

<a id="nestedatt--spec--deployment_config--rolling_update_policy--maximum_batch_size"></a>
### Nested Schema for `spec.deployment_config.rolling_update_policy.maximum_batch_size`

Optional:

- `type_` (String)
- `value` (Number)


<a id="nestedatt--spec--deployment_config--rolling_update_policy--rollback_maximum_batch_size"></a>
### Nested Schema for `spec.deployment_config.rolling_update_policy.rollback_maximum_batch_size`

Optional:

- `type_` (String)
- `value` (Number)




<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
