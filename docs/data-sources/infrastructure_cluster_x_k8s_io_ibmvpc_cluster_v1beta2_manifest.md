---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  IBMVPCCluster is the Schema for the ibmvpcclusters API.
---

# k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest (Data Source)

IBMVPCCluster is the Schema for the ibmvpcclusters API.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest" "example" {
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

- `spec` (Attributes) IBMVPCClusterSpec defines the desired state of IBMVPCCluster. (see [below for nested schema](#nestedatt--spec))

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

- `region` (String) The IBM Cloud Region the cluster lives in.
- `resource_group` (String) The VPC resources should be created under the resource group.

Optional:

- `control_plane_endpoint` (Attributes) ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. (see [below for nested schema](#nestedatt--spec--control_plane_endpoint))
- `control_plane_load_balancer` (Attributes) ControlPlaneLoadBalancer is optional configuration for customizing control plane behavior. (see [below for nested schema](#nestedatt--spec--control_plane_load_balancer))
- `network` (Attributes) network represents the VPC network to use for the cluster. (see [below for nested schema](#nestedatt--spec--network))
- `vpc` (String) The Name of VPC.
- `zone` (String) The Name of availability zone.

<a id="nestedatt--spec--control_plane_endpoint"></a>
### Nested Schema for `spec.control_plane_endpoint`

Required:

- `host` (String) The hostname on which the API server is serving.
- `port` (Number) The port on which the API server is serving.


<a id="nestedatt--spec--control_plane_load_balancer"></a>
### Nested Schema for `spec.control_plane_load_balancer`

Optional:

- `additional_listeners` (Attributes List) AdditionalListeners sets the additional listeners for the control plane load balancer. (see [below for nested schema](#nestedatt--spec--control_plane_load_balancer--additional_listeners))
- `id` (String) id of the loadbalancer
- `name` (String) Name sets the name of the VPC load balancer.
- `public` (Boolean) public indicates that load balancer is public or private

<a id="nestedatt--spec--control_plane_load_balancer--additional_listeners"></a>
### Nested Schema for `spec.control_plane_load_balancer.additional_listeners`

Required:

- `port` (Number) Port sets the port for the additional listener.



<a id="nestedatt--spec--network"></a>
### Nested Schema for `spec.network`

Optional:

- `control_plane_subnets` (Attributes List) controlPlaneSubnets is a set of Subnet's which define the Control Plane subnets. (see [below for nested schema](#nestedatt--spec--network--control_plane_subnets))
- `resource_group` (String) resourceGroup is the name of the Resource Group containing all of the newtork resources.This can be different than the Resource Group containing the remaining cluster resources.
- `vpc` (Attributes) vpc defines the IBM Cloud VPC for extended VPC Infrastructure support. (see [below for nested schema](#nestedatt--spec--network--vpc))
- `worker_subnets` (Attributes List) workerSubnets is a set of Subnet's which define the Worker subnets. (see [below for nested schema](#nestedatt--spec--network--worker_subnets))

<a id="nestedatt--spec--network--control_plane_subnets"></a>
### Nested Schema for `spec.network.control_plane_subnets`

Optional:

- `cidr` (String)
- `id` (String)
- `name` (String)
- `zone` (String)


<a id="nestedatt--spec--network--vpc"></a>
### Nested Schema for `spec.network.vpc`

Optional:

- `id` (String) id of the resource.
- `name` (String) name of the resource.


<a id="nestedatt--spec--network--worker_subnets"></a>
### Nested Schema for `spec.network.worker_subnets`

Optional:

- `cidr` (String)
- `id` (String)
- `name` (String)
- `zone` (String)
