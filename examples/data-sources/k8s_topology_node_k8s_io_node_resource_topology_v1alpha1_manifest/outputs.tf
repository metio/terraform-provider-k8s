output "manifests" {
  value = {
    "example" = data.k8s_topology_node_k8s_io_node_resource_topology_v1alpha1_manifest.example.yaml
  }
}
