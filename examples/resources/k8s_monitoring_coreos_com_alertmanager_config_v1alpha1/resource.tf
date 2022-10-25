resource "k8s_monitoring_coreos_com_alertmanager_config_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
