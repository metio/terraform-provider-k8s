---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "security-profiles-operator.x-k8s.io"
description: |-
  SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp
---

# k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest (Data Source)

SeccompProfile is a cluster level specification for a seccomp profile. See https://github.com/opencontainers/runtime-spec/blob/master/config-linux.md#seccomp

## Example Usage

```terraform
data "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    default_action = "SCMP_ACT_TRAP"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) SeccompProfileSpec defines the desired state of SeccompProfile. (see [below for nested schema](#nestedatt--spec))

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

- `default_action` (String) the default action for seccomp

Optional:

- `architectures` (List of String) the architecture used for system calls
- `base_profile_name` (String) BaseProfileName is the name of base profile (in the same namespace) that will be unioned into this profile. Base profiles can be references as remote OCI artifacts as well when prefixed with 'oci://'.
- `disabled` (Boolean) Whether the profile is disabled and should be skipped during reconciliation.
- `flags` (List of String) list of flags to use with seccomp(2)
- `listener_metadata` (String) opaque data to pass to the seccomp agent
- `listener_path` (String) path of UNIX domain socket to contact a seccomp agent for SCMP_ACT_NOTIFY
- `syscalls` (Attributes List) match a syscall in seccomp. While this property is OPTIONAL, some values of defaultAction are not useful without syscalls entries. For example, if defaultAction is SCMP_ACT_KILL and syscalls is empty or unset, the kernel will kill the container process on its first syscall (see [below for nested schema](#nestedatt--spec--syscalls))

<a id="nestedatt--spec--syscalls"></a>
### Nested Schema for `spec.syscalls`

Required:

- `action` (String) the action for seccomp rules
- `names` (List of String) the names of the syscalls

Optional:

- `args` (Attributes List) the specific syscall in seccomp (see [below for nested schema](#nestedatt--spec--syscalls--args))
- `errno_ret` (Number) the errno return code to use. Some actions like SCMP_ACT_ERRNO and SCMP_ACT_TRACE allow to specify the errno code to return

<a id="nestedatt--spec--syscalls--args"></a>
### Nested Schema for `spec.syscalls.args`

Required:

- `index` (Number) the index for syscall arguments in seccomp
- `op` (String) the operator for syscall arguments in seccomp

Optional:

- `value` (Number) the value for syscall arguments in seccomp
- `value_two` (Number) the value for syscall arguments in seccomp
