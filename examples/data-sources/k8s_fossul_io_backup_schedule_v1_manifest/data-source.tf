data "k8s_fossul_io_backup_schedule_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cron_schedule   = "59 23 * * *"
    deployment_name = "mariadb"
    policy          = "daily"
  }
}
