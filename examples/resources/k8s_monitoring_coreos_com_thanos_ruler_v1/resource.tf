resource "k8s_monitoring_coreos_com_thanos_ruler_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
