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

data "k8s_monitoring_coreos_com_pod_monitor_v1_manifest" "int_target_port" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    pod_metrics_endpoints = [
      {
        path        = "/metrics"
        port        = "metrics"
        target_port = 1234
      }
    ]
    selector = {
      match_labels = {
        "key" = "value"
      }
    }
  }
}

data "k8s_monitoring_coreos_com_pod_monitor_v1_manifest" "string_target_port" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    pod_metrics_endpoints = [
      {
        path        = "/metrics"
        port        = "metrics"
        target_port = "http"
      }
    ]
    selector = {
      match_labels = {
        "key" = "value"
      }
    }
  }
}
