data "k8s_monitoring_coreos_com_pod_monitor_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
