# declare any resource from the k8s provider
data "k8s_monitoring_coreos_com_pod_monitor_v1_manifest" "example" {
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

# use the 'yaml' attribute as input for the kubectl provider
resource "kubectl_manifest" "example" {
  yaml_body = data.k8s_monitoring_coreos_com_pod_monitor_v1_manifest.example.yaml
}
