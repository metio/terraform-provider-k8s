output "manifests" {
  value = {
    "example" = data.k8s_fossul_io_backup_schedule_v1_manifest.example.yaml
  }
}
