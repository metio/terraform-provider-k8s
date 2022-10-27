resource "k8s_longhorn_io_volume_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
