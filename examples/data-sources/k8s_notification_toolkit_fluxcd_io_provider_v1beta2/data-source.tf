data "k8s_notification_toolkit_fluxcd_io_provider_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
