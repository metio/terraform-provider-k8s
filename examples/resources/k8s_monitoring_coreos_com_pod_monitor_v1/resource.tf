resource "k8s_monitoring_coreos_com_pod_monitor_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    pod_metrics_endpoints = [
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
