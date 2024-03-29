output "manifests" {
  value = {
    "example" = data.k8s_velero_io_backup_storage_location_v1_manifest.example.yaml
  }
}
