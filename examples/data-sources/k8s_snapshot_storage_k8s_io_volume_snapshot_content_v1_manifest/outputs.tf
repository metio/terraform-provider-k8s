output "manifests" {
  value = {
    "example" = data.k8s_snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest.example.yaml
  }
}
