resource "k8s_loki_grafana_com_alerting_rule_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
