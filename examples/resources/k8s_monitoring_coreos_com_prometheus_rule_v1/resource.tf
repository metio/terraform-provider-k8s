resource "k8s_monitoring_coreos_com_prometheus_rule_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {

  }
}
