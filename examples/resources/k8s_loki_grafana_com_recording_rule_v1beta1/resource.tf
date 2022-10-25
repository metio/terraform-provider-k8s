resource "k8s_loki_grafana_com_recording_rule_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
