output "manifests" {
  value = {
    "example" = data.k8s_postgresql_cnpg_io_backup_v1_manifest.example.yaml
  }
}
