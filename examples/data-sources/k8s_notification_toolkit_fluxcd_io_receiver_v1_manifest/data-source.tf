data "k8s_notification_toolkit_fluxcd_io_receiver_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}