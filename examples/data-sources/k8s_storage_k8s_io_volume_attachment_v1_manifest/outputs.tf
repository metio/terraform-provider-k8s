output "manifests" {
  value = {
    "example" = data.k8s_storage_k8s_io_volume_attachment_v1_manifest.example.yaml
  }
}
