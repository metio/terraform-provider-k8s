resource "k8s_longhorn_io_backup_volume_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
}
