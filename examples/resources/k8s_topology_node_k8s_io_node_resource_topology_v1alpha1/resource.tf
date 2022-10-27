resource "k8s_topology_node_k8s_io_node_resource_topology_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  zones             = []
  topology_policies = []
}
