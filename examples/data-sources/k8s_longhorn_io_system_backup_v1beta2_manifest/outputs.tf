output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_system_backup_v1beta2_manifest.example.yaml
  }
}
