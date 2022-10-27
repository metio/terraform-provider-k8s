resource "k8s_longhorn_io_setting_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
  value = "some value"
}
