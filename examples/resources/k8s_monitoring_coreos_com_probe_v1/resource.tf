resource "k8s_monitoring_coreos_com_probe_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
