resource "k8s_monitoring_coreos_com_alertmanager_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
