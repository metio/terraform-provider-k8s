---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "anywhere.eks.amazonaws.com"
description: |-
  NutanixMachineConfig is the Schema for the nutanix machine configs API
---

# k8s_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest (Data Source)

NutanixMachineConfig is the Schema for the nutanix machine configs API

## Example Usage

```terraform
data "k8s_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) NutanixMachineConfigSpec defines the desired state of NutanixMachineConfig. (see [below for nested schema](#nestedatt--spec))

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

- `cluster` (Attributes) cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API. (see [below for nested schema](#nestedatt--spec--cluster))
- `image` (Attributes) image is to identify the OS image uploaded to the Prism Central (PC) The image identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API. It must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27. (see [below for nested schema](#nestedatt--spec--image))
- `memory_size` (String) memorySize is the memory size (in Quantity format) of the VM The minimum memorySize is 2Gi bytes
- `os_family` (String)
- `subnet` (Attributes) subnet is to identify the cluster's network subnet to use for the Machine's VM The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API. (see [below for nested schema](#nestedatt--spec--subnet))
- `system_disk_size` (String) systemDiskSize is size (in Quantity format) of the system disk of the VM The minimum systemDiskSize is 20Gi bytes
- `vcpu_sockets` (Number) vcpuSockets is the number of vCPU sockets of the VM
- `vcpus_per_socket` (Number) vcpusPerSocket is the number of vCPUs per socket of the VM

Optional:

- `additional_categories` (Attributes List) additionalCategories is a list of optional categories to be added to the VM. Categories must be created in Prism Central before they can be used. (see [below for nested schema](#nestedatt--spec--additional_categories))
- `project` (Attributes) Project is an optional property that specifies the Prism Central project so that machine resources can be linked to it. The project identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API. (see [below for nested schema](#nestedatt--spec--project))
- `users` (Attributes List) (see [below for nested schema](#nestedatt--spec--users))

<a id="nestedatt--spec--cluster"></a>
### Nested Schema for `spec.cluster`

Required:

- `type` (String) Type is the identifier type to use for this resource.

Optional:

- `name` (String) name is the resource name in the PC
- `uuid` (String) uuid is the UUID of the resource in the PC.


<a id="nestedatt--spec--image"></a>
### Nested Schema for `spec.image`

Required:

- `type` (String) Type is the identifier type to use for this resource.

Optional:

- `name` (String) name is the resource name in the PC
- `uuid` (String) uuid is the UUID of the resource in the PC.


<a id="nestedatt--spec--subnet"></a>
### Nested Schema for `spec.subnet`

Required:

- `type` (String) Type is the identifier type to use for this resource.

Optional:

- `name` (String) name is the resource name in the PC
- `uuid` (String) uuid is the UUID of the resource in the PC.


<a id="nestedatt--spec--additional_categories"></a>
### Nested Schema for `spec.additional_categories`

Optional:

- `key` (String) key is the Key of the category in the Prism Central.
- `value` (String) value is the category value linked to the key in the Prism Central.


<a id="nestedatt--spec--project"></a>
### Nested Schema for `spec.project`

Required:

- `type` (String) Type is the identifier type to use for this resource.

Optional:

- `name` (String) name is the resource name in the PC
- `uuid` (String) uuid is the UUID of the resource in the PC.


<a id="nestedatt--spec--users"></a>
### Nested Schema for `spec.users`

Required:

- `name` (String)
- `ssh_authorized_keys` (List of String)