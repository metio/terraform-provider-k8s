data "k8s_fossul_io_restore_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
