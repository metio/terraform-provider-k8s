resource "k8s_notification_toolkit_fluxcd_io_provider_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    type = "matrix"
  }
}
