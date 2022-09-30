resource "k8s_monitoring_coreos_com_service_monitor_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    endpoints = [
      {
        path = "/metrics"
        port = "metrics"
      }
    ]
    selector = {
      match_labels = {
        "key" = "value"
      }
    }
  }
}
