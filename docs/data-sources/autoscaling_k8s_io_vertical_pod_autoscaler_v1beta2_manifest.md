---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "autoscaling.k8s.io"
description: |-
  VerticalPodAutoscaler is the configuration for a vertical pod autoscaler, which automatically manages pod resources based on historical and real time resource utilization.
---

# k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest (Data Source)

VerticalPodAutoscaler is the configuration for a vertical pod autoscaler, which automatically manages pod resources based on historical and real time resource utilization.

## Example Usage

```terraform
data "k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      kind = "Deployment"
      name = "some-name"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Specification of the behavior of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status. (see [below for nested schema](#nestedatt--spec))

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

- `target_ref` (Attributes) TargetRef points to the controller managing the set of pods for the autoscaler to control - e.g. Deployment, StatefulSet. VerticalPodAutoscaler can be targeted at controller implementing scale subresource (the pod set is retrieved from the controller's ScaleStatus) or some well known controllers (e.g. for DaemonSet the pod set is read from the controller's spec). If VerticalPodAutoscaler cannot use specified target it will report ConfigUnsupported condition. Note that VerticalPodAutoscaler does not require full implementation of scale subresource - it will not use it to modify the replica count. The only thing retrieved is a label selector matching pods grouped by the target resource. (see [below for nested schema](#nestedatt--spec--target_ref))

Optional:

- `resource_policy` (Attributes) Controls how the autoscaler computes recommended resources. The resource policy may be used to set constraints on the recommendations for individual containers. If not specified, the autoscaler computes recommended resources for all containers in the pod, without additional constraints. (see [below for nested schema](#nestedatt--spec--resource_policy))
- `update_policy` (Attributes) Describes the rules on how changes are applied to the pods. If not specified, all fields in the 'PodUpdatePolicy' are set to their default values. (see [below for nested schema](#nestedatt--spec--update_policy))

<a id="nestedatt--spec--target_ref"></a>
### Nested Schema for `spec.target_ref`

Required:

- `kind` (String) Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names

Optional:

- `api_version` (String) API version of the referent


<a id="nestedatt--spec--resource_policy"></a>
### Nested Schema for `spec.resource_policy`

Optional:

- `container_policies` (Attributes List) Per-container resource policies. (see [below for nested schema](#nestedatt--spec--resource_policy--container_policies))

<a id="nestedatt--spec--resource_policy--container_policies"></a>
### Nested Schema for `spec.resource_policy.container_policies`

Optional:

- `container_name` (String) Name of the container or DefaultContainerResourcePolicy, in which case the policy is used by the containers that don't have their own policy specified.
- `max_allowed` (Map of String) Specifies the maximum amount of resources that will be recommended for the container. The default is no maximum.
- `min_allowed` (Map of String) Specifies the minimal amount of resources that will be recommended for the container. The default is no minimum.
- `mode` (String) Whether autoscaler is enabled for the container. The default is 'Auto'.



<a id="nestedatt--spec--update_policy"></a>
### Nested Schema for `spec.update_policy`

Optional:

- `update_mode` (String) Controls when autoscaler applies changes to the pod resources. The default is 'Auto'.