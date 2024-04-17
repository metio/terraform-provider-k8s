---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "anywhere.eks.amazonaws.com"
description: |-
  CloudStackMachineConfig is the Schema for the cloudstackmachineconfigs API.
---

# k8s_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest (Data Source)

CloudStackMachineConfig is the Schema for the cloudstackmachineconfigs API.

## Example Usage

```terraform
data "k8s_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) CloudStackMachineConfigSpec defines the desired state of CloudStackMachineConfig. (see [below for nested schema](#nestedatt--spec))

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

- `compute_offering` (Attributes) ComputeOffering refers to a compute offering which has been previously registered in CloudStack. It represents a VM’s instance size including number of CPU’s, memory, and CPU speed. It can either be specified as a UUID or name (see [below for nested schema](#nestedatt--spec--compute_offering))
- `template` (Attributes) Template refers to a VM image template which has been previously registered in CloudStack. It can either be specified as a UUID or name. When using a template name it must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27. (see [below for nested schema](#nestedatt--spec--template))

Optional:

- `affinity` (String) Defaults to 'no'. Can be 'pro' or 'anti'. If set to 'pro' or 'anti', will create an affinity group per machine set of the corresponding type
- `affinity_group_ids` (List of String) AffinityGroupIds allows users to pass in a list of UUIDs for previously-created Affinity Groups. Any VM’s created with this spec will be added to the affinity group, which will dictate which physical host(s) they can be placed on. Affinity groups can be type “affinity” or “anti-affinity” in CloudStack. If they are type “anti-affinity”, all VM’s in the group must be on separate physical hosts for high availability. If they are type “affinity”, all VM’s in the group must be on the same physical host for improved performance
- `disk_offering` (Attributes) DiskOffering refers to a disk offering which has been previously registered in CloudStack. It represents a disk offering with pre-defined size or custom specified disk size. It can either be specified as a UUID or name (see [below for nested schema](#nestedatt--spec--disk_offering))
- `symlinks` (Map of String) Symlinks create soft symbolic links folders. One use case is to use data disk to store logs
- `user_custom_details` (Map of String) UserCustomDetails allows users to pass in non-standard key value inputs, outside those defined [here](https://github.com/shapeblue/cloudstack/blob/main/api/src/main/java/com/cloud/vm/VmDetailConstants.java)
- `users` (Attributes List) Users consists of an array of objects containing the username, as well as a list of their public keys. These users will be authorized to ssh into the machines (see [below for nested schema](#nestedatt--spec--users))

<a id="nestedatt--spec--compute_offering"></a>
### Nested Schema for `spec.compute_offering`

Optional:

- `id` (String) Id of a resource in the CloudStack environment. Mutually exclusive with Name
- `name` (String) Name of a resource in the CloudStack environment. Mutually exclusive with Id


<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Optional:

- `id` (String) Id of a resource in the CloudStack environment. Mutually exclusive with Name
- `name` (String) Name of a resource in the CloudStack environment. Mutually exclusive with Id


<a id="nestedatt--spec--disk_offering"></a>
### Nested Schema for `spec.disk_offering`

Required:

- `device` (String) device name of the disk offering in VM, shows up in lsblk command
- `filesystem` (String) filesystem used to mkfs in disk offering partition
- `label` (String) disk label used to label disk partition
- `mount_path` (String) path the filesystem will use to mount in VM

Optional:

- `custom_size_in_gb` (Number) disk size in GB, > 0 for customized disk offering; = 0 for non-customized disk offering
- `id` (String) Id of a resource in the CloudStack environment. Mutually exclusive with Name
- `name` (String) Name of a resource in the CloudStack environment. Mutually exclusive with Id


<a id="nestedatt--spec--users"></a>
### Nested Schema for `spec.users`

Required:

- `name` (String)
- `ssh_authorized_keys` (List of String)