resource "k8s_notification_toolkit_fluxcd_io_receiver_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    type      = "some-type"
    resources = []
  }
}
