resource "k8s_monitoring_coreos_com_prometheus_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
