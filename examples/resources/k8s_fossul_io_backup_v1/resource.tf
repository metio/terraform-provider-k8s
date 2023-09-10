resource "k8s_fossul_io_backup_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
