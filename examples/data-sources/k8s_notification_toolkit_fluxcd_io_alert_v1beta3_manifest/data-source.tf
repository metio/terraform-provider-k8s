data "k8s_notification_toolkit_fluxcd_io_alert_v1beta3_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    provider_ref = {
      name = "test"
    }
    summary        = "some minor alert"
    event_severity = "info"
    event_sources  = []
  }
}