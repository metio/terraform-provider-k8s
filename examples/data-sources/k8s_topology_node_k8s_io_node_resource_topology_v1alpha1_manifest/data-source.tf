data "k8s_topology_node_k8s_io_node_resource_topology_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  zones             = []
  topology_policies = []
}
