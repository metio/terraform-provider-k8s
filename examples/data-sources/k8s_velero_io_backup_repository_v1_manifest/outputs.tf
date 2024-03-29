output "manifests" {
  value = {
    "example" = data.k8s_velero_io_backup_repository_v1_manifest.example.yaml
  }
}
