output "resources" {
  value = {
    "minimal" = k8s_storage_k8s_io_storage_class_v1.minimal.yaml
    "example" = k8s_storage_k8s_io_storage_class_v1.example.yaml
  }
}
