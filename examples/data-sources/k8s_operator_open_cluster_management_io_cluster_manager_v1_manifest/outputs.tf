output "manifests" {
  value = {
    "example" = data.k8s_operator_open_cluster_management_io_cluster_manager_v1_manifest.example.yaml
  }
}
