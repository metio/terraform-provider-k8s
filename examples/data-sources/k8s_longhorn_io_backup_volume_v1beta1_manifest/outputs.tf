output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_backup_volume_v1beta1_manifest.example.yaml
  }
}
