resource "k8s_longhorn_io_snapshot_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
