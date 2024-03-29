output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_log_storage_v1_manifest.example.yaml
  }
}
