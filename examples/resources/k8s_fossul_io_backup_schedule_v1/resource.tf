resource "k8s_fossul_io_backup_schedule_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_fossul_io_backup_schedule_v1" "example" {
  metadata = {
    name = "mariadb-sample"
  }
  spec = {
    cron_schedule   = "59 23 * * *"
    deployment_name = "mariadb"
    policy          = "daily"
  }
}
