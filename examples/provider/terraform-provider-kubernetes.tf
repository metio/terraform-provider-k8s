# declare any resource from the k8s provider
resource "k8s_monitoring_coreos_com_pod_monitor_v1" "example" {
  metadata = {
    name = "example"
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
        "app.kubernetes.io/name" = "some-name"
      }
    }
  }
}

# use the 'yaml' attribute as input for the kubernetes provider
resource "kubernetes_manifest" "example" {
  manifest = yamldecode(k8s_monitoring_coreos_com_pod_monitor_v1.example.yaml)
}
