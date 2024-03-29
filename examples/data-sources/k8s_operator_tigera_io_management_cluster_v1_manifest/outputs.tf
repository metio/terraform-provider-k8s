output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_management_cluster_v1_manifest.example.yaml
  }
}
