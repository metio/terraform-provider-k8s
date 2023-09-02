output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_backup_target_v1beta2_manifest.example.yaml
  }
}
