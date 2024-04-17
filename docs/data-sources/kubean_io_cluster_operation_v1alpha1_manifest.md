---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kubean_io_cluster_operation_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "kubean.io"
description: |-
  ClusterOperation represents the desire state and status of a member cluster.
---

# k8s_kubean_io_cluster_operation_v1alpha1_manifest (Data Source)

ClusterOperation represents the desire state and status of a member cluster.

## Example Usage

```terraform
data "k8s_kubean_io_cluster_operation_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of a member cluster. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `action` (String)
- `action_type` (String)
- `cluster` (String) Cluster the name of Cluster.kubean.io.
- `image` (String)

Optional:

- `action_source` (String)
- `action_source_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--action_source_ref))
- `active_deadline_seconds` (Number)
- `entrypoint_sh_ref` (Attributes) EntrypointSHRef will be filled by operator when it renders entrypoint.sh. (see [below for nested schema](#nestedatt--spec--entrypoint_sh_ref))
- `extra_args` (String)
- `hosts_conf_ref` (Attributes) HostsConfRef will be filled by operator when it performs backup. (see [below for nested schema](#nestedatt--spec--hosts_conf_ref))
- `post_hook` (Attributes List) (see [below for nested schema](#nestedatt--spec--post_hook))
- `pre_hook` (Attributes List) (see [below for nested schema](#nestedatt--spec--pre_hook))
- `resources` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--resources))
- `ssh_auth_ref` (Attributes) SSHAuthRef will be filled by operator when it performs backup. (see [below for nested schema](#nestedatt--spec--ssh_auth_ref))
- `vars_conf_ref` (Attributes) VarsConfRef will be filled by operator when it performs backup. (see [below for nested schema](#nestedatt--spec--vars_conf_ref))

<a id="nestedatt--spec--action_source_ref"></a>
### Nested Schema for `spec.action_source_ref`

Required:

- `name` (String)
- `namespace` (String)


<a id="nestedatt--spec--entrypoint_sh_ref"></a>
### Nested Schema for `spec.entrypoint_sh_ref`

Required:

- `name` (String)
- `namespace` (String)


<a id="nestedatt--spec--hosts_conf_ref"></a>
### Nested Schema for `spec.hosts_conf_ref`

Required:

- `name` (String)
- `namespace` (String)


<a id="nestedatt--spec--post_hook"></a>
### Nested Schema for `spec.post_hook`

Required:

- `action` (String)
- `action_type` (String)

Optional:

- `action_source` (String)
- `action_source_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--post_hook--action_source_ref))
- `extra_args` (String)

<a id="nestedatt--spec--post_hook--action_source_ref"></a>
### Nested Schema for `spec.post_hook.action_source_ref`

Required:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--pre_hook"></a>
### Nested Schema for `spec.pre_hook`

Required:

- `action` (String)
- `action_type` (String)

Optional:

- `action_source` (String)
- `action_source_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--pre_hook--action_source_ref))
- `extra_args` (String)

<a id="nestedatt--spec--pre_hook--action_source_ref"></a>
### Nested Schema for `spec.pre_hook.action_source_ref`

Required:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--resources"></a>
### Nested Schema for `spec.resources`

Optional:

- `claims` (Attributes List) Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers. (see [below for nested schema](#nestedatt--spec--resources--claims))
- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

<a id="nestedatt--spec--resources--claims"></a>
### Nested Schema for `spec.resources.claims`

Required:

- `name` (String) Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.



<a id="nestedatt--spec--ssh_auth_ref"></a>
### Nested Schema for `spec.ssh_auth_ref`

Required:

- `name` (String)
- `namespace` (String)


<a id="nestedatt--spec--vars_conf_ref"></a>
### Nested Schema for `spec.vars_conf_ref`

Required:

- `name` (String)
- `namespace` (String)