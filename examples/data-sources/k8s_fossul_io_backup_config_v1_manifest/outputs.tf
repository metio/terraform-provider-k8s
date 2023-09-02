output "manifests" {
  value = {
    "example" = data.k8s_fossul_io_backup_config_v1_manifest.example.yaml
  }
}
