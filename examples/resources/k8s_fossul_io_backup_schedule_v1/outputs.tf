output "resources" {
  value = {
    "minimal" = k8s_fossul_io_backup_schedule_v1.minimal.yaml
    "example" = k8s_fossul_io_backup_schedule_v1.example.yaml
  }
}
