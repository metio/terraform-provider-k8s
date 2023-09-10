resource "k8s_monitoring_coreos_com_prometheus_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
