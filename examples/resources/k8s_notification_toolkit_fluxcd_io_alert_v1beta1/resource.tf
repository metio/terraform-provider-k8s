resource "k8s_notification_toolkit_fluxcd_io_alert_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}