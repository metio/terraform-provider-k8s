resource "k8s_notification_toolkit_fluxcd_io_alert_v1beta1" "big" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
    labels = {
      "test" = "abc"
    }
    annotations = {
      "try" = "this"
    }
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

resource "k8s_notification_toolkit_fluxcd_io_alert_v1beta1" "small" {
  metadata = {
    name = "test"
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
