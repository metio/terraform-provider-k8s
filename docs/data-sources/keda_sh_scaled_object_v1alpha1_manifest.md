---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_keda_sh_scaled_object_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "keda.sh"
description: |-
  ScaledObject is a specification for a ScaledObject resource
---

# k8s_keda_sh_scaled_object_v1alpha1_manifest (Data Source)

ScaledObject is a specification for a ScaledObject resource

## Example Usage

```terraform
data "k8s_keda_sh_scaled_object_v1alpha1_manifest" "example" {
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
- `spec` (Attributes) ScaledObjectSpec is the spec for a ScaledObject resource (see [below for nested schema](#nestedatt--spec))

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

- `scale_target_ref` (Attributes) ScaleTarget holds the reference to the scale target Object (see [below for nested schema](#nestedatt--spec--scale_target_ref))
- `triggers` (Attributes List) (see [below for nested schema](#nestedatt--spec--triggers))

Optional:

- `advanced` (Attributes) AdvancedConfig specifies advance scaling options (see [below for nested schema](#nestedatt--spec--advanced))
- `cooldown_period` (Number)
- `fallback` (Attributes) Fallback is the spec for fallback options (see [below for nested schema](#nestedatt--spec--fallback))
- `idle_replica_count` (Number)
- `max_replica_count` (Number)
- `min_replica_count` (Number)
- `polling_interval` (Number)

<a id="nestedatt--spec--scale_target_ref"></a>
### Nested Schema for `spec.scale_target_ref`

Required:

- `name` (String)

Optional:

- `api_version` (String)
- `env_source_container_name` (String)
- `kind` (String)


<a id="nestedatt--spec--triggers"></a>
### Nested Schema for `spec.triggers`

Required:

- `metadata` (Map of String)
- `type` (String)

Optional:

- `authentication_ref` (Attributes) AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object thatis used to authenticate the scaler with the environment (see [below for nested schema](#nestedatt--spec--triggers--authentication_ref))
- `metric_type` (String) MetricTargetType specifies the type of metric being targeted, and should be either'Value', 'AverageValue', or 'Utilization'
- `name` (String)
- `use_cached_metrics` (Boolean)

<a id="nestedatt--spec--triggers--authentication_ref"></a>
### Nested Schema for `spec.triggers.authentication_ref`

Required:

- `name` (String)

Optional:

- `kind` (String) Kind of the resource being referred to. Defaults to TriggerAuthentication.



<a id="nestedatt--spec--advanced"></a>
### Nested Schema for `spec.advanced`

Optional:

- `horizontal_pod_autoscaler_config` (Attributes) HorizontalPodAutoscalerConfig specifies horizontal scale config (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config))
- `restore_to_original_replica_count` (Boolean)
- `scaling_modifiers` (Attributes) ScalingModifiers describes advanced scaling logic options like formula (see [below for nested schema](#nestedatt--spec--advanced--scaling_modifiers))

<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config`

Optional:

- `behavior` (Attributes) HorizontalPodAutoscalerBehavior configures the scaling behavior of the targetin both Up and Down directions (scaleUp and scaleDown fields respectively). (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config--behavior))
- `name` (String)

<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config--behavior"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config.behavior`

Optional:

- `scale_down` (Attributes) scaleDown is scaling policy for scaling Down.If not set, the default value is to allow to scale down to minReplicas pods, with a300 second stabilization window (i.e., the highest recommendation forthe last 300sec is used). (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_down))
- `scale_up` (Attributes) scaleUp is scaling policy for scaling Up.If not set, the default value is the higher of:  * increase no more than 4 pods per 60 seconds  * double the number of pods per 60 secondsNo stabilization is used. (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_up))

<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_down"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config.name.scale_down`

Optional:

- `policies` (Attributes List) policies is a list of potential scaling polices which can be used during scaling.At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_down--policies))
- `select_policy` (String) selectPolicy is used to specify which policy should be used.If not set, the default value Max is used.
- `stabilization_window_seconds` (Number) stabilizationWindowSeconds is the number of seconds for which past recommendations should beconsidered while scaling up or scaling down.StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour).If not set, use the default values:- For scale up: 0 (i.e. no stabilization is done).- For scale down: 300 (i.e. the stabilization window is 300 seconds long).

<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_down--policies"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config.name.scale_down.policies`

Required:

- `period_seconds` (Number) periodSeconds specifies the window of time for which the policy should hold true.PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
- `type` (String) type is used to specify the scaling policy.
- `value` (Number) value contains the amount of change which is permitted by the policy.It must be greater than zero



<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_up"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config.name.scale_up`

Optional:

- `policies` (Attributes List) policies is a list of potential scaling polices which can be used during scaling.At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid (see [below for nested schema](#nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_up--policies))
- `select_policy` (String) selectPolicy is used to specify which policy should be used.If not set, the default value Max is used.
- `stabilization_window_seconds` (Number) stabilizationWindowSeconds is the number of seconds for which past recommendations should beconsidered while scaling up or scaling down.StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour).If not set, use the default values:- For scale up: 0 (i.e. no stabilization is done).- For scale down: 300 (i.e. the stabilization window is 300 seconds long).

<a id="nestedatt--spec--advanced--horizontal_pod_autoscaler_config--name--scale_up--policies"></a>
### Nested Schema for `spec.advanced.horizontal_pod_autoscaler_config.name.scale_up.policies`

Required:

- `period_seconds` (Number) periodSeconds specifies the window of time for which the policy should hold true.PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
- `type` (String) type is used to specify the scaling policy.
- `value` (Number) value contains the amount of change which is permitted by the policy.It must be greater than zero





<a id="nestedatt--spec--advanced--scaling_modifiers"></a>
### Nested Schema for `spec.advanced.scaling_modifiers`

Optional:

- `activation_target` (String)
- `formula` (String)
- `metric_type` (String) MetricTargetType specifies the type of metric being targeted, and should be either'Value', 'AverageValue', or 'Utilization'
- `target` (String)



<a id="nestedatt--spec--fallback"></a>
### Nested Schema for `spec.fallback`

Required:

- `failure_threshold` (Number)
- `replicas` (Number)