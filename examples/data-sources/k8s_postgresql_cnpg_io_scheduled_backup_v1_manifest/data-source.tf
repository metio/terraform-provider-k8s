data "k8s_postgresql_cnpg_io_scheduled_backup_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
