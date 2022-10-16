output "resources" {
  value = {
    "minimal" = k8s_storage_k8s_io_volume_attachment_v1.minimal.yaml
  }
}
