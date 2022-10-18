output "resources" {
  value = {
    "minimal" = k8s_operator_open_cluster_management_io_cluster_manager_v1.minimal.yaml
    "example" = k8s_operator_open_cluster_management_io_cluster_manager_v1.example.yaml
  }
}
