data "k8s_notification_toolkit_fluxcd_io_alert_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    exclusion_list = [
      "first",
      "second"
    ]
    provider_ref = {
      name = "test"
    }
    summary        = "some critical alert"
    suspend        = true
    event_severity = "error"
    event_sources  = []
  }
}
