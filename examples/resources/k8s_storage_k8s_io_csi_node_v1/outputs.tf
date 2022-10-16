output "resources" {
  value = {
    "minimal" = k8s_storage_k8s_io_csi_node_v1.minimal.yaml
    "example" = k8s_storage_k8s_io_csi_node_v1.example.yaml
  }
}
