output "manifests" {
  value = {
    "example" = data.k8s_velero_io_delete_backup_request_v1_manifest.example.yaml
  }
}
