output "manifests" {
  value = {
    "example" = data.k8s_storage_k8s_io_storage_class_v1_manifest.example.yaml
  }
}
