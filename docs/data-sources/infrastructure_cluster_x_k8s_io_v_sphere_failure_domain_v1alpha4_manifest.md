---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  VSphereFailureDomain is the Schema for the vspherefailuredomains API Deprecated: This type will be removed in one of the next releases.
---

# k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest (Data Source)

VSphereFailureDomain is the Schema for the vspherefailuredomains API Deprecated: This type will be removed in one of the next releases.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1alpha4_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain (see [below for nested schema](#nestedatt--spec))

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

- `region` (Attributes) Region defines the name and type of a region (see [below for nested schema](#nestedatt--spec--region))
- `topology` (Attributes) Topology describes a given failure domain using vSphere constructs (see [below for nested schema](#nestedatt--spec--topology))
- `zone` (Attributes) Zone defines the name and type of a zone (see [below for nested schema](#nestedatt--spec--zone))

<a id="nestedatt--spec--region"></a>
### Nested Schema for `spec.region`

Required:

- `name` (String) Name is the name of the tag that represents this failure domain
- `tag_category` (String) TagCategory is the category used for the tag
- `type` (String) Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'

Optional:

- `auto_configure` (Boolean) AutoConfigure tags the Type which is specified in the Topology


<a id="nestedatt--spec--topology"></a>
### Nested Schema for `spec.topology`

Required:

- `datacenter` (String) The underlying infrastructure for this failure domain Datacenter as the failure domain

Optional:

- `compute_cluster` (String) ComputeCluster as the failure domain
- `datastore` (String) Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.
- `hosts` (Attributes) Hosts has information required for placement of machines on VSphere hosts. (see [below for nested schema](#nestedatt--spec--topology--hosts))
- `networks` (List of String) Networks is the list of networks within this failure domain

<a id="nestedatt--spec--topology--hosts"></a>
### Nested Schema for `spec.topology.hosts`

Required:

- `host_group_name` (String) HostGroupName is the name of the Host group
- `vm_group_name` (String) VMGroupName is the name of the VM group



<a id="nestedatt--spec--zone"></a>
### Nested Schema for `spec.zone`

Required:

- `name` (String) Name is the name of the tag that represents this failure domain
- `tag_category` (String) TagCategory is the category used for the tag
- `type` (String) Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'

Optional:

- `auto_configure` (Boolean) AutoConfigure tags the Type which is specified in the Topology
